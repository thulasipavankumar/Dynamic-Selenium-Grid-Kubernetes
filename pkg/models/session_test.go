package models

import (
	"encoding/json"
	"testing"
)

const VALID_SESSION string = `{
	"capabilities": {
		"alwaysMatch": {
            "browserName": "chrome",
            "platformName": "LINUX",
            "browserVersion": "104.0"
		},
		"firstMatch": [{"browserName": "chrome"},{"platformName": "LINUX"},{"browserVersion": "104.0"}]
	}
}
`

func TestSessionValidate(t *testing.T) {
	inValidSession := Session{}
	if inValidSession.IsValidSession() {
		t.Fail()
	}
	var validSession Session
	json.Unmarshal([]byte(VALID_SESSION), &validSession)
	if validSession.IsValidSession() == false {
		t.Fail()
	}

}
