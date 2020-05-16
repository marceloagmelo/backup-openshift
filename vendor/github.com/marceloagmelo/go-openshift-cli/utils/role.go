package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetRole recuperar Role
func GetRole(token string, url string, projeto string, nome string) (resultado int, role model.Role) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/roles/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRole] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, role
		}
		role = model.Role{}
		err = json.Unmarshal(corpo, &role)
		if err != nil {
			fmt.Println("[GetRole] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, role
}

// GetRoleString recuperar Role
func GetRoleString(token string, url string, projeto string, nome string) (resultado int, roleString string) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/roles/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRoleString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		roleString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, roleString
}

// ListRole listar todos Roles
func ListRole(token string, url string) (resultado int, role model.Roles) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "roles"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRole] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, role
		}
		role = model.Roles{}
		err = json.Unmarshal(corpo, &role)
		if err != nil {
			fmt.Println("[ListRole] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, role
}

// ListRoleString listar todos Roles
func ListRoleString(token string, url string) (resultado int, roleString string) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "roles"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		roleString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, roleString
}

// ListRoleProjeto listar Roles por projetos
func ListRoleProjeto(token string, url string, projeto string) (resultado int, role model.Roles) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/roles"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, role
		}
		role = model.Roles{}
		err = json.Unmarshal(corpo, &role)
		if err != nil {
			fmt.Println("[ListRoleProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, role
}

// ListRoleProjetoString listar Roles por projetos
func ListRoleProjetoString(token string, url string, projeto string) (resultado int, roleString string) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/roles"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoleString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		roleString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, roleString
}

// CriarRole criar um Roles
func CriarRole(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apisAuthorizationOpenshiftV1 + "namespaces/" + projeto + "/roles"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarRole] Erro ao executar o comando no OS.")
		erro = "[CriarRole] Erro ao executar o comando no OS."
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
