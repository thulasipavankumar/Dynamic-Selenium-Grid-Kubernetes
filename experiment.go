package main

import (
	"fmt"
	"os"
	"time"

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
	//config.GetDB().AutoMigrate(&models.DatabaseModel{})
}
func main() {

	sessions := models.GetAllOpenSessions()
	fmt.Println(len(sessions))

	for _, session := range sessions {

		diff := time.Now().Sub(session.CreatedAt)
		hours, minutes, seconds := diff.Hours(), diff.Minutes(), diff.Seconds()
		fmt.Println("%v %v %v \n", diff.Hours(), diff.Minutes(), diff.Seconds())
		if minutes > 3 {
			fmt.Println("Succes found entry > 3 minutes ", session)
		} else {
			fmt.Printf("Failed session not satisfying %v \n", session)
		}
		_, _ = hours, seconds

	}

}
