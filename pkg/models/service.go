package models

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

// service constants
const (
	SERVICE_API_VERSION = "v1"
	SERVICE_KIND        = "Service"
	SERVICE_URL_POSTFIX = "services/"
)

type Service struct {
	namespaceDetails NamespaceDetails
	APIVersion       string `json:"apiVersion"`
	Kind             string `json:"kind"`
	Metadata         struct {
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

func (s *Service) GetServiceUrl() string {
	return "http://" + s.GetName() + ":" + strconv.Itoa(HUB_PORT)
}
func (s *Service) Init(c Common, n NamespaceDetails) {
	s.namespaceDetails = n
	s.SetValues()
	metadata := &s.Metadata
	spec := &s.Spec
	metadata.Name = "service-" + c.App
	metadata.Labels.App = c.App
	spec.Ports = append(spec.Ports, Port{Port: c.Port, Protocol: constants.PROTOCOL, Name: c.App})
	spec.Selector.App = c.App
}
func (s *Service) constructUrl() (url string) {
	return s.namespaceDetails.Url + "api/v1/namespaces/" + s.namespaceDetails.Namespace + "/" + SERVICE_URL_POSTFIX
}
func (s *Service) Deploy() {
	bytes, err := json.Marshal(s)
	if err != nil {
		log.Println("Error in pod marshall", err)
	}
	response := utils.Make_Post_Call_With_Bearer(s.constructUrl(), bytes, s.namespaceDetails.Token)
	log.Println(response)
}
func (s *Service) GetName() (name string) {
	return s.Metadata.Name
}
func (s *Service) SetValues() {
	s.Kind = SERVICE_KIND
	s.APIVersion = SERVICE_API_VERSION

}
