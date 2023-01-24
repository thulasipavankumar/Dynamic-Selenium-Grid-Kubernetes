package models

type Session struct {
	Capabilities struct {
		AlwaysMatch struct {
			BrowserName    string `json:"browserName"`
			BrowserVersion string `json:"browserVersion"`
			PlatformName   string `json:"platformName"`
		} `json:"alwaysMatch"`

		Matches []firstMatch `json:"firstMatch"`
	} `json:"capabilities"`
}

type firstMatch struct {
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	PlatformName   string `json:"platformName"`
}

func (s Session) IsValidSession() bool {
	return false
}
