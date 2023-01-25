package models

//Pod constants
const (
	POD_API_VERSION   = "v1"
	POD_KIND          = "Pod"
	POD_URL_POSTFIX   = "pods/"
	IMAGE_PULL_POLICY = "Always"
)

// Image constants
const (
	IMAGE_SELENIUM_HUB_V4       = ""
	IMAGE_SELENIUM_NODE_CHROME  = ""
	IMAGE_SELENIUM_NODE_FIREFOX = ""
	PLATFORM_LINUX              = "linux"
)

type Pod struct {
	url        string
	token      string
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		GenerateName string `json:"generateName"`
		Labels       struct {
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

type Container struct {
	Envs            []Env  `json:"env"`
	Image           string `json:"image"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	Name            string `json:"name"`
	Resources       struct {
		Limits struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"limits"`
		Requests struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"requests"`
	} `json:"resources"`
	Ports []struct {
		ContainerPort int    `json:"containerPort"`
		Protocol      string `json:"protocol"`
	} `json:"ports"`
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

	p.Spec.Containers = append(p.Spec.Containers, hub)
}
func (p *Pod) appendNodeContainer() {
	node := Container{}
	p.Spec.Containers = append(p.Spec.Containers, node)
}
func (p *Pod) Deploy() {
	p.appendHubContainer()
	p.appendNodeContainer()
}
func (p *Pod) GetName() (name string) {
	return name
}
func (p *Pod) SetValues() {
	p.Kind = POD_KIND
	p.APIVersion = POD_API_VERSION
}
