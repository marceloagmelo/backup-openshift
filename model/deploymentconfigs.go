package model

//DeploymentConfigs dados
type DeploymentConfigs struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		Metadata struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata"`
	} `json:"items"`
}
