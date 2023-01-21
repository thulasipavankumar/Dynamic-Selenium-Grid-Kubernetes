package models

type Pod struct {
	Image string `json:"title" bson:"title"`
	// Author    string             `json:"author" bson:"author,omitempty"`
	Port int `json:"createdAt" bson:"createdAt"`
	Env  []map[string]string
	app  string
}
