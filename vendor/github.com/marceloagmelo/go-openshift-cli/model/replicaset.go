package model

//ReplicaSet dados
type ReplicaSet struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Generation      int    `json:"generation"`
		Labels          struct {
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
				Name   string `json:"name"`
				Labels struct {
					App             string `json:"app"`
					PodTemplateHash string `json:"pod-template-hash"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Volumes []struct {
					Name      string `json:"name"`
					ConfigMap struct {
						Name        string `json:"name"`
						DefaultMode int    `json:"defaultMode"`
					} `json:"configMap,omitempty"`
					Secret struct {
						SecretName  string `json:"secretName"`
						DefaultMode int    `json:"defaultMode"`
					} `json:"secret,omitempty"`
					PersistentVolumeClaim struct {
						ClaimName string `json:"claimName"`
					} `json:"persistentVolumeClaim,omitempty"`
				} `json:"volumes"`
				Containers []struct {
					Name  string   `json:"name"`
					Image string   `json:"image"`
					Args  []string `json:"args,omitempty"`
					Ports []struct {
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
					TerminationMessagePath   string   `json:"terminationMessagePath"`
					TerminationMessagePolicy string   `json:"terminationMessagePolicy"`
					ImagePullPolicy          string   `json:"imagePullPolicy"`
					Command                  []string `json:"command,omitempty"`
					EnvFrom                  []struct {
						ConfigMapRef struct {
							Name string `json:"name"`
						} `json:"configMapRef"`
					} `json:"envFrom,omitempty"`
				} `json:"containers"`
				RestartPolicy                 string `json:"restartPolicy"`
				TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
				DNSPolicy                     string `json:"dnsPolicy"`
				NodeSelector                  struct {
					App     string `json:"app"`
					Entorno string `json:"entorno"`
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
		Replicas           int `json:"replicas"`
		ObservedGeneration int `json:"observedGeneration"`
	} `json:"status"`
}
