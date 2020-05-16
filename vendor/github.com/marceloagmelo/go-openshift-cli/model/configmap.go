package model

//ConfigMap dados
type ConfigMap struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Labels          struct {
			Template string `json:"template"`
		} `json:"labels"`
	} `json:"metadata"`
}
