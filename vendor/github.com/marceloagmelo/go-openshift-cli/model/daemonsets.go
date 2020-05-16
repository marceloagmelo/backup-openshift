package model

import "time"

//DaemonSets dados
type DaemonSets struct {
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
			Generation        int       `json:"generation"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				App string `json:"app"`
			} `json:"labels"`
		} `json:"metadata,omitempty"`
		Spec struct {
			Selector struct {
				MatchLabels struct {
					App string `json:"app"`
				} `json:"matchLabels"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						App string `json:"app"`
					} `json:"labels"`
					Annotations struct {
						CaHash string `json:"ca_hash"`
					} `json:"annotations"`
				} `json:"metadata"`
				Spec struct {
					Volumes []struct {
						Name   string `json:"name"`
						Secret struct {
							SecretName string `json:"secretName"`
							Items      []struct {
								Key  string `json:"key"`
								Path string `json:"path"`
							} `json:"items"`
							DefaultMode int `json:"defaultMode"`
						} `json:"secret,omitempty"`
						HostPath struct {
							Path string `json:"path"`
							Type string `json:"type"`
						} `json:"hostPath,omitempty"`
						EmptyDir struct {
						} `json:"emptyDir,omitempty"`
					} `json:"volumes"`
					Containers []struct {
						Name    string   `json:"name"`
						Image   string   `json:"image"`
						Command []string `json:"command"`
						Args    []string `json:"args"`
						Ports   []struct {
							ContainerPort int    `json:"containerPort"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
						Resources struct {
						} `json:"resources"`
						VolumeMounts []struct {
							Name      string `json:"name"`
							ReadOnly  bool   `json:"readOnly"`
							MountPath string `json:"mountPath"`
						} `json:"volumeMounts"`
						LivenessProbe struct {
							HTTPGet struct {
								Path   string `json:"path"`
								Port   int    `json:"port"`
								Scheme string `json:"scheme"`
							} `json:"httpGet"`
							InitialDelaySeconds int `json:"initialDelaySeconds"`
							TimeoutSeconds      int `json:"timeoutSeconds"`
							PeriodSeconds       int `json:"periodSeconds"`
							SuccessThreshold    int `json:"successThreshold"`
							FailureThreshold    int `json:"failureThreshold"`
						} `json:"livenessProbe"`
						ReadinessProbe struct {
							HTTPGet struct {
								Path   string `json:"path"`
								Port   int    `json:"port"`
								Scheme string `json:"scheme"`
							} `json:"httpGet"`
							InitialDelaySeconds int `json:"initialDelaySeconds"`
							TimeoutSeconds      int `json:"timeoutSeconds"`
							PeriodSeconds       int `json:"periodSeconds"`
							SuccessThreshold    int `json:"successThreshold"`
							FailureThreshold    int `json:"failureThreshold"`
						} `json:"readinessProbe"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						ImagePullPolicy          string `json:"imagePullPolicy"`
					} `json:"containers"`
					RestartPolicy                 string `json:"restartPolicy"`
					TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
					DNSPolicy                     string `json:"dnsPolicy"`
					NodeSelector                  struct {
						NodeRoleKubernetesIoMaster string `json:"node-role.kubernetes.io/master"`
					} `json:"nodeSelector"`
					ServiceAccountName string `json:"serviceAccountName"`
					ServiceAccount     string `json:"serviceAccount"`
					SecurityContext    struct {
					} `json:"securityContext"`
					SchedulerName string `json:"schedulerName"`
				} `json:"spec"`
			} `json:"template"`
			TemplateGeneration   int `json:"templateGeneration"`
			RevisionHistoryLimit int `json:"revisionHistoryLimit"`
		} `json:"spec"`
		Status struct {
			CurrentNumberScheduled int `json:"currentNumberScheduled"`
			NumberMisscheduled     int `json:"numberMisscheduled"`
			DesiredNumberScheduled int `json:"desiredNumberScheduled"`
			NumberReady            int `json:"numberReady"`
			ObservedGeneration     int `json:"observedGeneration"`
			UpdatedNumberScheduled int `json:"updatedNumberScheduled"`
			NumberAvailable        int `json:"numberAvailable"`
		} `json:"status"`
	} `json:"items"`
}
