package models

import "github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"

// Pod constants
const (
	POD_API_VERSION   = "v1"
	POD_KIND          = "Pod"
	POD_URL_POSTFIX   = "pods/"
	IMAGE_PULL_POLICY = "Always"
	RestartPolicy     = "Never"
)

// Image constants
const (
	IMAGE_SELENIUM_HUB_V4       = ""
	IMAGE_SELENIUM_NODE_CHROME  = ""
	IMAGE_SELENIUM_NODE_FIREFOX = ""
	PLATFORM_LINUX              = "linux"
)

var hubImages map[string]string
var nodeImages map[string]map[string]string

func init() {
	hubImages = make(map[string]string)
	nodeImages = make(map[string]map[string]string)
}

type Pod struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name   string `json:"name"`
		Labels struct {
			App string `json:"app"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
		Containers    []Container `json:"containers"`
		RestartPolicy string      `json:"restartPolicy"`
	} `json:"spec"`
}
type Resource struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type Container struct {
	Envs            []Env  `json:"env"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	Name            string `json:"name"`
	Resources       struct {
		Limits   Resource `json:"limits"`
		Requests Resource `json:"requests"`
	} `json:"resources"`
	Ports []PortStruct `json:"ports"`
}

type PortStruct struct {
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

func (p *Pod) Init(c Common) {
	p.Metadata.Name = "pod-" + c.App
	p.Metadata.Labels.App = c.App
	p.Spec.Selector.App = c.App

	p.SetValues()

}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (p *Pod) constructUrl(baseUrl, namespace string) (url string) {
	return baseUrl + "api/v1/namespaces/" + namespace + "/" + POD_URL_POSTFIX
}
func (p *Pod) appendHubContainer() {
	hub := Container{}
	requests := Resource{"cpu", "ram"}
	limits := Resource{"cpu", "ram"}
	hub.Resources.Limits = limits
	hub.Resources.Requests = requests
	hub.Name = "<hub-name>"
	hub.Image = "<hub-image>"
	hub.Ports = append(hub.Ports, PortStruct{4444, constants.PROTOCOL})
	hub.ImagePullPolicy = IMAGE_PULL_POLICY
	p.Spec.Containers = append(p.Spec.Containers, hub)
}
func (p *Pod) appendNodeContainer() {
	/*
		"SE_EVENT_BUS_HOST", "localhost"
		"SE_EVENT_BUS_PUBLISH_PORT", "4442"
		"SE_EVENT_BUS_SUBSCRIBE_PORT",  "4443"
		"REMOTE_HOST", "http://${HOSTNAME}:5555"
		"SE_NODE_PORT", "5555"
		"SE_NODE_SESSION_TIMEOUT", "600"
		"SE_NODE_MAX_SESSIONS", "1"
		"SE_DRAIN_AFTER_SESSION_COUNT", "1"

	*/

	node := Container{}
	envArr := append(node.Envs,
		Env{"SE_EVENT_BUS_HOST", "localhost"},
		Env{"SE_EVENT_BUS_PUBLISH_PORT", "4442"},
		Env{"SE_EVENT_BUS_SUBSCRIBE_PORT", "4443"},
		Env{"REMOTE_HOST", "http://${HOSTNAME}:5555"},
		Env{"SE_NODE_PORT", "5555"},
		Env{"SE_NODE_SESSION_TIMEOUT", "600"},
		Env{"SE_NODE_MAX_SESSIONS", "1"},
		Env{"SE_DRAIN_AFTER_SESSION_COUNT", "1"})
	node.Envs = envArr
	requests := Resource{"cpu", "ram"}
	limits := Resource{"cpu", "ram"}
	node.Resources.Limits = limits
	node.Resources.Requests = requests
	node.Name = "<node-name>"
	node.Image = "<node-image>"
	node.Ports = append(node.Ports, PortStruct{5555, constants.PROTOCOL})
	node.ImagePullPolicy = IMAGE_PULL_POLICY
	p.Spec.Containers = append(p.Spec.Containers, node)
}
func (p *Pod) Deploy() {

}
func (p *Pod) GetName() (name string) {
	return name
}
func (p *Pod) SetValues() {
	p.Kind = POD_KIND
	p.APIVersion = POD_API_VERSION
	p.Spec.RestartPolicy = RestartPolicy
	p.appendHubContainer()
	p.appendNodeContainer()
}
