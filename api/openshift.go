package api

import (
	"fmt"
	"strings"

	"github.com/marceloagmelo/go-backup-openshift/logger"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

//Listar
func Listar(endPoint string) (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderOpenshift("application/json")
	apiRequest.EndPoint = endPoint
	apiRequest.Metodo = "GET"

	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

//Recuperar
func Recuperar(endPoint string) (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderOpenshift("application/json")
	apiRequest.EndPoint = endPoint
	apiRequest.Metodo = "GET"
	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao recuperar um recurso", err.Error())
		logger.Erro.Println(mensagem)
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

//getHeader
func montarHeaderOpenshift(contentType string) []Header {
	var headers []Header
	header := Header{}

	var bearerAuth = "Bearer " + strings.TrimSuffix(variaveis.OpenshiftToken, "\n")
	header.Chave = "Authorization"
	header.Valor = bearerAuth
	headers = append(headers, header)

	header.Chave = "Accept"
	header.Valor = "application/json"
	headers = append(headers, header)

	header.Chave = "Content-Type"
	header.Valor = contentType
	headers = append(headers, header)

	return headers
}
