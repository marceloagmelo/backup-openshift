package resource

import (
	"net/http"

	"github.com/marceloagmelo/go-backup-openshift/api"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

// GetRoute recuperar
func GetRoute(namespace, nome string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApiRoutes + "namespaces/" + namespace + "/routes/" + nome + "?export=true"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}

// ListarRoutes listar
func ListarRoutes(namespace string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApiRoutes + "namespaces/" + namespace + "/routes"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
