package models

//service constants
const (
	SERVICE_API_VERSION = "v1"
	SERVICE_KIND        = "Service"
	SERVICE_URL_POSTFIX = "services/"
)

type Service struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		GenerateName string `json:"generateName"`
		Labels       struct {
			App string `json:"app"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Port     int    `json:"port"`
			Protocol string `json:"protocol"`
			Name     string `json:"name"`
		} `json:"ports"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
	} `json:"spec"`
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
