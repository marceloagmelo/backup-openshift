package resource

import (
	"net/http"

	"github.com/marceloagmelo/go-backup-openshift/api"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

// GetStateFulSet recuperar
func GetStateFulSet(namespace, nome string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApisAppsv1beta1 + "namespaces/" + namespace + "/statefulsets/" + nome + "?export=true"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}

// ListarStateFulSets listar
func ListarStateFulSets(namespace string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApisAppsv1beta1 + "namespaces/" + namespace + "/statefulsets"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
