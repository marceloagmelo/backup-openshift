package model

//RoleBinding dados
type RoleBinding struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Annotations     struct {
			OpenshiftIoDescription string `json:"openshift.io/description"`
		} `json:"annotations"`
	} `json:"metadata"`
	UserNames  []string    `json:"userNames"`
	GroupNames interface{} `json:"groupNames"`
	Subjects   []struct {
		Kind      string `json:"kind"`
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	} `json:"subjects"`
	RoleRef struct {
		Name string `json:"name"`
	} `json:"roleRef"`
}
