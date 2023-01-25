package models

//ingress constants
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
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
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
		Rules []struct {
			HTTP struct {
				Paths []struct {
					Path     string `json:"path"`
					PathType string `json:"pathType"`
					Backend  struct {
						Service struct {
							Name string `json:"name"`
							Port struct {
								Number int `json:"number"`
							} `json:"port"`
						} `json:"service"`
					} `json:"backend"`
				} `json:"paths"`
			} `json:"http"`
		} `json:"rules"`
	} `json:"spec"`
}

func (i *Ingress) constructUrl(baseUrl, namespace string) (url string) {

	return baseUrl + "apis/networking.k8s.io/v1/namespaces/" + namespace + "/"
}
func (i *Ingress) Deploy() {

}
func (i *Ingress) GetName() (name string) {
	return name
}
func (i *Ingress) SetValues() {
	i.Kind = INGRESS_KIND
	i.APIVersion = INGRESS_API_VERSION
}
