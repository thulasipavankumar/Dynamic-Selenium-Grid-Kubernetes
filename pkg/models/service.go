package models

import "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"

// service constants
const (
	SERVICE_API_VERSION = "v1"
	SERVICE_KIND        = "Service"
	SERVICE_URL_POSTFIX = "services/"
)

type Service struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name   string `json:"name"`
		Labels struct {
			App string `json:"app"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Ports    []Port `json:"ports"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
	} `json:"spec"`
}

type Port struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Name     string `json:"name"`
}

func (s *Service) Init(c Common) {
	s.SetValues()
	metadata := &s.Metadata
	spec := &s.Spec
	metadata.Name = "service-" + c.App
	metadata.Labels.App = c.App
	spec.Ports = append(spec.Ports, Port{Port: c.Port, Protocol: constants.PROTOCOL, Name: c.App})
	spec.Selector.App = c.App
}
func (s *Service) constructUrl(baseUrl, namespace string) (url string) {
	return baseUrl + "api/v1/namespaces/" + namespace + "/" + SERVICE_URL_POSTFIX
}
func (s *Service) Deploy() {

}
func (s *Service) GetName() (name string) {
	return name
}
func (s *Service) SetValues() {
	s.Kind = SERVICE_KIND
	s.APIVersion = SERVICE_API_VERSION

}
