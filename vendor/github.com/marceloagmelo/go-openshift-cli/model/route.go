package model

import "time"

//Route dados
type Route struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		UID             string `json:"uid"`
		ResourceVersion string `json:"resourceVersion"`
		Labels          struct {
			App      string `json:"app"`
			AppName  string `json:"app_name"`
			Template string `json:"template"`
		} `json:"labels"`
		Annotations struct {
			HaproxyRouterOpenshiftIoBalance string `json:"haproxy.router.openshift.io/balance"`
			OpenshiftIoHostGenerated        string `json:"openshift.io/host.generated"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		Host string `json:"host"`
		To   struct {
			Kind   string `json:"kind"`
			Name   string `json:"name"`
			Weight int    `json:"weight"`
		} `json:"to"`
		Port struct {
			TargetPort string `json:"targetPort"`
		} `json:"port"`
		TLS struct {
			Termination string `json:"termination"`
		} `json:"tls"`
		WildcardPolicy string `json:"wildcardPolicy"`
	} `json:"spec"`
	Status struct {
		Ingress []struct {
			Host       string `json:"host"`
			RouterName string `json:"routerName"`
			Conditions []struct {
				Type               string    `json:"type"`
				Status             string    `json:"status"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
			} `json:"conditions"`
			WildcardPolicy string `json:"wildcardPolicy"`
		} `json:"ingress"`
	} `json:"status"`
}
