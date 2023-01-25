package utils

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/models"
)

func Make_Get_Call(url string) {

}
func Make_Post_Call(url string, body []byte) models.Response {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Panicln(err)
	}
	client := &http.Client{}
	res, err := client.Do(r)
	//defer res.Body.Close()
	if err != nil {
		return models.Response{ResData: nil, Err: err, ResponseCode: 422}
	} else {
		return models.Response{ResData: nil, Err: fmt.Errorf("An error occurred whilemake json call status code is not 200"), ResponseCode: res.StatusCode}
	}
}
