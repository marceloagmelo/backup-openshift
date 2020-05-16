package model

//BuildConfig dados
type BuildConfig struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Labels          struct {
			App string `json:"app"`
		} `json:"labels"`
		Annotations struct {
			OpenshiftIoGeneratedBy string `json:"openshift.io/generated-by"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		Triggers []struct {
			Type        string `json:"type"`
			ImageChange struct {
				LastTriggeredImageID string `json:"lastTriggeredImageID"`
			} `json:"imageChange,omitempty"`
			Generic struct {
				Secret string `json:"secret"`
			} `json:"generic,omitempty"`
			Github struct {
				Secret string `json:"secret"`
			} `json:"github,omitempty"`
		} `json:"triggers"`
		RunPolicy string `json:"runPolicy"`
		Source    struct {
			Type string `json:"type"`
			Git  struct {
				URI string `json:"uri"`
				Ref string `json:"ref"`
			} `json:"git"`
		} `json:"source"`
		Strategy struct {
			Type           string `json:"type"`
			SourceStrategy struct {
				From struct {
					Kind      string `json:"kind"`
					Namespace string `json:"namespace"`
					Name      string `json:"name"`
				} `json:"from"`
			} `json:"sourceStrategy"`
		} `json:"strategy"`
		Output struct {
			To struct {
				Kind string `json:"kind"`
				Name string `json:"name"`
			} `json:"to"`
		} `json:"output"`
		Resources struct {
		} `json:"resources"`
		PostCommit struct {
		} `json:"postCommit"`
		NodeSelector interface{} `json:"nodeSelector"`
	} `json:"spec"`
	Status struct {
		LastVersion int `json:"lastVersion"`
	} `json:"status"`
}
