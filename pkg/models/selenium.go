package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
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

func CreateSession(m []byte, posturl string) (response utils.Response) {

	response = utils.Make_Post_Call(posturl, m)
	newSession := selenim{}
	if response.Err != nil {
		return
	}
	if response.GetResponseCode() == http.StatusOK {
		//utils.ParseBody(res, newSession)

		err := json.Unmarshal([]byte(response.GetResponseData()), &newSession)
		if err != nil {

			return utils.Response{ResData: response.GetResponseData(), Err: err, ResponseCode: constants.Unable_TO_UNMARSHALL_JSON}

		}
		//	log.Printf(" data is %v", newSession.Value)
		return utils.Response{ResData: response.GetResponseData(), Err: nil, ResponseCode: response.GetResponseCode()}

	} else {
		return utils.Response{ResData: nil, Err: fmt.Errorf("An error occurred whilemake json call status code is not 200"), ResponseCode: response.GetResponseCode()}
	}

}
func DeleteSession(sessionId, deleteUrl string) (response utils.Response) {
	r, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		return utils.Response{ResData: nil, Err: fmt.Errorf("Unable to Create Delete Request object"), ResponseCode: constants.Unable_TO_CREATE_REQUEST_OBJECT}
	}
	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		return utils.Response{ResData: nil, Err: err, ResponseCode: 422}
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		//utils.ParseBody(res, newSession)
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Panic(err)
		}
		//	log.Printf(" data is %v", newSession.Value)
		return utils.Response{ResData: data, Err: nil, ResponseCode: res.StatusCode}

	} else {
		log.Printf("Unkown response %v, %v \n", res.StatusCode, res.Body)
		return utils.Response{ResData: nil, Err: fmt.Errorf("An error occurred whilemake json call status code is not 200"), ResponseCode: res.StatusCode}

	}
}
