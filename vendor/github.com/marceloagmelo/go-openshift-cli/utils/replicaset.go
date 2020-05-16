package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetReplicaSet recuperar ReplicaSet
func GetReplicaSet(token string, url string, projeto string, nome string) (resultado int, replicaset model.ReplicaSet) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/replicasets/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetReplicaSet] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, replicaset
		}
		replicaset = model.ReplicaSet{}
		err = json.Unmarshal(corpo, &replicaset)
		if err != nil {
			fmt.Println("[GetReplicaSet] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, replicaset
}

// GetReplicaSetString recuperar ReplicaSet
func GetReplicaSetString(token string, url string, projeto string, nome string) (resultado int, replicasetString string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/replicasets/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetReplicaSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		replicasetString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, replicasetString
}

// ListReplicaSet listar todos ReplicaSet
func ListReplicaSet(token string, url string) (resultado int, replicasets model.ReplicaSets) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "replicasets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListReplicaSet] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, replicasets
		}
		replicasets = model.ReplicaSets{}
		err = json.Unmarshal(corpo, &replicasets)
		if err != nil {
			fmt.Println("[ListReplicaSet] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, replicasets
}

// ListReplicaSetString listar todos ReplicaSet
func ListReplicaSetString(token string, url string) (resultado int, replicasetsString string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "replicasets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListReplicaSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		replicasetsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, replicasetsString
}

// ListReplicaSetProjeto listar ReplicaSet por projetos
func ListReplicaSetProjeto(token string, url string, projeto string) (resultado int, replicasets model.ReplicaSets) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/replicasets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListReplicaSetProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, replicasets
		}
		replicasets = model.ReplicaSets{}
		err = json.Unmarshal(corpo, &replicasets)
		if err != nil {
			fmt.Println("[ListReplicaSetProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, replicasets
}

// ListReplicaSetProjetoString listar ReplicaSet por projetos
func ListReplicaSetProjetoString(token string, url string, projeto string) (resultado int, replicasetsString string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/replicasets"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListReplicaSetString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		replicasetsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, replicasetsString
}

// CriarReplicaSet criar um replicaset
func CriarReplicaSet(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/replicasets"

	/*cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarServiceAccount] Erro ao executar o comando no OS.")
		erro = "[CriarServiceAccount] Erro ao executar o comando no OS."
		return resultado, erro
	}*/

	resultado, resposta := PostRequestFile(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
