package model

import "time"

//Deployments dados
type Deployments struct {
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
				App       string `json:"app"`
				Component string `json:"component"`
			} `json:"labels"`
			Annotations struct {
				DeploymentKubernetesIoRevision              string `json:"deployment.kubernetes.io/revision"`
				KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
			} `json:"annotations"`
		} `json:"metadata"`
		Spec struct {
			Replicas int `json:"replicas"`
			Selector struct {
				MatchLabels struct {
					App       string `json:"app"`
					Component string `json:"component"`
				} `json:"matchLabels"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					Name              string      `json:"name"`
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						App       string `json:"app"`
						Component string `json:"component"`
					} `json:"labels"`
					Annotations struct {
						InstallerTriggeredRollout string `json:"installer-triggered-rollout"`
					} `json:"annotations"`
				} `json:"metadata"`
				Spec struct {
					Volumes []struct {
						Name   string `json:"name"`
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
						Name    string   `json:"name"`
						Image   string   `json:"image"`
						Command []string `json:"command"`
						Ports   []struct {
							ContainerPort int    `json:"containerPort"`
							Protocol      string `json:"protocol"`
						} `json:"ports"`
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
							ReadOnly  bool   `json:"readOnly,omitempty"`
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
							TimeoutSeconds   int `json:"timeoutSeconds"`
							PeriodSeconds    int `json:"periodSeconds"`
							SuccessThreshold int `json:"successThreshold"`
							FailureThreshold int `json:"failureThreshold"`
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
					SecurityContext struct {
					} `json:"securityContext"`
					Affinity struct {
						PodAntiAffinity struct {
							PreferredDuringSchedulingIgnoredDuringExecution []struct {
								Weight          int `json:"weight"`
								PodAffinityTerm struct {
									LabelSelector struct {
										MatchLabels struct {
											App string `json:"app"`
										} `json:"matchLabels"`
									} `json:"labelSelector"`
									TopologyKey string `json:"topologyKey"`
								} `json:"podAffinityTerm"`
							} `json:"preferredDuringSchedulingIgnoredDuringExecution"`
						} `json:"podAntiAffinity"`
					} `json:"affinity"`
					SchedulerName string `json:"schedulerName"`
				} `json:"spec"`
			} `json:"template"`
			Strategy struct {
				Type          string `json:"type"`
				RollingUpdate struct {
					MaxUnavailable string `json:"maxUnavailable"`
					MaxSurge       string `json:"maxSurge"`
				} `json:"rollingUpdate"`
			} `json:"strategy"`
			RevisionHistoryLimit    int `json:"revisionHistoryLimit"`
			ProgressDeadlineSeconds int `json:"progressDeadlineSeconds"`
		} `json:"spec"`
		Status struct {
			ObservedGeneration int `json:"observedGeneration"`
			Replicas           int `json:"replicas"`
			UpdatedReplicas    int `json:"updatedReplicas"`
			ReadyReplicas      int `json:"readyReplicas"`
			AvailableReplicas  int `json:"availableReplicas"`
			Conditions         []struct {
				Type               string    `json:"type"`
				Status             string    `json:"status"`
				LastUpdateTime     time.Time `json:"lastUpdateTime"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				Reason             string    `json:"reason"`
				Message            string    `json:"message"`
			} `json:"conditions"`
		} `json:"status"`
	} `json:"items"`
}
