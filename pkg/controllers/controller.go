package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
)

func init() {

}
func Create_selenium_controller(w http.ResponseWriter, r *http.Request) {
	deployment := models.Deployment
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
	response := models.CreateSession(responseData, os.Getenv("hub_url"))
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
