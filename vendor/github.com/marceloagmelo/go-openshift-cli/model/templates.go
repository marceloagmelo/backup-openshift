package model

import "time"

//Templates dados
type Templates struct {
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
				Description                                 string `json:"description"`
				IconClass                                   string `json:"iconClass"`
				KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
				OpenshiftIoDisplayName                      string `json:"openshift.io/display-name"`
				OpenshiftIoDocumentationURL                 string `json:"openshift.io/documentation-url"`
				OpenshiftIoProviderDisplayName              string `json:"openshift.io/provider-display-name"`
				Tags                                        string `json:"tags"`
			} `json:"annotations"`
		} `json:"metadata"`
		Objects []struct {
			APIVersion string `json:"apiVersion"`
			Kind       string `json:"kind"`
			Metadata   struct {
				Labels struct {
					AppName string `json:"app_name"`
				} `json:"labels"`
				Name string `json:"name"`
			} `json:"metadata,omitempty"`
			Spec struct {
				Ports []struct {
					Name       string `json:"name"`
					NodePort   int    `json:"nodePort"`
					Port       int    `json:"port"`
					Protocol   string `json:"protocol"`
					TargetPort int    `json:"targetPort"`
				} `json:"ports"`
				Selector struct {
					Deploymentconfig string `json:"deploymentconfig"`
				} `json:"selector"`
				SessionAffinity string `json:"sessionAffinity"`
				Type            string `json:"type"`
			} `json:"spec,omitempty"`
			Status struct {
				LoadBalancer struct {
				} `json:"loadBalancer"`
			} `json:"status,omitempty"`
		} `json:"objects"`
		Parameters []struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			Value       string `json:"value,omitempty"`
			Required    bool   `json:"required,omitempty"`
		} `json:"parameters"`
		Labels struct {
			Template string `json:"template"`
		} `json:"labels,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"items"`
}
