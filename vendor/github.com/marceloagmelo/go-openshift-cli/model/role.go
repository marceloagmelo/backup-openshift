package model

import "time"

//Role dados
type Role struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
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
}
