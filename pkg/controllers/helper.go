package controllers

import (
	"net/http"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

func send_Error_To_Client(w http.ResponseWriter, errorMessage string, errorCode int) {
	utils.Send_Error_To_Client(utils.ErrorTemplate{
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
	}, w)
}
func createDeployment() models.Deployment {
	return models.Deployment{}
}
