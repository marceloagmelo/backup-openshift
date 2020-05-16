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
				RevisionHistoryLimit int `json:"revisionHistoryLimit"`
				Selector             struct {
					Deploymentconfig string `json:"deploymentconfig"`
				} `json:"selector"`
				Template struct {
					Metadata struct {
						Labels struct {
							Deploymentconfig string `json:"deploymentconfig"`
						} `json:"labels"`
					} `json:"metadata"`
					Spec struct {
						Containers []struct {
							Env []struct {
								Name      string `json:"name"`
								Value     string `json:"value,omitempty"`
								ValueFrom struct {
									FieldRef struct {
										FieldPath string `json:"fieldPath"`
									} `json:"fieldRef"`
								} `json:"valueFrom,omitempty"`
							} `json:"env"`
							Image           string `json:"image"`
							ImagePullPolicy string `json:"imagePullPolicy"`
							LivenessProbe   struct {
								HTTPGet struct {
									Path string `json:"path"`
									Port int    `json:"port"`
								} `json:"httpGet"`
								InitialDelaySeconds int `json:"initialDelaySeconds"`
								TimeoutSeconds      int `json:"timeoutSeconds"`
							} `json:"livenessProbe"`
							Name  string `json:"name"`
							Ports []struct {
								ContainerPort int    `json:"containerPort"`
								Protocol      string `json:"protocol"`
							} `json:"ports"`
							ReadinessProbe struct {
								HTTPGet struct {
									Path string `json:"path"`
									Port int    `json:"port"`
								} `json:"httpGet"`
								InitialDelaySeconds int `json:"initialDelaySeconds"`
								TimeoutSeconds      int `json:"timeoutSeconds"`
							} `json:"readinessProbe"`
							TerminationMessagePath string `json:"terminationMessagePath"`
							VolumeMounts           []struct {
								MountPath string `json:"mountPath"`
								Name      string `json:"name"`
								ReadOnly  bool   `json:"readOnly"`
							} `json:"volumeMounts"`
						} `json:"containers"`
						DNSPolicy     string `json:"dnsPolicy"`
						RestartPolicy string `json:"restartPolicy"`
						Volumes       []struct {
							Metadata struct {
								Items []struct {
									FieldRef struct {
										FieldPath string `json:"fieldPath"`
									} `json:"fieldRef"`
									Name string `json:"name"`
								} `json:"items"`
							} `json:"metadata"`
							Name string `json:"name"`
						} `json:"volumes"`
					} `json:"spec"`
				} `json:"template"`
				Triggers []struct {
					Type string `json:"type"`
				} `json:"triggers"`
			} `json:"spec,omitempty"`
		} `json:"objects"`
		Parameters []struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			Required    bool   `json:"required"`
		} `json:"parameters"`
		Labels struct {
			Template string `json:"template"`
		} `json:"labels,omitempty"`
	} `json:"items"`
}
