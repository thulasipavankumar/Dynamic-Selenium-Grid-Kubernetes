package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/constants"
	"github.com/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/pkg/utils"
)

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
	HUB_PORT                    = 4444
)

func init() {
	// hubImages = make(map[string]string)
	// nodeImages = make(map[string]map[string]string)
	// err := godotenv.Load("../keys.env")
	// _ = err
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(cwd)
	// hub_image = os.Getenv("hub_image")
	// node_image = os.Getenv("node_image")
}

type Pod struct {
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
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
		Containers    []Container `json:"containers"`
		RestartPolicy string      `json:"restartPolicy"`
	} `json:"spec"`
	nodeImage string
	hubImage  string
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

func (p *Pod) SaveNamespaceDetails(n NamespaceDetails) {
	p.namespaceDetails = n
}
func (p *Pod) Init(c Common, n NamespaceDetails) {
	p.Metadata.Name = "pod-" + c.App
	p.Metadata.Labels.App = c.App
	p.Spec.Selector.App = c.App
	p.SaveNamespaceDetails(n)
	p.SetValues()

}
func (p *Pod) PopulateImagesFronRequest(m Match) error {
	var err error
	p.hubImage, err = GetHubImage("4.0")
	if err != nil {
		return err
	}

	if m.BrowserName == "chrome" {
		p.nodeImage, _ = GetChromeNodeImage(m.BrowserVersion)
		if err != nil {
			return err
		}
	} else {
		p.nodeImage, _ = GetFirefoxNodeImage(m.BrowserVersion)
		if err != nil {
			return err
		}
	}
	return nil
}

type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (p *Pod) Delete() error {
	if p.Metadata.Name == "" {
		return fmt.Errorf("pod name cannot be empty for delete")
	}
	log.Println("Deleting Pod", p.GetName())
	response := utils.Make_Delete_Call_With_Bearer(p.constructDeleteUrl(), p.namespaceDetails.Token)
	response.Printf("Pod delete response:")
	if response.Err != nil {
		return response.Err
	}
	return nil
}
func (p *Pod) SetName(podName string) {
	p.Metadata.Name = podName

}
func (p *Pod) constructUrl() (url string) {
	return p.namespaceDetails.Url + "api/v1/namespaces/" + p.namespaceDetails.Namespace + "/" + POD_URL_POSTFIX
}
func (p *Pod) constructDeleteUrl() (url string) {
	return p.constructUrl() + p.GetName()
}
func (p *Pod) appendHubContainer() {
	hub := Container{}
	requests := Resource{"300m", "500Mi"} // <----- CPU, RAM
	limits := Resource{"300m", "500Mi"}   // <----- CPU, RAM
	hub.Resources.Limits = limits
	hub.Resources.Requests = requests
	hub.Name = "selenium-hub" // <---- "<hub-name>"
	hub.Image = p.hubImage    // <---- "<hub-image>"
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
	requests := Resource{"300m", "500Mi"} // <----- CPU, RAM
	limits := Resource{"300m", "500Mi"}   // <----- CPU, RAM
	node.Resources.Limits = limits
	node.Resources.Requests = requests
	node.Name = "selenium-node" // <----- "<node-name>"
	// <----- "<node-image>"
	node.Image = p.nodeImage
	node.Ports = append(node.Ports, PortStruct{5555, constants.PROTOCOL})
	node.ImagePullPolicy = IMAGE_PULL_POLICY
	p.Spec.Containers = append(p.Spec.Containers, node)
}
func (p *Pod) Deploy() error {
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Println("Error in pod marshall", err)
		return err
	}
	response := utils.Make_Post_Call_With_Bearer(p.constructUrl(), bytes, p.namespaceDetails.Token)
	response.Printf("Pod response:")
	if response.Err != nil {
		return response.Err
	}
	return nil
}
func (p *Pod) GetName() string {
	return p.Metadata.Name
}
func (p *Pod) SetValues() {
	p.Kind = POD_KIND
	p.APIVersion = POD_API_VERSION
	p.Spec.RestartPolicy = RestartPolicy
	p.appendHubContainer()
	p.appendNodeContainer()
}
