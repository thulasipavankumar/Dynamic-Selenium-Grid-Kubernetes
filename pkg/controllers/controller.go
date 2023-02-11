package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

var hub_url string

func init() {
	hub_url = os.Getenv("hub_url")
}
func Create_selenium_Containers(match models.Match) (models.Deployment, error) {
	deployment := createDeployment()
	err := deployment.LoadRequestedCapabilites(match)
	if err != nil {
		return models.Deployment{}, err
	}
	err = deployment.Deploy()
	if err != nil {
		return models.Deployment{}, err
	}
	return deployment, nil
	//service := deployment.GetService()
	//return service.GetServiceUrl()
}
func validate(w http.ResponseWriter, r *http.Request) (data []byte, matched models.Match, error bool) {
	emptyMatch := models.Match{}
	responseData, readerr := ioutil.ReadAll(r.Body)
	if readerr != nil {
		send_Error_To_Client(w, readerr.Error(), constants.Unable_TO_READ_REQUEST_DATA)
		return nil, emptyMatch, true
	}
	log.Printf("Got data: %v \n", string(responseData))
	session := models.Session{}
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
func Create_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	responseData, match, err := validate(w, r)
	if err {
		return
	}
	deployment, createErr := Create_selenium_Containers(match)
	if createErr != nil {
		send_Error_To_Client(w, "UNABLE_TO_CREATE_DEPLOYMENT", constants.UNABLE_TO_CREATE_DEPLOYMENT)
		return
	}

	session_result := models.CreateSession(responseData, utils.ConstructCreateSessionURL(hub_url)) // <--- replace url
	// <--- get sessionId
	if session_result.GetErr() != nil {
		send_Error_To_Client(w, session_result.GetErr().Error(), session_result.GetResponseCode())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(session_result.GetResponseData())
	}
	_ = deployment
}
func Delete_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["sessionId"]

	log.Printf("Got sessionId: %v \n", string(sessionId))
	response := models.DeleteSession(sessionId, utils.ConstructDeleteSessionURL(sessionId, os.Getenv("hub_url")))

	if response.GetErr() != nil {
		send_Error_To_Client(w, response.GetErr().Error(), response.GetResponseCode())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response.GetResponseData())
	}
}
