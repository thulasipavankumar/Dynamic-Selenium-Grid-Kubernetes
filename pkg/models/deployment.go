package models

type template interface {
	Deploy()
	SetValues()
	GetName() string
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
	return Details{
		PodName:     d.pod.GetName(),
		ServiceName: d.service.GetName(),
		IngressName: d.ingress.GetName(),
	}
}
func (d *Deployment) Deploy() {
	d.deployService()
	d.deployPod()
	d.deployIngress()
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
