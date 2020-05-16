package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetService recuperar SVC
func GetService(token string, url string, projeto string, nome string) (resultado int, service model.Service) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetService] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, service
		}
		service = model.Service{}
		err = json.Unmarshal(corpo, &service)
		if err != nil {
			fmt.Println("[GetService] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, service
}

// GetServiceString recuperar SVC
func GetServiceString(token string, url string, projeto string, nome string) (resultado int, serviceString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetServiceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceString
}

// ListService listar todos SVC
func ListService(token string, url string) (resultado int, services model.Services) {
	resultado = 0
	endpoint := url + apiV1 + "services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListService] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, services
		}
		services = model.Services{}
		err = json.Unmarshal(corpo, &services)
		if err != nil {
			fmt.Println("[ListService] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, services
}

// ListServiceString listar todos SVC
func ListServiceString(token string, url string) (resultado int, servicesString string) {
	resultado = 0
	endpoint := url + apiV1 + "services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, servicesString
		}
		servicesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, servicesString
}

// ListServiceProjeto listar SVC por projetos
func ListServiceProjeto(token string, url string, projeto string) (resultado int, services model.Services) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, services
		}
		services = model.Services{}
		err = json.Unmarshal(corpo, &services)
		if err != nil {
			fmt.Println("[ListServiceProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, services
}

// ListServiceProjetoString listar SVC por projetos
func ListServiceProjetoString(token string, url string, projeto string) (resultado int, servicesString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		servicesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, servicesString
}

// CriarService criar um SVC
func CriarService(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/services"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarService] Erro ao executar o comando no OS.")
		erro = "[CriarService] Erro ao executar o comando no OS."
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
