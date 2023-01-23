package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
)

type selenim struct {
	Value struct {
		SessionId    string `json:"sessionId"`
		Capabilities struct {
			BrowserName    string `json:"browserName"`
			BrowserVersion string `json:"browserVersion"`
			PlatformName   string `json:"platformName"`
		} `json:"capabilities"`
	} `json:"value"`
}
type Response struct {
	resData      []byte
	err          error
	responseCode int
}

func (r Response) GetResponseData() (byteSlice []byte) {
	return r.resData
}
func (r Response) GetErr() error {
	return r.err
}
func (r Response) GetResponseCode() int {
	return r.responseCode
}

func CreateSession(m []byte, posturl string) (response Response) {
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(m))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	res, err := client.Do(r)
	//defer res.Body.Close()
	if err != nil {
		return Response{nil, err, 422}
	}
	newSession := selenim{}
	if res.StatusCode == http.StatusOK {
		//utils.ParseBody(res, newSession)
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal([]byte(data), &newSession)
		log.Println("erro: ", err, string(data))
		//	log.Printf(" data is %v", newSession.Value)
		return Response{data, nil, res.StatusCode}

	} else {
		log.Printf("Unkown response %v, %v \n", res.StatusCode, res.Body)
		return Response{nil, fmt.Errorf("An error occurred whilemake json call status code is not 200"), res.StatusCode}

	}

}
func DeleteSession(sessionId, deleteUrl string) (response Response) {
	r, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		return Response{nil, fmt.Errorf("Unable to Create Delete Request object"), constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		return Response{nil, err, 422}
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		//utils.ParseBody(res, newSession)
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		//	log.Printf(" data is %v", newSession.Value)
		return Response{data, nil, res.StatusCode}

	} else {
		log.Printf("Unkown response %v, %v \n", res.StatusCode, res.Body)
		return Response{nil, fmt.Errorf("An error occurred whilemake json call status code is not 200"), res.StatusCode}

	}
}
