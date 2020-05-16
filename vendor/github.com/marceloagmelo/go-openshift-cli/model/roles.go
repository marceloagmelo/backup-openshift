package model

import "time"

//Roles dados
type Roles struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink string `json:"selfLink"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Rules []struct {
			Verbs                 []string    `json:"verbs"`
			AttributeRestrictions interface{} `json:"attributeRestrictions"`
			APIGroups             []string    `json:"apiGroups"`
			Resources             []string    `json:"resources"`
		} `json:"rules"`
	} `json:"items"`
}
