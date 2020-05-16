package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetPvc recuperar PVC
func GetPvc(token string, url string, projeto string, nome string) (resultado int, pvc model.Pvc) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/persistentvolumeclaims/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetPvc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, pvc
		}
		pvc = model.Pvc{}
		err = json.Unmarshal(corpo, &pvc)
		if err != nil {
			fmt.Println("[GetPvc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pvc
}

// GetPvcString recuperar PVC
func GetPvcString(token string, url string, projeto string, nome string) (resultado int, pvcString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/persistentvolumeclaims/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetPvcString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, pvcString
}

// ListPvc listar todos PVC
func ListPvc(token string, url string) (resultado int, pvcs model.Pvcs) {
	resultado = 0
	endpoint := url + apiV1 + "persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvc] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, pvcs
		}
		pvcs = model.Pvcs{}
		err = json.Unmarshal(corpo, &pvcs)
		if err != nil {
			fmt.Println("[ListPvc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pvcs
}

// ListPvcString listar todos PVC
func ListPvcString(token string, url string) (resultado int, pvcsString string) {
	resultado = 0
	endpoint := url + apiV1 + "persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvcString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, pvcsString
}

// ListPvcProjeto listar PVC por projetos
func ListPvcProjeto(token string, url string, projeto string) (resultado int, pvcs model.Pvcs) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvcProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, pvcs
		}
		pvcs = model.Pvcs{}
		err = json.Unmarshal(corpo, &pvcs)
		if err != nil {
			fmt.Println("[ListPvcProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, pvcs
}

// ListPvcProjetoString listar PVC por projetos
func ListPvcProjetoString(token string, url string, projeto string) (resultado int, pvcsString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/persistentvolumeclaims"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListPvcString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		pvcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, pvcsString
}

// CriarPvc criar um PVC
func CriarPvc(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/persistentvolumeclaims"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarPvc] Erro ao executar o comando no OS.")
		erro = "[CriarPvc] Erro ao executar o comando no OS. " + cmd
		return resultado, erro
	}

	cmd = "sed -i s/\\\"volumeName[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarPvc] Erro ao executar o comando no OS.")
		erro = "[CriarPvc] Erro ao executar o comando no OS. " + cmd
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
