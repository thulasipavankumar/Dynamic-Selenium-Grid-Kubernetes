package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func CreateSession(m []byte, posturl string) (resData []byte, err error) {
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(m))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	res, err := client.Do(r)
	//defer res.Body.Close()

	newSession := selenim{}
	if res.StatusCode == http.StatusOK {
		//utils.ParseBody(res, newSession)
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal([]byte(data), &newSession)
		log.Println("erro: ", err, string(data))
		log.Printf(" data is %#v", newSession.Value)
		return data, nil

	} else {
		log.Printf("Unkown response %v, %#v", res.StatusCode, res.Body)
		return nil, fmt.Errorf("An error occurred whilemake json call")
	}

}
