package resource

import (
	"net/http"

	"github.com/marceloagmelo/go-backup-openshift/api"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

// GetConfigMap recuperar
func GetConfigMap(namespace, nome string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApiV1 + "namespaces/" + namespace + "/configmaps/" + nome + "?export=true"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}

// ListarConfigMaps listar
func ListarConfigMaps(namespace string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApiV1 + "namespaces/" + namespace + "/configmaps"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
