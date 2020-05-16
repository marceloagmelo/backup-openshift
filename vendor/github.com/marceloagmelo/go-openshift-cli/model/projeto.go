package model

//Projeto dados
type Projeto struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Labels          struct {
			Router string `json:"router"`
		} `json:"labels"`
		Annotations struct {
			OpenshiftIoDisplayName             string `json:"openshift.io/display-name"`
			OpenshiftIoNodeSelector            string `json:"openshift.io/node-selector"`
			OpenshiftIoSaSccMcs                string `json:"openshift.io/sa.scc.mcs"`
			OpenshiftIoSaSccSupplementalGroups string `json:"openshift.io/sa.scc.supplemental-groups"`
			OpenshiftIoSaSccUIDRange           string `json:"openshift.io/sa.scc.uid-range"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		Finalizers []string `json:"finalizers"`
	} `json:"spec"`
	Status struct {
		Phase string `json:"phase"`
	} `json:"status"`
}
