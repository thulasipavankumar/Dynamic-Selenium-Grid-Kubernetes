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
	deployment := models.Deployment{}
	deployment.Deploy()
	Create_Selenium_Session(w, r)
}
func Create_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Send_Error_To_Client(utils.ErrorTemplate{
			ErrorMessage: (err.Error()),
			ErrorCode:    constants.Unable_TO_READ_REQUEST_DATA,
		}, w)

		return
	}
	log.Printf("Got data: %v \n", string(responseData))
	session := models.Session{}
	err = json.Unmarshal([]byte(responseData), &session)
	if err != nil {
		utils.Send_Error_To_Client(utils.ErrorTemplate{
			ErrorMessage: (err.Error()),
			ErrorCode:    constants.Unable_TO_UNMARSHALL_JSON,
		}, w)
		return
	}
	if !session.IsValidSession() {
		log.Println(`Please pass a valid session`)
	}
	session_result := models.CreateSession(responseData, utils.ConstructCreateSessionURL(os.Getenv("hub_url")))

	if session_result.GetErr() != nil {
		utils.Send_Error_To_Client(utils.ErrorTemplate{
			ErrorMessage: session_result.GetErr().Error(),
			ErrorCode:    session_result.GetResponseCode(),
		}, w)

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
		utils.Send_Error_To_Client(utils.ErrorTemplate{
			ErrorMessage: response.GetErr().Error(),
			ErrorCode:    response.GetResponseCode(),
		}, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response.GetResponseData())
	}
}
