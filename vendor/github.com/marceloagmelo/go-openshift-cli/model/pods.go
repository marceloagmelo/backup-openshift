package model

import "time"

//Pods dados
type Pods struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			GenerateName      string    `json:"generateName"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				App              string `json:"app"`
				AppName          string `json:"app_name"`
				Deployment       string `json:"deployment"`
				Deploymentconfig string `json:"deploymentconfig"`
			} `json:"labels"`
			Annotations struct {
				OpenshiftIoDeploymentConfigLatestVersion string `json:"openshift.io/deployment-config.latest-version"`
				OpenshiftIoDeploymentConfigName          string `json:"openshift.io/deployment-config.name"`
				OpenshiftIoDeploymentName                string `json:"openshift.io/deployment.name"`
				OpenshiftIoGeneratedBy                   string `json:"openshift.io/generated-by"`
				OpenshiftIoScc                           string `json:"openshift.io/scc"`
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
			Volumes []struct {
				Name     string `json:"name"`
				EmptyDir struct {
				} `json:"emptyDir,omitempty"`
				Secret struct {
					SecretName  string `json:"secretName"`
					DefaultMode int    `json:"defaultMode"`
				} `json:"secret,omitempty"`
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
					ReadOnly  bool   `json:"readOnly,omitempty"`
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
				SecurityContext          struct {
					Capabilities struct {
						Drop []string `json:"drop"`
					} `json:"capabilities"`
					Privileged bool `json:"privileged"`
				} `json:"securityContext"`
			} `json:"containers"`
			RestartPolicy                 string `json:"restartPolicy"`
			TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
			DNSPolicy                     string `json:"dnsPolicy"`
			NodeSelector                  struct {
				Entorno string `json:"entorno"`
				Type    string `json:"type"`
			} `json:"nodeSelector"`
			ServiceAccountName string `json:"serviceAccountName"`
			ServiceAccount     string `json:"serviceAccount"`
			NodeName           string `json:"nodeName"`
			SecurityContext    struct {
				SeLinuxOptions struct {
					Level string `json:"level"`
				} `json:"seLinuxOptions"`
				SupplementalGroups []int `json:"supplementalGroups"`
			} `json:"securityContext"`
			ImagePullSecrets []struct {
				Name string `json:"name"`
			} `json:"imagePullSecrets"`
			SchedulerName string `json:"schedulerName"`
			Tolerations   []struct {
				Key      string `json:"key"`
				Operator string `json:"operator"`
				Effect   string `json:"effect"`
			} `json:"tolerations"`
			Priority int `json:"priority"`
		} `json:"spec"`
		Status struct {
			Phase      string `json:"phase"`
			Conditions []struct {
				Type               string      `json:"type"`
				Status             string      `json:"status"`
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime time.Time   `json:"lastTransitionTime"`
			} `json:"conditions"`
			HostIP            string    `json:"hostIP"`
			PodIP             string    `json:"podIP"`
			StartTime         time.Time `json:"startTime"`
			ContainerStatuses []struct {
				Name  string `json:"name"`
				State struct {
					Running struct {
						StartedAt time.Time `json:"startedAt"`
					} `json:"running"`
				} `json:"state"`
				LastState struct {
				} `json:"lastState"`
				Ready        bool   `json:"ready"`
				RestartCount int    `json:"restartCount"`
				Image        string `json:"image"`
				ImageID      string `json:"imageID"`
				ContainerID  string `json:"containerID"`
			} `json:"containerStatuses"`
			QosClass string `json:"qosClass"`
		} `json:"status"`
	} `json:"items"`
}
