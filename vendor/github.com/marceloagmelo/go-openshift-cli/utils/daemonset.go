package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetDaemonSet recuperar DaemonSet
func GetDaemonSet(token string, url string, projeto string, nome string) (resultado int, daemonSet model.DaemonSet) {
	resultado = 0
	endpoint := url + apisExtensionsV1beta1 + "namespaces/" + projeto + "/daemonsets/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDaemonSet] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, daemonSet
		}
		daemonSet = model.DaemonSet{}
		err = json.Unmarshal(corpo, &daemonSet)
		if err != nil {
			fmt.Println("[GetDaemonSet] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, daemonSet
}

// GetDaemonSetString recuperar DaemonSet
func GetDaemonSetString(token string, url string, projeto string, nome string) (resultado int, daemonSetString string) {
	resultado = 0
	endpoint := url + apisExtensionsV1beta1 + "namespaces/" + projeto + "/daemonsets/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDaemonSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		daemonSetString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, daemonSetString
}

// ListDaemonSet listar todos DaemonSets
func ListDaemonSet(token string, url string) (resultado int, daemonset model.DaemonSets) {
	resultado = 0
	endpoint := url + apisExtensionsV1beta1 + "daemonsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDaemonSet] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, daemonset
		}
		daemonset = model.DaemonSets{}
		err = json.Unmarshal(corpo, &daemonset)
		if err != nil {
			fmt.Println("[ListDaemonSet] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, daemonset
}

// ListDaemonSetString listar todos DaemonSets
func ListDaemonSetString(token string, url string) (resultado int, daemonsetString string) {
	resultado = 0
	endpoint := url + apisExtensionsV1beta1 + "daemonsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDaemonSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		daemonsetString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, daemonsetString
}

// ListDaemonSetProjeto listar DaemonSets por projetos
func ListDaemonSetProjeto(token string, url string, projeto string) (resultado int, daemonset model.DaemonSets) {
	resultado = 0
	endpoint := url + apisExtensionsV1beta1 + "namespaces/" + projeto + "/daemonsets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDaemonSetProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, daemonset
		}
		daemonset = model.DaemonSets{}
		err = json.Unmarshal(corpo, &daemonset)
		if err != nil {
			fmt.Println("[ListDaemonSetProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, daemonset
}

// CriarDeployment criar um DaemonSets
func CriarDaemonSet(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apisExtensionsV1beta1 + "namespaces/" + projeto + "/daemonsets"

	resultado, resposta := PostRequest(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
