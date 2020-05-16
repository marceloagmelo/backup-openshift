package resource

import (
	"io"
)

//Resource
type (
	OpenshiftInterf interface{}
	Openshift       struct {
		Metodo      string
		Namespace   string
		NomeRecurso string
		Body        io.Reader
	}
)

//Executar
func Executar(openshift Openshift) (interface{}, int, error) {
	var resposta interface{}
	var statusCode int
	var err error

	if openshift.Metodo == "ListarBuildConfigs" {
		resposta, statusCode, err = ListarBuildConfigs(openshift.Namespace)
	} else if openshift.Metodo == "GetBuildConfig" {
		resposta, statusCode, err = GetBuildConfig(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarTemplates" {
		resposta, statusCode, err = ListarTemplates(openshift.Namespace)
	} else if openshift.Metodo == "GetTemplate" {
		resposta, statusCode, err = GetTemplate(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarServices" {
		resposta, statusCode, err = ListarServices(openshift.Namespace)
	} else if openshift.Metodo == "GetService" {
		resposta, statusCode, err = GetService(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarDeploymentConfigs" {
		resposta, statusCode, err = ListarDeploymentConfigs(openshift.Namespace)
	} else if openshift.Metodo == "GetDeploymentConfig" {
		resposta, statusCode, err = GetDeploymentConfig(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarRoleBindings" {
		resposta, statusCode, err = ListarRoleBindings(openshift.Namespace)
	} else if openshift.Metodo == "GetRoleBinding" {
		resposta, statusCode, err = GetRoleBinding(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarConfigMaps" {
		resposta, statusCode, err = ListarConfigMaps(openshift.Namespace)
	} else if openshift.Metodo == "GetConfigMap" {
		resposta, statusCode, err = GetConfigMap(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarDaemonSets" {
		resposta, statusCode, err = ListarDaemonSets(openshift.Namespace)
	} else if openshift.Metodo == "GetDaemonSet" {
		resposta, statusCode, err = GetDaemonSet(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarDeployments" {
		resposta, statusCode, err = ListarDeployments(openshift.Namespace)
	} else if openshift.Metodo == "GetDeployment" {
		resposta, statusCode, err = GetDeployment(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarImageStreams" {
		resposta, statusCode, err = ListarImageStreams(openshift.Namespace)
	} else if openshift.Metodo == "GetImageStream" {
		resposta, statusCode, err = GetImageStream(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarLimitRanges" {
		resposta, statusCode, err = ListarLimitRanges(openshift.Namespace)
	} else if openshift.Metodo == "GetLimitRange" {
		resposta, statusCode, err = GetLimitRange(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarReplicaSets" {
		resposta, statusCode, err = ListarReplicaSets(openshift.Namespace)
	} else if openshift.Metodo == "GetReplicaSet" {
		resposta, statusCode, err = GetReplicaSet(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarResourceQuotas" {
		resposta, statusCode, err = ListarResourceQuotas(openshift.Namespace)
	} else if openshift.Metodo == "GetResourceQuota" {
		resposta, statusCode, err = GetResourceQuota(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarRoles" {
		resposta, statusCode, err = ListarRoles(openshift.Namespace)
	} else if openshift.Metodo == "GetRole" {
		resposta, statusCode, err = GetRole(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarRoutes" {
		resposta, statusCode, err = ListarRoutes(openshift.Namespace)
	} else if openshift.Metodo == "GetRoute" {
		resposta, statusCode, err = GetRoute(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarSecrets" {
		resposta, statusCode, err = ListarSecrets(openshift.Namespace)
	} else if openshift.Metodo == "GetSecret" {
		resposta, statusCode, err = GetSecret(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarServiceAccounts" {
		resposta, statusCode, err = ListarServiceAccounts(openshift.Namespace)
	} else if openshift.Metodo == "GetServiceAccount" {
		resposta, statusCode, err = GetServiceAccount(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarStateFulSets" {
		resposta, statusCode, err = ListarStateFulSets(openshift.Namespace)
	} else if openshift.Metodo == "GetStateFulSet" {
		resposta, statusCode, err = GetStateFulSet(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarCronJobs" {
		resposta, statusCode, err = ListarCronJobs(openshift.Namespace)
	} else if openshift.Metodo == "GetCronJob" {
		resposta, statusCode, err = GetCronJob(openshift.Namespace, openshift.NomeRecurso)
	} else if openshift.Metodo == "ListarNamespaces" {
		resposta, statusCode, err = ListarNamespaces()
	} else if openshift.Metodo == "GetNamespace" {
		resposta, statusCode, err = GetNamespace(openshift.NomeRecurso)
	}
	return resposta, statusCode, err
}
