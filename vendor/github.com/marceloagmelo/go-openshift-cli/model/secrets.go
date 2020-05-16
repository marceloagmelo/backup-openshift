package model

import "time"

//Secrets dados
type Secrets struct {
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
				KubernetesIoServiceAccountName string `json:"kubernetes.io/service-account.name"`
				KubernetesIoServiceAccountUID  string `json:"kubernetes.io/service-account.uid"`
				OpenshiftIoTokenSecretName     string `json:"openshift.io/token-secret.name"`
				OpenshiftIoTokenSecretValue    string `json:"openshift.io/token-secret.value"`
			} `json:"annotations"`
		} `json:"metadata,omitempty"`
		Type string `json:"type"`
	} `json:"items"`
}
