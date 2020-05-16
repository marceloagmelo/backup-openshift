package model

import "time"

//ReplicaSets dados
type ReplicaSets struct {
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
				App             string `json:"app"`
				PodTemplateHash string `json:"pod-template-hash"`
			} `json:"labels"`
			Annotations struct {
				DeploymentKubernetesIoDesiredReplicas string `json:"deployment.kubernetes.io/desired-replicas"`
				DeploymentKubernetesIoMaxReplicas     string `json:"deployment.kubernetes.io/max-replicas"`
				DeploymentKubernetesIoRevision        string `json:"deployment.kubernetes.io/revision"`
			} `json:"annotations"`
			OwnerReferences []struct {
				APIVersion         string `json:"apiVersion"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				UID                string `json:"uid"`
				Controller         bool   `json:"controller"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
			} `json:"ownerReferences"`
		} `json:"metadata"`
		Spec struct {
			Replicas int `json:"replicas"`
			Selector struct {
				MatchLabels struct {
					App             string `json:"app"`
					PodTemplateHash string `json:"pod-template-hash"`
				} `json:"matchLabels"`
			} `json:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `json:"creationTimestamp"`
					Labels            struct {
						App             string `json:"app"`
						PodTemplateHash string `json:"pod-template-hash"`
					} `json:"labels"`
				} `json:"metadata"`
				Spec struct {
					Volumes []struct {
						Name string `json:"name"`
						Nfs  struct {
							Server string `json:"server"`
							Path   string `json:"path"`
						} `json:"nfs"`
					} `json:"volumes"`
					Containers []struct {
						Name  string `json:"name"`
						Image string `json:"image"`
						Env   []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"env"`
						Resources struct {
						} `json:"resources"`
						VolumeMounts []struct {
							Name      string `json:"name"`
							MountPath string `json:"mountPath"`
						} `json:"volumeMounts"`
						TerminationMessagePath   string `json:"terminationMessagePath"`
						TerminationMessagePolicy string `json:"terminationMessagePolicy"`
						ImagePullPolicy          string `json:"imagePullPolicy"`
					} `json:"containers"`
					RestartPolicy                 string `json:"restartPolicy"`
					TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
					DNSPolicy                     string `json:"dnsPolicy"`
					NodeSelector                  struct {
						Entorno string `json:"entorno"`
						Region  string `json:"region"`
						Type    string `json:"type"`
					} `json:"nodeSelector"`
					ServiceAccountName string `json:"serviceAccountName"`
					ServiceAccount     string `json:"serviceAccount"`
					SecurityContext    struct {
					} `json:"securityContext"`
					SchedulerName string `json:"schedulerName"`
				} `json:"spec"`
			} `json:"template"`
		} `json:"spec"`
		Status struct {
			Replicas             int `json:"replicas"`
			FullyLabeledReplicas int `json:"fullyLabeledReplicas"`
			ReadyReplicas        int `json:"readyReplicas"`
			AvailableReplicas    int `json:"availableReplicas"`
			ObservedGeneration   int `json:"observedGeneration"`
		} `json:"status,omitempty"`
	} `json:"items"`
}
