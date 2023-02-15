package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/config"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/routes"
)

var userName, password, database, host string

func init() {
	userName = os.Getenv("DBUserName")
	password = os.Getenv("DBPassword")
	database = os.Getenv("DBDatabase")
	host = os.Getenv("DBHost")
	config.Connect(userName, password, host, database)
	config.GetDB().AutoMigrate(&models.DatabaseModel{})
}
func main() {
	app := mux.NewRouter()
	url := fmt.Sprintf(":%d", config.GetPort())
	routes.Register_Create_Session_route(app)
	log.Println("Starting the server on ", config.GetService())
	http.Handle("/", app)
	log.Println("Server Started")

	log.Fatal(http.ListenAndServe(url, app))

}
