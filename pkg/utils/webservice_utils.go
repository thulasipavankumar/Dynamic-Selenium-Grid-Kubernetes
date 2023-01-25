package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
)

func Make_Get_Call(url string) {

}
func Make_Post_Call(url string, body []byte) Response {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return Response{ResData: nil, Err: err, ResponseCode: constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		return Response{ResData: nil, Err: err, ResponseCode: constants.Unable_TO_CREATE_CLIENT_OBJECT}
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return Response{ResData: nil, Err: err, ResponseCode: constants.Unable_TO_READ_REQUEST_DATA}
	}
	return Response{ResData: data, Err: nil, ResponseCode: res.StatusCode}
}
