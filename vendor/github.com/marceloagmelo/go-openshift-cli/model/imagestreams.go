package model

import "time"

//ImageStreams dados
type ImageStreams struct {
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
			Generation        int       `json:"generation"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata,omitempty"`
		Spec struct {
			LookupPolicy struct {
				Local bool `json:"local"`
			} `json:"lookupPolicy"`
			Tags []struct {
				Name        string      `json:"name"`
				Annotations interface{} `json:"annotations"`
				From        struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
				} `json:"from"`
				Reference    bool `json:"reference"`
				Generation   int  `json:"generation"`
				ImportPolicy struct {
				} `json:"importPolicy"`
				ReferencePolicy struct {
					Type string `json:"type"`
				} `json:"referencePolicy"`
			} `json:"tags"`
		} `json:"spec,omitempty"`
		Status struct {
			DockerImageRepository string `json:"dockerImageRepository"`
			Tags                  []struct {
				Tag   string `json:"tag"`
				Items []struct {
					Created              time.Time `json:"created"`
					DockerImageReference string    `json:"dockerImageReference"`
					Image                string    `json:"image"`
					Generation           int       `json:"generation"`
				} `json:"items"`
			} `json:"tags"`
		} `json:"status,omitempty"`
	} `json:"items"`
}
