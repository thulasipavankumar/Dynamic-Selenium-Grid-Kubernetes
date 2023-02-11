package models

import "fmt"

/*
https://w3c.github.io/webdriver/#processing-capabilities

	{
	    "capabilities": {
	        "alwaysMatch": {
	            "cloud:user": "alice",
	            "cloud:password": "hunter2",
	            "platformName": "linux"
	        },
	        "firstMatch": [
	            {"browserName": "chrome"},
	            {"browserName": "edge"}
	        ]
	    }
	}

Once all capabilities are merged from this example, an endpoint node would receive New Session capabilities identical to:

[

	{"browserName": "chrome", "platformName": "linux"},
	{"browserName": "edge", "platformName": "linux"}

]
*/
type Session struct {
	Capabilities struct {
		AlwaysMatch struct {
			BrowserName    string `json:"browserName"`
			BrowserVersion string `json:"browserVersion"`
			PlatformName   string `json:"platformName"`
		} `json:"alwaysMatch"`

		FirstMatch []struct {
			BrowserName    string `json:"browserName"`
			BrowserVersion string `json:"browserVersion"`
			PlatformName   string `json:"platformName"`
		} `json:"firstMatch"`
	} `json:"capabilities"`
}

var defaultBrowserName, defaultBrowserVersion, defaultPlatformName string

func init() {
	defaultBrowserName = "chrome"
	defaultBrowserVersion = "104.0"
	defaultPlatformName = "linux"
}

type Match struct {
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	PlatformName   string `json:"platformName"`
}

/*
Negative scenario when the request cannot be fullfiled
1. First Match array is empty array
2. Both "Always Match" and "First Match" are empty
3. Overlapping Keys in "Always Match" and "First Match"
*/
func (s Session) IsValidSession() bool {

	if s.Capabilities.AlwaysMatch.BrowserName == "" && s.Capabilities.AlwaysMatch.PlatformName == "" &&
		s.Capabilities.AlwaysMatch.BrowserVersion == "" && len(s.Capabilities.FirstMatch) == 0 {
		return false
	}
	return true

}
func (s Session) GetValidatedSession() (Match, error) {
	if s.IsValidSession() {
		return Match{BrowserName: defaultBrowserName, BrowserVersion: defaultBrowserVersion, PlatformName: defaultPlatformName}, nil
	} else {
		return Match{}, fmt.Errorf("Not a valid session object")
	}

}
func (s Session) experiment() {

}
