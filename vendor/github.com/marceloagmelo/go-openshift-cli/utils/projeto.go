package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetProjeto recuperar projeto
func GetProjeto(token string, url string, nome string) (resultado int, proj model.Projeto) {
	resultado = 0
	endpoint := url + apiProjeto + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, proj
		}
		proj = model.Projeto{}
		err = json.Unmarshal(corpo, &proj)
		if err != nil {
			fmt.Println("[GetProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, proj
}

// GetProjetoString recuperar projeto
func GetProjetoString(token string, url string, nome string) (resultado int, projString string) {
	resultado = 0
	endpoint := url + apiProjeto + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetProjetoString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		projString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, projString
}

// Projetos listar projetos
func Projetos(token string, url string) (resultado int, projetos model.Projetos) {
	resultado = 0
	endpoint := url + apiProjeto

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[Projetos] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, projetos
		}
		projetos = model.Projetos{}
		err = json.Unmarshal(corpo, &projetos)
		if err != nil {
			fmt.Println("[Projetos] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, projetos
}

// GetNamespace recuperar namespace
func GetNamespace(token string, url string, nome string) (resultado int, proj model.Projeto) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetNamespace] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, proj
		}
		proj = model.Projeto{}
		err = json.Unmarshal(corpo, &proj)
		if err != nil {
			fmt.Println("[GetNamespace] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, proj
}

// GetNamespaceString recuperar namespace
func GetNamespaceString(token string, url string, nome string) (resultado int, projString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetNamespaceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		projString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, projString
}

// Namespaces listar namespaces
func Namespaces(token string, url string) (resultado int, projetos model.Projetos) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/"

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[Namespaces] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, projetos
		}
		projetos = model.Projetos{}
		err = json.Unmarshal(corpo, &projetos)
		if err != nil {
			fmt.Println("[Namespaces] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, projetos
}

// CriarNamespace criar um namespace
func CriarNamespace(token string, url string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/"

	resultado, resposta := PostRequest(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
