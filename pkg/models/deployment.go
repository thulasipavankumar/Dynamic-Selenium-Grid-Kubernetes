package models

type deploy interface {
	Deploy()
}
type Deployment struct {
	pod     Pod
	service Service
	ingress Ingress
}
type Details struct {
	PodName     string
	ServiceName string
	IngressName string
}

func (d *Deployment) GetDetails() Details {
	return Details{}
}
func (d *Deployment) Deploy() {
	d.deployService()
	d.deployIngress()
	d.deployPod()

}

func (d *Deployment) deployPod() {

}
func (d *Deployment) deployService() {

}
func (d *Deployment) deployIngress() {

}

type Pod struct {
	Image string `json:"title" bson:"title"`
	Port  int    `json:"createdAt" bson:"createdAt"`
	Env   []map[string]string
	app   string
}
type Ingress struct {
	Name         string `json:"title" bson:"title"`
	Service_Port int    `json:"createdAt" bson:"createdAt"`

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
