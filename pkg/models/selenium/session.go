package models

import (
	"fmt"
	"log"
	"strings"
)

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
		AlwaysMatch Match `json:"alwaysMatch"`

		FirstMatch []Match `json:"firstMatch"`
	} `json:"capabilities"`
}
type Match struct {
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	PlatformName   string `json:"platformName"`
}

func (m *Match) hasBrowserName() bool {
	return m.BrowserName != ""
}
func (m *Match) hasBrowserVersion() bool {
	return m.BrowserVersion != ""
}
func (m *Match) hasPlatformName() bool {
	return m.PlatformName != ""
}

var defaultBrowserName, defaultBrowserVersion, defaultPlatformName string

func init() {
	defaultBrowserName = "chrome"
	defaultBrowserVersion = "104.0"
	defaultPlatformName = "linux"
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
	if strings.ToLower(s.Capabilities.AlwaysMatch.PlatformName) == "windows" {
		return false
	}
	return true

}
func (s Session) GetValidatedSession() (Match, error) {
	match := Match{}
	if !s.IsValidSession() {
		return Match{}, fmt.Errorf("not a valid session object")
	}
	match = *s.mapFirstMatch(&match)
	match = *s.mapAlwaysMatch(&match)
	log.Printf("session: returning the match object:%v\n", match)
	return match, nil
	//return Match{BrowserName: defaultBrowserName, BrowserVersion: defaultBrowserVersion, PlatformName: defaultPlatformName}, nil

}

func (s Session) mapFirstMatch(match *Match) *Match {
	firstMatch := &s.Capabilities.FirstMatch
	for _, val := range *firstMatch {
		if val.hasPlatformName() {
			match.PlatformName = val.PlatformName
		}
		if val.hasBrowserName() {
			match.BrowserName = val.BrowserName
		}
		if val.hasBrowserVersion() {
			match.BrowserVersion = val.BrowserVersion
		}
	}

	return match
}
func (s Session) mapAlwaysMatch(match *Match) *Match {
	val := &s.Capabilities.AlwaysMatch

	if val.hasPlatformName() {
		match.PlatformName = val.PlatformName
	}
	if val.hasBrowserName() {
		match.BrowserName = val.BrowserName
	}
	if val.hasBrowserVersion() {
		match.BrowserVersion = val.BrowserVersion
	}

	return match
}
