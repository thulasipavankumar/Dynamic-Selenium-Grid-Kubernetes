package models

type Ingress struct {
	Name string `json:"title" bson:"title"`
	// Author    string             `json:"author" bson:"author,omitempty"`
	Sercice_Port int `json:"createdAt" bson:"createdAt"`

	App         string
	Serice_Name string
}
