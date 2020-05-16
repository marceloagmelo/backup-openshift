package resource

import (
	"net/http"

	"github.com/marceloagmelo/go-backup-openshift/api"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

// GetImageStream recuperar
func GetImageStream(namespace, nome string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApisImageV1 + "namespaces/" + namespace + "/imagestreams/" + nome + "?export=true"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}

// ListarImageStreams listar
func ListarImageStreams(namespace string) (interface{}, int, error) {
	var interf interface{}

	endpoint := variaveis.OpenshiftApiURL + variaveis.OpenshiftApisImageV1 + "namespaces/" + namespace + "/imagestreams"

	interf, statusCode, err := api.Recuperar(endpoint)
	if err != nil {
		return interf, http.StatusInternalServerError, err
	}

	return interf, statusCode, nil
}
