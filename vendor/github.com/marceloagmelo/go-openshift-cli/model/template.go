package model

//Template dados
type Template struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Annotations     struct {
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
		Description string `json:"description,omitempty"`
		Required    bool   `json:"required,omitempty"`
		Value       string `json:"value,omitempty"`
	} `json:"parameters"`
	Labels struct {
		Template string `json:"template"`
	} `json:"labels"`
}
