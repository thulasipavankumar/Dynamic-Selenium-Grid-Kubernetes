package models

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

type firstMatch struct {
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	PlatformName   string `json:"platformName"`
}

func (s Session) IsValidSession() bool {
	if s.Capabilities.AlwaysMatch.BrowserName == "" && s.Capabilities.AlwaysMatch.PlatformName == "" &&
		s.Capabilities.AlwaysMatch.BrowserVersion == "" && len(s.Capabilities.FirstMatch) == 0 {
		return false
	}
	return true

}
