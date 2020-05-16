package model

//Service dados de service
type Service struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		Labels    struct {
			App      string `json:"app"`
			AppName  string `json:"app_name"`
			Template string `json:"template"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Name       string `json:"name"`
			Protocol   string `json:"protocol"`
			Port       int    `json:"port"`
			TargetPort int    `json:"targetPort"`
		} `json:"ports"`
		Selector struct {
			Deploymentconfig string `json:"deploymentconfig"`
		} `json:"selector"`
		Type string `json:"type"`
	} `json:"spec"`
}
