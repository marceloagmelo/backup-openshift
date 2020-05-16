package model

import "time"

//DeploymentConfigs dados
type DeploymentConfigs struct {
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
				App        string `json:"app"`
				AppName    string `json:"app_name"`
				Technology string `json:"technology"`
			} `json:"labels"`
			Annotations struct {
				OpenshiftIoGeneratedBy string `json:"openshift.io/generated-by"`
			} `json:"annotations"`
		} `json:"metadata"`
		Spec struct {
			Strategy struct {
				Type      string `json:"type"`
				Resources struct {
				} `json:"resources"`
				ActiveDeadlineSeconds int `json:"activeDeadlineSeconds"`
			} `json:"strategy"`
			Triggers             []interface{} `json:"triggers"`
			Replicas             int           `json:"replicas"`
			RevisionHistoryLimit int           `json:"revisionHistoryLimit"`
			Test                 bool          `json:"test"`
			Selector             struct {
				AppName          string `json:"app_name"`
				Deploymentconfig string `json:"deploymentconfig"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						App              string `json:"app"`
						AppName          string `json:"app_name"`
						Deploymentconfig string `json:"deploymentconfig"`
					} `json:"labels"`
					Annotations struct {
						OpenshiftIoGeneratedBy string `json:"openshift.io/generated-by"`
					} `json:"annotations"`
				} `json:"metadata"`
				Spec struct {
					Volumes []struct {
						Name     string `json:"name"`
						EmptyDir struct {
						} `json:"emptyDir,omitempty"`
						Secret struct {
							SecretName  string `json:"secretName"`
							DefaultMode int    `json:"defaultMode"`
						} `json:"secret,omitempty"`
						ConfigMap struct {
							Name        string `json:"name"`
							DefaultMode int    `json:"defaultMode"`
						} `json:"configMap,omitempty"`
					} `json:"volumes"`
					Containers []struct {
						Name  string `json:"name"`
						Image string `json:"image"`
						Ports []struct {
							ContainerPort int    `json:"containerPort"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
						Env []struct {
							Name      string `json:"name"`
							Value     string `json:"value,omitempty"`
							ValueFrom struct {
								FieldRef struct {
									APIVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
							} `json:"valueFrom,omitempty"`
						} `json:"env"`
						Resources struct {
							Limits struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"limits"`
							Requests struct {
								CPU    string `json:"cpu"`
								Memory string `json:"memory"`
							} `json:"requests"`
						} `json:"resources"`
						VolumeMounts []struct {
							Name      string `json:"name"`
							MountPath string `json:"mountPath"`
						} `json:"volumeMounts"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						ImagePullPolicy          string `json:"imagePullPolicy"`
						SecurityContext          struct {
							Capabilities struct {
							} `json:"capabilities"`
							Privileged bool `json:"privileged"`
						} `json:"securityContext"`
					} `json:"containers"`
					RestartPolicy                 string `json:"restartPolicy"`
					TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
					DNSPolicy                     string `json:"dnsPolicy"`
					SecurityContext               struct {
						SupplementalGroups []int `json:"supplementalGroups"`
					} `json:"securityContext"`
					SchedulerName string `json:"schedulerName"`
				} `json:"spec"`
			} `json:"template"`
		} `json:"spec"`
		Status struct {
			LatestVersion       int `json:"latestVersion"`
			ObservedGeneration  int `json:"observedGeneration"`
			Replicas            int `json:"replicas"`
			UpdatedReplicas     int `json:"updatedReplicas"`
			AvailableReplicas   int `json:"availableReplicas"`
			UnavailableReplicas int `json:"unavailableReplicas"`
			Details             struct {
				Message string `json:"message"`
				Causes  []struct {
					Type string `json:"type"`
				} `json:"causes"`
			} `json:"details"`
			Conditions []struct {
				Type               string    `json:"type"`
				Status             string    `json:"status"`
				LastUpdateTime     time.Time `json:"lastUpdateTime"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				Message            string    `json:"message"`
				Reason             string    `json:"reason,omitempty"`
			} `json:"conditions"`
		} `json:"status"`
	} `json:"items"`
}
