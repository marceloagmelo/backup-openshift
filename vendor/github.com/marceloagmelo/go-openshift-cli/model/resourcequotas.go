package model

import "time"

//ResourceQuotas dados
type ResourceQuotas struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Spec struct {
			Hard struct {
				LimitsMemory string `json:"limits.memory"`
			} `json:"hard"`
			Scopes []string `json:"scopes"`
		} `json:"spec,omitempty"`
		Status struct {
			Hard struct {
				LimitsMemory string `json:"limits.memory"`
			} `json:"hard"`
			Used struct {
				LimitsMemory string `json:"limits.memory"`
			} `json:"used"`
		} `json:"status"`
	} `json:"items"`
}
