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
