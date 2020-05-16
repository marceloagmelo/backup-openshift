package model

import "time"

//Routes dados
type Routes struct {
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
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Annotations       struct {
				OpenshiftIoHostGenerated string `json:"openshift.io/host.generated"`
			} `json:"annotations"`
		} `json:"metadata,omitempty"`
		Spec struct {
			Host string `json:"host"`
			To   struct {
				Kind   string `json:"kind"`
				Name   string `json:"name"`
				Weight int    `json:"weight"`
			} `json:"to"`
			TLS struct {
				Termination string `json:"termination"`
			} `json:"tls"`
			WildcardPolicy string `json:"wildcardPolicy"`
		} `json:"spec,omitempty"`
		Status struct {
			Ingress []struct {
				Host       string `json:"host"`
				RouterName string `json:"routerName"`
				Conditions []struct {
					Type               string    `json:"type"`
					Status             string    `json:"status"`
					LastTransitionTime time.Time `json:"lastTransitionTime"`
				} `json:"conditions"`
				WildcardPolicy string `json:"wildcardPolicy"`
			} `json:"ingress"`
		} `json:"status"`
	} `json:"items"`
}
