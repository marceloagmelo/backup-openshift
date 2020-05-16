package model

import "time"

//ImageStream dados
type ImageStream struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Generation      int    `json:"generation"`
		Annotations     struct {
			OpenshiftIoDisplayName                string    `json:"openshift.io/display-name"`
			OpenshiftIoImageDockerRepositoryCheck time.Time `json:"openshift.io/image.dockerRepositoryCheck"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		LookupPolicy struct {
			Local bool `json:"local"`
		} `json:"lookupPolicy"`
		Tags []struct {
			Name        string `json:"name"`
			Annotations struct {
				Description                    string `json:"description"`
				IconClass                      string `json:"iconClass"`
				OpenshiftIoDisplayName         string `json:"openshift.io/display-name"`
				OpenshiftIoProviderDisplayName string `json:"openshift.io/provider-display-name"`
				SampleRepo                     string `json:"sampleRepo"`
				Supports                       string `json:"supports"`
				Tags                           string `json:"tags"`
				Version                        string `json:"version"`
			} `json:"annotations,omitempty"`
			From struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
			} `json:"from"`
			Generation   int `json:"generation"`
			ImportPolicy struct {
			} `json:"importPolicy"`
			ReferencePolicy struct {
				Type string `json:"type"`
			} `json:"referencePolicy"`
		} `json:"tags"`
	} `json:"spec"`
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
	} `json:"status"`
}
