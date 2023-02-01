package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
)

func Make_Get_Call(url string) Response {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{ResData: nil, Err: fmt.Errorf("Unable to Create GET Request object"), ResponseCode: constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	return makeClientCall(r)
}
func Make_Delete_Call(url string) Response {
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return Response{ResData: nil, Err: fmt.Errorf("Unable to Create Delete Request object"), ResponseCode: constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	return makeClientCall(r)
}
func Make_Post_Call(url string, body []byte) Response {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return Response{ResData: nil, Err: err, ResponseCode: constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	return makeClientCall(r)
}
func Make_Post_Call_With_Bearer(url string, body []byte, token string) Response {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	bearer := "Bearer " + token
	r.Header.Add("Authorization", bearer)
	if err != nil {
		return Response{ResData: nil, Err: err, ResponseCode: constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	return makeClientCall(r)
}
func makeClientCall(r *http.Request) Response {
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
