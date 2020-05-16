package model

import "time"

//LimitRange dados
type LimitRange struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
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
}
