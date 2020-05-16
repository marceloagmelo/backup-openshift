package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetConfigMap recuperar ConfigMap
func GetConfigMap(token string, url string, projeto string, nome string) (resultado int, configMap model.ConfigMap) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/configmaps/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetConfigMap] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, configMap
		}
		configMap = model.ConfigMap{}
		err = json.Unmarshal(corpo, &configMap)
		if err != nil {
			fmt.Println("[GetConfigMap] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, configMap
}

// GetConfigMapString recuperar ConfigMap
func GetConfigMapString(token string, url string, projeto string, nome string) (resultado int, configMapString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/configmaps/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetConfigMapString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configMapString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, configMapString
}

// ListConfigMap listar todos ConfigMaps
func ListConfigMap(token string, url string) (resultado int, configmap model.ConfigMaps) {
	resultado = 0
	endpoint := url + apiV1 + "configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMap] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, configmap
		}
		configmap = model.ConfigMaps{}
		err = json.Unmarshal(corpo, &configmap)
		if err != nil {
			fmt.Println("[ListConfigMap] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, configmap
}

// ListConfigMapString listar todos ConfigMaps
func ListConfigMapString(token string, url string) (resultado int, configmapString string) {
	resultado = 0
	endpoint := url + apiV1 + "configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMapString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configmapString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, configmapString
}

// ListConfigMapProjeto listar ConfigMaps por projetos
func ListConfigMapProjeto(token string, url string, projeto string) (resultado int, configmap model.ConfigMaps) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMapProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, configmap
		}
		configmap = model.ConfigMaps{}
		err = json.Unmarshal(corpo, &configmap)
		if err != nil {
			fmt.Println("[ListConfigMapProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, configmap
}

// ListConfigMapProjetoString listar ConfigMaps por projetos
func ListConfigMapProjetoString(token string, url string, projeto string) (resultado int, configmapString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/configmaps"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListConfigMapString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		configmapString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, configmapString
}

// CriarConfigMap criar um config map
func CriarConfigMap(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/configmaps"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarConfigMap] Erro ao executar o comando no OS.")
		erro = "[CriarConfigMap] Erro ao executar o comando no OS."
		return resultado, erro
	}

	resultado, resposta := PostRequestFile(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
