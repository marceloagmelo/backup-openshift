package resource

import (
	"net/http"

	"github.com/marceloagmelo/go-backup-openshift/api"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

// GetTemplate recuperar Template
func GetTemplate(namespace, nome string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApiTemplates + "namespaces/" + namespace + "/templates/" + nome + "?export=true"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}

// ListarTemplates listar todos Templates
func ListarTemplates(namespace string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApiTemplates + "namespaces/" + namespace + "/templates"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
