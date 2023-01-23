package routes

import (
	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/controllers"
)

func Register_Create_Session_route(router *mux.Router) {
	router.HandleFunc("/wd/hub/session", controllers.Create_selenium_controller).Methods("POST")
	router.HandleFunc("/wd/hub/session", controllers.Create_selenium_controller).Methods("GET")
	router.HandleFunc("/wd/hub/session/{sessionId}", controllers.Delete_Selenium_Session).Methods("DELETE")
}

func create_selenium() {

}
