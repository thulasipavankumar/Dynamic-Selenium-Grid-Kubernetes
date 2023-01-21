package models

type Service struct {
	Name string `json:"title" bson:"title"`
	// Author    string             `json:"author" bson:"author,omitempty"`
	Port int `json:"createdAt" bson:"createdAt"`
	Env  []map[string]string
	App  string
}
