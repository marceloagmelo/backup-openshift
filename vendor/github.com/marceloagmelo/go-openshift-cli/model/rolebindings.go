package model

import "time"

//RoleBindings dados
type RoleBindings struct {
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
		} `json:"metadata,omitempty"`
		UserNames  []string    `json:"userNames"`
		GroupNames interface{} `json:"groupNames"`
		Subjects   []struct {
			Kind      string `json:"kind"`
			Namespace string `json:"namespace"`
			Name      string `json:"name"`
		} `json:"subjects"`
		RoleRef struct {
			Namespace string `json:"namespace"`
			Name      string `json:"name"`
		} `json:"roleRef,omitempty"`
	} `json:"items"`
}
