package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Response, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	} else {
		fmt.Println(err)
	}

}
func ConstructCreateSessionURL(givenUrl string) (url string) {
	url = givenUrl + "session"
	return
}
func ConstructDeleteSessionURL(sessionId, givenUrl string) (url string) {
	url = ConstructCreateSessionURL(givenUrl) + "/" + sessionId
	return
}
func Send_Error_To_Client(e ErrorTemplate, w http.ResponseWriter) {
	return_data, _ := json.Marshal(&e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.ErrorCode)
	w.Write(return_data)
}

type ErrorTemplate struct {
	ErrorMessage string
	ErrorCode    int
}
