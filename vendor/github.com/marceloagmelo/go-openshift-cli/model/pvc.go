package model

//Pvc dados
type Pvc struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Labels          struct {
			Template string `json:"template"`
		} `json:"labels"`
		Annotations struct {
			PvKubernetesIoBindCompleted     string `json:"pv.kubernetes.io/bind-completed"`
			PvKubernetesIoBoundByController string `json:"pv.kubernetes.io/bound-by-controller"`
		} `json:"annotations"`
		Finalizers []string `json:"finalizers"`
	} `json:"metadata"`
	Spec struct {
		AccessModes []string `json:"accessModes"`
		Resources   struct {
			Requests struct {
				Storage string `json:"storage"`
			} `json:"requests"`
		} `json:"resources"`
		VolumeName string `json:"volumeName"`
	} `json:"spec"`
	Status struct {
		Phase       string   `json:"phase"`
		AccessModes []string `json:"accessModes"`
		Capacity    struct {
			Storage string `json:"storage"`
		} `json:"capacity"`
	} `json:"status"`
}
