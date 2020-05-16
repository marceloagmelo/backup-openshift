package model

import "time"

//Services dados
type Services struct {
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
			Labels            struct {
				App      string `json:"app"`
				AppName  string `json:"app_name"`
				Template string `json:"template"`
			} `json:"labels"`
		} `json:"metadata,omitempty"`
		Spec struct {
			Ports []struct {
				Name     string `json:"name"`
				Protocol string `json:"protocol"`
				Port     int    `json:"port"`
			} `json:"ports"`
			Selector struct {
				Name             string `json:"name"`
				AppName          string `json:"app_name"`
				Deploymentconfig string `json:"deploymentconfig"`
			} `json:"selector"`
			ClusterIP       string `json:"clusterIP"`
			Type            string `json:"type"`
			SessionAffinity string `json:"sessionAffinity"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct {
			} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
}
