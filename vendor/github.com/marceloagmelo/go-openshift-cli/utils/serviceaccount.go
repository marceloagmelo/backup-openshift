package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetServiceAccount recuperar ServiceAccount
func GetServiceAccount(token string, url string, projeto string, nome string) (resultado int, serviceAccount model.ServiceAccount) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getServiceAccount] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, serviceAccount
		}
		serviceAccount = model.ServiceAccount{}
		err = json.Unmarshal(corpo, &serviceAccount)
		if err != nil {
			fmt.Println("[getServiceAccount] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, serviceAccount
}

// GetServiceAccountString recuperar ServiceAccount
func GetServiceAccountString(token string, url string, projeto string, nome string) (resultado int, serviceAccountString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getServiceAccountString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccountString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceAccountString
}

// ListServiceAccount listar todos serviceaccounts
func ListServiceAccount(token string, url string) (resultado int, serviceAccounts model.ServiceAccounts) {
	resultado = 0
	endpoint := url + apiV1 + "serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceAccount] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, serviceAccounts
		}
		serviceAccounts = model.ServiceAccounts{}
		err = json.Unmarshal(corpo, &serviceAccounts)
		if err != nil {
			fmt.Println("[ListServiceAccount] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, serviceAccounts
}

// ListserviceAccountString listar todos serviceaccounts
func ListserviceAccountString(token string, url string) (resultado int, serviceAccountString string) {
	resultado = 0
	endpoint := url + apiV1 + "serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListserviceAccountString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccountString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceAccountString
}

// ListServiceAccountProjeto listar serviceaccounts por projetos
func ListServiceAccountProjeto(token string, url string, projeto string) (resultado int, serviceAccount model.ServiceAccount) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListServiceAccountProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, serviceAccount
		}
		serviceAccount = model.ServiceAccount{}
		err = json.Unmarshal(corpo, &serviceAccount)
		if err != nil {
			fmt.Println("[ListServiceAccountProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, serviceAccount
}

// ListServiceAccountProjetoString listar serviceaccounts por projetos
func ListServiceAccountProjetoString(token string, url string, projeto string) (resultado int, serviceAccountString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListserviceAccountString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		serviceAccountString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, serviceAccountString
}

// CriarServiceAccount criar um serviceaccounts
func CriarServiceAccount(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/serviceaccounts"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarServiceAccount] Erro ao executar o comando no OS.")
		erro = "[CriarServiceAccount] Erro ao executar o comando no OS."
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
