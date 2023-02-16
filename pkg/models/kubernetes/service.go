package models

import (
	"encoding/json"
	"fmt"
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

// TODO change the hardcoded values to config or fetch from service struct
func (s *Service) GetServiceUrl() string {
	return "http://" + s.GetName() + ":" + strconv.Itoa(HUB_PORT)
}
func (s *Service) SaveNamespaceDetails(n NamespaceDetails) {
	s.namespaceDetails = n
}
func (s *Service) Init(c Common, n NamespaceDetails) {
	s.SaveNamespaceDetails(n)
	s.SetValues()
	metadata := &s.Metadata
	spec := &s.Spec
	metadata.Name = "service-" + c.App
	metadata.Labels.App = c.App
	spec.Ports = append(spec.Ports, Port{Port: c.Port, Protocol: constants.PROTOCOL, Name: c.App})
	spec.Selector.App = c.App
}
func (s *Service) Delete(string) error {
	if s.Metadata.Name == "" {
		return fmt.Errorf("service name cannot be empty for delete")
	}
	log.Println("Deleting Service", s.GetName())
	response := utils.Make_Delete_Call_With_Bearer(s.constructDeleteUrl(), s.namespaceDetails.Token)
	response.Printf("service delete response:")
	if response.Err != nil {
		return response.Err
	}
	return nil
}
func (s *Service) SetName(serviceName string) {
	s.Metadata.Name = serviceName

}
func (s *Service) constructUrl() (url string) {
	return s.namespaceDetails.Url + "api/v1/namespaces/" + s.namespaceDetails.Namespace + "/" + SERVICE_URL_POSTFIX
}
func (s *Service) constructDeleteUrl() (url string) {
	return s.constructUrl() + s.GetName()
}

func (s *Service) Deploy() error {
	bytes, err := json.Marshal(s)
	if err != nil {
		log.Println("Error in service marshall", err)
		return err
	}
	response := utils.Make_Post_Call_With_Bearer(s.constructUrl(), bytes, s.namespaceDetails.Token)
	response.Printf("Service response:")
	if response.Err != nil {
		return response.Err
	}
	return nil
}
func (s *Service) GetName() (name string) {
	return s.Metadata.Name
}
func (s *Service) SetValues() {
	s.Kind = SERVICE_KIND
	s.APIVersion = SERVICE_API_VERSION

}
