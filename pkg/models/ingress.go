package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

/*
/session/7b72c331c5b3afa152cc0a203b2f0362/$1
*/
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
			NginxIngressKubernetesIoProxyTimeout  string `json:"nginx.org/proxy-connect-timeout"`
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

var base string

func init() {
	base = os.Getenv("IngressPrefix") // TODO replace it with the ingress prefix in nginx
}
func (i *Ingress) createSeleniumPath(serviceName, sessionId string) {
	// port := PortService{4444}
	// service := ServiceStruct{serviceName, port}
	// backend := BackendStruct{service}
	// path := PathStruct{fmt.Sprintf("%s/session/%s/(.+)", base, sessionId), "Prefix", backend}
	// http := HTTPStruct{}
	// http.Paths = append(http.Paths, path)
	// rule := &i.Spec.Rules[0]
	// rule.HTTP.Paths = append(rule.HTTP.Paths, path)
	i.addPath(serviceName, fmt.Sprintf("%s/session/%s/(.+)", base, sessionId), "Prefix", 4444)
}
func (i *Ingress) createDeletePath(dynamicGridService, sessionId string, dynamicGridServicePort int) {
	// port := PortService{8080} // TODO read the port from config or pass onby parameters
	// service := ServiceStruct{serviceName, port}
	// backend := BackendStruct{service}
	// path := PathStruct{fmt.Sprintf("%s/session/%v", base, sessionId), "Prefix", backend}
	// http := HTTPStruct{}
	// http.Paths = append(http.Paths, path)
	// rule := &i.Spec.Rules[0]
	// rule.HTTP.Paths = append(rule.HTTP.Paths, path)
	i.addPath(dynamicGridService, fmt.Sprintf("%s/session/%v", base, sessionId), "Prefix", dynamicGridServicePort)
}
func (i *Ingress) createSeleniumPathFor3(serviceName, sessionId string) {
	// port := PortService{4444}
	// service := ServiceStruct{serviceName, port}
	// backend := BackendStruct{service}
	// path := PathStruct{fmt.Sprintf("%s/hub/session/%s/(.+)", base, sessionId), "Prefix", backend}
	// http := HTTPStruct{}
	// http.Paths = append(http.Paths, path)
	// rule := &i.Spec.Rules[0]
	// rule.HTTP.Paths = append(rule.HTTP.Paths, path)
	i.addPath(serviceName, fmt.Sprintf("%s/hub/session/%s/(.+)", base, sessionId), "Prefix", 4444)
}
func (i *Ingress) createDeletePathFor3(serviceName, sessionId string) {
	// port := PortService{8080} // TODO read the port from config or pass onby parameters
	// service := ServiceStruct{serviceName, port}
	// backend := BackendStruct{service}
	// path := PathStruct{fmt.Sprintf("%s/hub/session/%v", base, sessionId), "Prefix", backend}
	// http := HTTPStruct{}
	// http.Paths = append(http.Paths, path)
	// rule := &i.Spec.Rules[0]
	// rule.HTTP.Paths = append(rule.HTTP.Paths, path)
	i.addPath(serviceName, fmt.Sprintf("%s/hub/session/%v", base, sessionId), "Prefix", 8080)
}
func (i *Ingress) addPath(serviceName, pathUrl, pathType string, servicePort int) {
	port := PortService{servicePort} // TODO read the port from config or pass onby parameters
	service := ServiceStruct{serviceName, port}
	backend := BackendStruct{service}
	path := PathStruct{pathUrl, pathType, backend}
	http := HTTPStruct{}
	http.Paths = append(http.Paths, path)
	rule := &i.Spec.Rules[0]
	rule.HTTP.Paths = append(rule.HTTP.Paths, path)
}
func (i *Ingress) SaveNamespaceDetails(n NamespaceDetails) {
	i.namespaceDetails = n
}
func (i *Ingress) Init(c Common, n NamespaceDetails) {
	i.SaveNamespaceDetails(n)
	i.SetValues()

	metadata := &i.Metadata
	spec := &i.Spec
	metadata.Name = "ingress-" + c.App
	metadata.Labels.App = c.App
	spec.Selector.App = c.App
	metadata.Annotations.NginxIngressKubernetesIoProxyBodySize = "1000M"
	metadata.Annotations.NginxIngressKubernetesIoProxyTimeout = "1200s"
	metadata.Annotations.NginxIngressKubernetesIoRewriteTarget = "$1"
	rule := Rule{}
	spec.Rules = append(spec.Rules, rule)

}
func (i *Ingress) SaveServiceAndSession(serviceName, sessionId, dynamicGridService string, dynamicGridServicePort int) {

	metadata := &i.Metadata
	rewriteTarget := "/session/" + sessionId + "/$1"
	metadata.Annotations.NginxIngressKubernetesIoRewriteTarget = rewriteTarget

	//TODO if serviceName or sessionId is empty throw error
	i.createSeleniumPath(serviceName, sessionId)
	//i.createSeleniumPathFor3(serviceName, sessionId)
	i.createDeletePath(dynamicGridService, sessionId, dynamicGridServicePort)
	//i.createDeletePathFor3(serviceName, sessionId)

	log.Printf("Changed the Ingress annotaions:%v\n", metadata)

}
func (i *Ingress) Delete(ingressName string) error {
	log.Println("Deleting Ingress", ingressName)
	response := utils.Make_Delete_Call_With_Bearer(i.constructDeleteUrl(ingressName), i.namespaceDetails.Token)
	response.Printf("Ingress delete response:")
	if response.Err != nil {
		return response.Err
	}
	return nil
}
func (i *Ingress) constructUrl() (url string) {

	return i.namespaceDetails.Url + "apis/networking.k8s.io/v1/namespaces/" + i.namespaceDetails.Namespace + "/ingresses/"
}
func (i *Ingress) constructDeleteUrl(ingressName string) (url string) {

	return i.namespaceDetails.Url + "apis/networking.k8s.io/v1/namespaces/" + i.namespaceDetails.Namespace + "/ingresses/" + ingressName
}
func (i *Ingress) Deploy() error {
	log.Printf(" Ingress before deploy:%v\n", i)
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Println("Error in ingress marshall", err)
		return err
	}
	response := utils.Make_Post_Call_With_Bearer(i.constructUrl(), bytes, i.namespaceDetails.Token)
	response.Printf("Ingress response:")
	if response.Err != nil {
		return response.Err
	}
	if response.ResponseCode != 201 {
		return fmt.Errorf("didnot return 201 responce in ingress")
	}
	return nil
}
func (i *Ingress) GetName() string {
	return i.Metadata.Name
}
func (i *Ingress) SetValues() {
	i.Kind = INGRESS_KIND
	i.APIVersion = INGRESS_API_VERSION
}
