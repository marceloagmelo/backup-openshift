package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// ListPod listar todos pods
func ListPod(token string, url string) (resultado int, pods model.Pods) {
	resultado = 0
	endpoint := url + apiPods
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPod] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, pods
		}
		pods = model.Pods{}
		err = json.Unmarshal(corpo, &pods)
		if err != nil {
			fmt.Println("[ListPod] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pods
}

// ListPodProjeto listar pod por projetos
func ListPodProjeto(token string, url string, projeto string) (resultado int, pods model.Pods) {
	resultado = 0
	endpoint := url + apiPods + "namespaces/" + projeto + "/pods"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPodProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, pods
		}
		pods = model.Pods{}
		err = json.Unmarshal(corpo, &pods)
		if err != nil {
			fmt.Println("[ListPodProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pods
}
