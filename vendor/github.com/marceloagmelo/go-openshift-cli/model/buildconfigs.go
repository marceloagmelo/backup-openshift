package model

import "time"

//BuildConfigs dados
type BuildConfigs struct {
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
			Labels            struct {
				App      string `json:"app"`
				AppName  string `json:"app_name"`
				Template string `json:"template"`
			} `json:"labels"`
			Annotations struct {
				Description string `json:"description"`
			} `json:"annotations"`
		} `json:"metadata"`
		Spec struct {
			Triggers []struct {
				Type   string `json:"type"`
				Github struct {
					Secret string `json:"secret"`
				} `json:"github,omitempty"`
				Generic struct {
					Secret string `json:"secret"`
				} `json:"generic,omitempty"`
			} `json:"triggers"`
			RunPolicy string `json:"runPolicy"`
			Source    struct {
				Type string `json:"type"`
				Git  struct {
					URI string `json:"uri"`
				} `json:"git"`
			} `json:"source"`
			Strategy struct {
				Type           string `json:"type"`
				DockerStrategy struct {
					From struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
					} `json:"from"`
					ForcePull bool `json:"forcePull"`
				} `json:"dockerStrategy"`
			} `json:"strategy"`
			Output struct {
				To struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
				} `json:"to"`
				PushSecret struct {
					Name string `json:"name"`
				} `json:"pushSecret"`
			} `json:"output"`
			Resources struct {
			} `json:"resources"`
			PostCommit struct {
			} `json:"postCommit"`
			NodeSelector                 interface{} `json:"nodeSelector"`
			SuccessfulBuildsHistoryLimit int         `json:"successfulBuildsHistoryLimit"`
			FailedBuildsHistoryLimit     int         `json:"failedBuildsHistoryLimit"`
		} `json:"spec,omitempty"`
		Status struct {
			LastVersion int `json:"lastVersion"`
		} `json:"status"`
	} `json:"items"`
}
