package routes

import (
	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/controllers"
)

func Register_Create_Session_route(router *mux.Router) {
	router.HandleFunc("/wd/hub/session", controllers.Create_selenium_controller).Methods("POST")
	router.HandleFunc("/book/", controllers.Create_selenium_controller).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.Create_selenium_controller).Methods("DELETE")
}

func create_selenium() {

}
