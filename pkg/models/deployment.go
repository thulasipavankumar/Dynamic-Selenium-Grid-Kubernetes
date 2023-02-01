package models

import (
	"fmt"
	"strings"

	"github.com/lithammer/shortuuid/v4"
)

type template interface {
	Deploy()
	SetValues()
	GetName() string
	Init(c Common)
}
type Deployment struct {
	pod     Pod
	service Service
	ingress Ingress
}
type Common struct {
	App    string
	EnvArr []Env
	Port   int
}
type Details struct {
	PodName     string
	ServiceName string
	IngressName string
}

func (d *Deployment) GetDetails() Details {
	return Details{
		PodName:     d.pod.GetName(),
		ServiceName: d.service.GetName(),
		IngressName: d.ingress.GetName(),
	}
}
func (d *Deployment) Deploy() {
	c := Common{App: "app-" + strings.ToLower(shortuuid.New()), EnvArr: nil, Port: 4444}
	d.Init(c)
	d.deployService()
	d.deployPod()
	//	d.deployIngress()
}
func (d *Deployment) Init(c Common) {
	d.pod.Init(c)
	d.service.Init(c)
	d.ingress.Init(c)

	fmt.Printf("%#v \n%#v	\n%#v", d.pod, d.service, d.ingress)

}
func (d *Deployment) deployPod() {
	d.pod.Deploy()
}
func (d *Deployment) deployService() {
	d.service.Deploy()
}
func (d *Deployment) deployIngress() {
	d.ingress.Deploy()
}

func (i *Deployment) setValues() {

}
