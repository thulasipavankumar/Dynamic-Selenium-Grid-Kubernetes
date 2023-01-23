package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

func init() {

}
func Create_selenium_controller(w http.ResponseWriter, r *http.Request) {
	deployment := models.Deployment{}
	deployment.Deploy()
	Create_Selenium_Session(w, r)
}
func Create_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	log.Printf("Got data: %v \n", string(responseData))
	response := models.CreateSession(responseData, utils.ConstructCreateSessionURL(os.Getenv("hub_url")))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.GetResponseCode())
	if response.GetErr() != nil {
		errData := struct {
			Err string
		}{
			Err: response.GetErr().Error(),
		}
		return_data, _ := json.Marshal(&errData)
		w.Write(return_data)
	} else {
		w.Write(response.GetResponseData())
	}
}
func Delete_Selenium_Session(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["sessionId"]

	log.Printf("Got sessionId: %v \n", string(sessionId))
	response := models.DeleteSession(sessionId, utils.ConstructDeleteSessionURL(sessionId, os.Getenv("hub_url")))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.GetResponseCode())
	if response.GetErr() != nil {
		errData := struct {
			Err string
		}{
			Err: response.GetErr().Error(),
		}
		return_data, _ := json.Marshal(&errData)
		w.Write(return_data)
	} else {
		w.Write(response.GetResponseData())
	}
}
