package models

type deploy interface {
	Deploy()
}
type Deployment struct {
	pod     Pod
	service Service
	ingress Ingress
}

func (d *Deployment) Deploy() {

}

func (d *Deployment) deployPod() {

}
func (d *Deployment) deployService() {

}
func (d *Deployment) deployIngress() {

}

type Pod struct {
	Image string `json:"title" bson:"title"`
	// Author    string             `json:"author" bson:"author,omitempty"`
	Port int `json:"createdAt" bson:"createdAt"`
	Env  []map[string]string
	app  string
}
type Ingress struct {
	Name string `json:"title" bson:"title"`
	// Author    string             `json:"author" bson:"author,omitempty"`
	Service_Port int `json:"createdAt" bson:"createdAt"`

	App         string
	Serice_Name string
}

func (i *Ingress) setValues(name string, port int, app, service string) {
	i.Name = name
	i.Service_Port = port
	i.App = app
	i.Serice_Name = service
}

type Service struct {
	Name string `json:"title" bson:"title"`
	// Author    string             `json:"author" bson:"author,omitempty"`
	Port int `json:"createdAt" bson:"createdAt"`
	Env  []map[string]string
	App  string
}
