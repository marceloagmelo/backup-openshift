package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetDeployment recuperar Deployment
func GetDeployment(token string, url string, projeto string, nome string) (resultado int, deployment model.Deployment) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/deployments/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDeployment] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, deployment
		}
		deployment = model.Deployment{}
		err = json.Unmarshal(corpo, &deployment)
		if err != nil {
			fmt.Println("[GetDeployment] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, deployment
}

// GetDeploymentString recuperar Deployment
func GetDeploymentString(token string, url string, projeto string, nome string) (resultado int, deploymentString string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/deployments/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDeploymentString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		deploymentString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, deploymentString
}

// ListDeployment listar todos Deployment
func ListDeployment(token string, url string) (resultado int, deployments model.Deployments) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "deployments"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeployment] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, deployments
		}
		deployments = model.Deployments{}
		err = json.Unmarshal(corpo, &deployments)
		if err != nil {
			fmt.Println("[ListDeployment] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, deployments
}

// ListDeploymentString listar todos Deployment
func ListDeploymentString(token string, url string) (resultado int, deploymentsString string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "deployments"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		deploymentsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, deploymentsString
}

// ListDeploymentProjeto listar Deployment por projetos
func ListDeploymentProjeto(token string, url string, projeto string) (resultado int, deployments model.Deployments) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/deployments"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, deployments
		}
		deployments = model.Deployments{}
		err = json.Unmarshal(corpo, &deployments)
		if err != nil {
			fmt.Println("[ListDeploymentProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, deployments
}

// ListDeploymentProjetoString listar Deployment por projetos
func ListDeploymentProjetoString(token string, url string, projeto string) (resultado int, deploymentsString string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/deployments"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		deploymentsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, deploymentsString
}

// CriarDeployment criar um deployment
func CriarDeployment(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apisAppsv1beta1 + "namespaces/" + projeto + "/deployments"

	resultado, resposta := PostRequest(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
