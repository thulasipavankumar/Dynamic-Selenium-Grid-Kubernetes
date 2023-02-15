package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/config"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
)

var userName, password, database, host string

func init() {
	godotenv.Load("keys.env")
	userName = os.Getenv("DBUserName")
	password = os.Getenv("DBPassword")
	database = os.Getenv("DBDatabase")
	host = os.Getenv("DBHost")
	config.Connect(userName, password, host, database)
	config.GetDB().AutoMigrate(&models.DatabaseModel{})
}
func main() {
	config.GetDB().NewRecord(models.DatabaseModel{})
	models.AddValuesToDB(models.DatabaseModel{
		SessionID:  "sessionId",
		Service:    "Service",
		Pod:        "Pod",
		Ingress:    "Ingress",
		ServiceUrl: "ServiceUrl",
	})
	fmt.Println("Done")
}
