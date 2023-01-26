package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/routes"
)

func main() {
	app := mux.NewRouter()

	routes.Register_Create_Session_route(app)
	http.Handle("/", app)
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe("localhost:8081", app))

}
