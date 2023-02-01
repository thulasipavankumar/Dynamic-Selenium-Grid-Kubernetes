package models

import (
	"encoding/json"
	"log"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

// ingress constants
const (
	INGRESS_API_VERSION   = "networking.k8s.io/v1"
	INGRESS_KIND          = "Ingress"
	INGRESS_PATH_WILDCARD = `(/|$)(.*)`
	INGRESS_TYPE_PREFIX   = "Prefix"
	INGRESS_TYPE_EXACT    = "Exact"
	INGRESS_TYPE_MIXED    = "Mixed"
	INGRESS_URL_POSTFIX   = "ingresses/"
)

type Ingress struct {
	namespaceDetails NamespaceDetails
	APIVersion       string `json:"apiVersion"`
	Kind             string `json:"kind"`
	Metadata         struct {
		Name        string `json:"name"`
		Annotations struct {
			NginxIngressKubernetesIoRewriteTarget string `json:"nginx.ingress.kubernetes.io/rewrite-target"`
			NginxIngressKubernetesIoProxyBodySize string `json:"nginx.ingress.kubernetes.io/proxy-body-size"`
		} `json:"annotations"`
		Labels struct {
			App string `json:"app"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
		Rules []Rule `json:"rules"`
	} `json:"spec"`
}
type Rule struct {
	HTTP HTTPStruct `json:"http"`
}
type HTTPStruct struct {
	Paths []PathStruct `json:"paths"`
}
type PathStruct struct {
	Path     string        `json:"path"`
	PathType string        `json:"pathType"`
	Backend  BackendStruct `json:"backend"`
}
type BackendStruct struct {
	Service ServiceStruct `json:"service"`
}
type ServiceStruct struct {
	Name string      `json:"name"`
	Port PortService `json:"port"`
}
type PortService struct {
	Number int `json:"number"`
}

func (i *Ingress) createSeleniumPath() {
	port := PortService{4444}
	service := ServiceStruct{"<service-name-selenium>", port}
	backend := BackendStruct{service}
	path := PathStruct{"session/<sessionID>/(.+)", "Prefix", backend}
	http := HTTPStruct{}
	http.Paths = append(http.Paths, path)
	rule := &i.Spec.Rules[0]
	rule.HTTP.Paths = append(rule.HTTP.Paths, path)
}
func (i *Ingress) createDeletePath() {
	port := PortService{8080}
	service := ServiceStruct{"<service-name-dynamic-grid>", port}
	backend := BackendStruct{service}
	path := PathStruct{"session/<sessionID>", "Prefix", backend}
	http := HTTPStruct{}
	http.Paths = append(http.Paths, path)
	rule := &i.Spec.Rules[0]
	rule.HTTP.Paths = append(rule.HTTP.Paths, path)
}
func (i *Ingress) Init(c Common, n NamespaceDetails) {
	i.namespaceDetails = n
	i.SetValues()

	metadata := &i.Metadata
	spec := &i.Spec
	metadata.Name = "ingress-" + c.App
	metadata.Labels.App = c.App
	spec.Selector.App = c.App

	rule := Rule{}
	spec.Rules = append(spec.Rules, rule)

	i.createSeleniumPath()
	i.createDeletePath()
}
func (i *Ingress) constructUrl() (url string) {

	return i.namespaceDetails.Url + "apis/networking.k8s.io/v1/namespaces/" + i.namespaceDetails.Namespace + "/"
}
func (i *Ingress) Deploy() {
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Println("Error in pod marshall", err)
	}
	response := utils.Make_Post_Call_With_Bearer(i.constructUrl(), bytes, i.namespaceDetails.Token)
	log.Println(response)
}
func (i *Ingress) GetName() (name string) {
	return name
}
func (i *Ingress) SetValues() {
	i.Kind = INGRESS_KIND
	i.APIVersion = INGRESS_API_VERSION
}
