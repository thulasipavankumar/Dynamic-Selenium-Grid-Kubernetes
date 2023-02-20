package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/config"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
	kubernetes "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models/kubernetes"
	selenium "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models/selenium"
)

func init() {

}
func Create_selenium3_Containers(match selenium.Match) (*kubernetes.Deployment, error) {
	deployment := createDeployment()
	deployment.SetPodSel3()
	return Create_selenium_Containers(match, &deployment)
	//service := deployment.GetService()
	//return service.GetServiceUrl()
}
func Create_selenium4_Containers(match selenium.Match) (*kubernetes.Deployment, error) {
	deployment := createDeployment()
	return Create_selenium_Containers(match, &deployment)
	//service := deployment.GetService()
	//return service.GetServiceUrl()
}
func Create_selenium_Containers(match selenium.Match, deployment *kubernetes.Deployment) (*kubernetes.Deployment, error) {
	err := deployment.LoadRequestedCapabilites(match)
	if err != nil {
		return &kubernetes.Deployment{}, err
	}
	err = deployment.Deploy()
	if err != nil {
		return &kubernetes.Deployment{}, err
	}
	return deployment, nil
	//service := deployment.GetService()
	//return service.GetServiceUrl()
}
func validate(w http.ResponseWriter, r *http.Request) (data []byte, matched selenium.Match, error bool) {
	emptyMatch := selenium.Match{}
	responseData, readerr := ioutil.ReadAll(r.Body)
	if readerr != nil {
		send_Error_To_Client(w, readerr.Error(), constants.Unable_TO_READ_REQUEST_DATA)
		return nil, emptyMatch, true
	}
	log.Printf("Got data: %v \n", string(responseData))
	session := selenium.Session{}
	unmarshalError := json.Unmarshal([]byte(responseData), &session)
	if unmarshalError != nil {
		send_Error_To_Client(w, unmarshalError.Error(), constants.Unable_TO_UNMARSHALL_JSON)
		return nil, emptyMatch, true
	}
	if !session.IsValidSession() {
		send_Error_To_Client(w, "Session Object validation failed", constants.UNABLE_TO_VALIDATE_REQUEST)
		return nil, emptyMatch, true
	}
	match, err := session.GetValidatedSession()
	_ = err
	return responseData, match, false
}
func Create_Selenium_Session(w http.ResponseWriter, r *http.Request, isSel3 bool) {
	responseData, match, err := validate(w, r)
	if err {
		send_Error_To_Client(w, "Validation of session capabilites failed", 422)
		return
	}
	var deployment *kubernetes.Deployment
	var createErr error
	if isSel3 {
		deployment, createErr = Create_selenium3_Containers(match)
	} else {
		deployment, createErr = Create_selenium4_Containers(match)
	}

	if createErr != nil {
		log.Println("Controller: error in creating deployment", createErr)
		send_Error_To_Client(w, "UNABLE_TO_CREATE_DEPLOYMENT", constants.UNABLE_TO_CREATE_DEPLOYMENT)
		return
	}
	service := deployment.GetService()
	hubEndpoint := fmt.Sprintf("%s/session", service.GetServiceUrl()) // TODO replace the url later

	session_result := selenium.CreateSession(responseData, hubEndpoint)
	// <--- get sessionId

	if session_result.GetErr() != nil {
		send_Error_To_Client(w, session_result.GetErr().Error(), session_result.GetResponseCode())
		return
	}

	newSession := selenium.Selenium{}

	_ = json.Unmarshal([]byte(session_result.GetResponseData()), &newSession)
	i := deployment.GetIngress()
	i.SaveServiceAndSession(deployment.GetDetails().ServiceName, newSession.Value.SessionId, config.GetService(), config.GetPort())
	ingressErr := deployment.DeployIngress()
	if ingressErr != nil {
		log.Println(`Error in creating ingress`, ingressErr)
	}
	time.Sleep(5 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	w.Write(session_result.GetResponseData())
	defer models.AddValuesToDB(models.DatabaseModel{
		SessionID:      newSession.Value.SessionId,
		Service:        deployment.GetDetails().ServiceName,
		Pod:            deployment.GetDetails().PodName,
		Ingress:        deployment.GetDetails().IngressName,
		SessionUrl:     deployment.GetService().GetSessionUrl(),
		Browser:        deployment.GetPod().BrowserName,
		BrowserVersion: deployment.GetPod().BrowserVersion,
	})
	_ = deployment
}
func Create_Selenium3_Session(w http.ResponseWriter, r *http.Request) {
	isSel3 := true
	Create_Selenium_Session(w, r, isSel3)
}
func Create_Selenium4_Session(w http.ResponseWriter, r *http.Request) {
	isSel3 := false
	Create_Selenium_Session(w, r, isSel3)
}
func Delete_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["sessionId"]

	log.Printf("Got sessionId: %v \n", string(sessionId))
	/*
		Below for deleting session it require <service-name>:<port>/sessionId
		In-order to achieve it , there should be a session to service mapping done
	*/
	val := models.GetSessionIDObject(sessionId)
	if val.Pod == "" || val.Service == "" || val.Ingress == "" {
		send_Error_To_Client(w, fmt.Sprintf("Values came empty pod:%s service:%s ingress:%s", val.Pod, val.Service, val.Ingress), 422)
		return
	}
	deployment := kubernetes.Deployment{}

	response := selenium.DeleteSession(sessionId, val.SessionUrl+"/"+sessionId)

	if response.GetErr() != nil {
		send_Error_To_Client(w, response.GetErr().Error(), response.GetResponseCode())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response.GetResponseData())
	}
	defer deployment.DeleteDeployment(val.Pod, val.Service, val.Ingress)
	defer models.DeleteDBCell(sessionId)
}
