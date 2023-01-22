package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
)

func init() {

}
func Create_selenium_controller(w http.ResponseWriter, r *http.Request) {

	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(responseData))
	data, err := models.CreateSession(responseData, "https://box.ic.ing.net/ingress/tchp/default/hub/session")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
