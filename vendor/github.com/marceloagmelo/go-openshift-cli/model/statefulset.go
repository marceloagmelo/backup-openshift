package model

import "time"

//StateFulSet dados
type StateFulSet struct {
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
				App      string `json:"app"`
				AppName  string `json:"app_name"`
				Template string `json:"template"`
			} `json:"labels"`
		} `json:"metadata"`
		Spec struct {
			Replicas int `json:"replicas"`
			Selector struct {
				MatchLabels struct {
					App string `json:"app"`
				} `json:"matchLabels"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						App     string `json:"app"`
						AppName string `json:"app_name"`
					} `json:"labels"`
				} `json:"metadata"`
				Spec struct {
					Volumes []struct {
						Name      string `json:"name"`
						ConfigMap struct {
							Name  string `json:"name"`
							Items []struct {
								Key  string `json:"key"`
								Path string `json:"path"`
							} `json:"items"`
							DefaultMode int `json:"defaultMode"`
						} `json:"configMap,omitempty"`
						PersistentVolumeClaim struct {
							ClaimName string `json:"claimName"`
						} `json:"persistentVolumeClaim,omitempty"`
					} `json:"volumes"`
					Containers []struct {
						Name    string   `json:"name"`
						Image   string   `json:"image"`
						Command []string `json:"command"`
						Args    []string `json:"args"`
						Ports   []struct {
							Name          string `json:"name"`
							ContainerPort int    `json:"containerPort"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
						Resources struct {
						} `json:"resources"`
						VolumeMounts []struct {
							Name      string `json:"name"`
							MountPath string `json:"mountPath"`
						} `json:"volumeMounts"`
						LivenessProbe struct {
							Exec struct {
								Command []string `json:"command"`
							} `json:"exec"`
							InitialDelaySeconds int `json:"initialDelaySeconds"`
							TimeoutSeconds      int `json:"timeoutSeconds"`
							PeriodSeconds       int `json:"periodSeconds"`
							SuccessThreshold    int `json:"successThreshold"`
							FailureThreshold    int `json:"failureThreshold"`
						} `json:"livenessProbe"`
						ReadinessProbe struct {
							Exec struct {
								Command []string `json:"command"`
							} `json:"exec"`
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
					ServiceAccountName            string `json:"serviceAccountName"`
					ServiceAccount                string `json:"serviceAccount"`
					SecurityContext               struct {
					} `json:"securityContext"`
					SchedulerName string `json:"schedulerName"`
				} `json:"spec"`
			} `json:"template"`
			ServiceName         string `json:"serviceName"`
			PodManagementPolicy string `json:"podManagementPolicy"`
			UpdateStrategy      struct {
				Type string `json:"type"`
			} `json:"updateStrategy"`
			RevisionHistoryLimit int `json:"revisionHistoryLimit"`
		} `json:"spec"`
		Status struct {
			ObservedGeneration int    `json:"observedGeneration"`
			Replicas           int    `json:"replicas"`
			ReadyReplicas      int    `json:"readyReplicas"`
			CurrentReplicas    int    `json:"currentReplicas"`
			UpdatedReplicas    int    `json:"updatedReplicas"`
			CurrentRevision    string `json:"currentRevision"`
			UpdateRevision     string `json:"updateRevision"`
			CollisionCount     int    `json:"collisionCount"`
		} `json:"status"`
	} `json:"items"`
}
