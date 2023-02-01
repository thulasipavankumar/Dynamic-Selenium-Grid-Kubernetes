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

func init() {

}
func Create_selenium_Containers(w http.ResponseWriter, r *http.Request) {
	deployment := createDeployment()
	deployment.Deploy()
	Create_Selenium_Session(w, r)
}
func Create_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		send_Error_To_Client(w, err.Error(), constants.Unable_TO_READ_REQUEST_DATA)
		return
	}
	log.Printf("Got data: %v \n", string(responseData))
	session := models.Session{}
	err = json.Unmarshal([]byte(responseData), &session)
	if err != nil {
		send_Error_To_Client(w, err.Error(), constants.Unable_TO_UNMARSHALL_JSON)
		return
	}
	if !session.IsValidSession() {
		send_Error_To_Client(w, "Session Object validation failed", constants.UNABLE_TO_VALIDATE_REQUEST)
		return
	}
	session_result := models.CreateSession(responseData, utils.ConstructCreateSessionURL(os.Getenv("hub_url")))

	if session_result.GetErr() != nil {
		send_Error_To_Client(w, session_result.GetErr().Error(), session_result.GetResponseCode())
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(session_result.GetResponseData())
	}
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
func send_Error_To_Client(w http.ResponseWriter, errorMessage string, errorCode int) {
	utils.Send_Error_To_Client(utils.ErrorTemplate{
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
	}, w)
}
func createDeployment() models.Deployment {
	return models.Deployment{}
}
