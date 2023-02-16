package controllers

import (
	"net/http"

	kubernetes "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models/kubernetes"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

func send_Error_To_Client(w http.ResponseWriter, errorMessage string, errorCode int) {
	utils.Send_Error_To_Client(utils.ErrorTemplate{
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
	}, w)
}
func createDeployment() kubernetes.Deployment {
	return kubernetes.Deployment{}
}
