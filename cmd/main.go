package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/config"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/routes"
)

func main() {
	app := mux.NewRouter()
	url := fmt.Sprintf(":%d", config.Port)
	routes.Register_Create_Session_route(app)
	log.Println("Starting the server on ", config.Service)
	http.Handle("/", app)
	log.Println("Server Started")

	log.Fatal(http.ListenAndServe(url, app))

}
