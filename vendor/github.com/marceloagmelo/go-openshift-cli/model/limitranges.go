package model

import "time"

//LimitRanges dados
type LimitRanges struct {
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
			Limits []struct {
				Type string `json:"type"`
				Max  struct {
					CPU    string `json:"cpu"`
					Memory string `json:"memory"`
				} `json:"max"`
				Min struct {
					CPU    string `json:"cpu"`
					Memory string `json:"memory"`
				} `json:"min"`
				Default struct {
					CPU    string `json:"cpu"`
					Memory string `json:"memory"`
				} `json:"default"`
				DefaultRequest struct {
					CPU    string `json:"cpu"`
					Memory string `json:"memory"`
				} `json:"defaultRequest"`
			} `json:"limits"`
		} `json:"spec"`
	} `json:"items"`
}
