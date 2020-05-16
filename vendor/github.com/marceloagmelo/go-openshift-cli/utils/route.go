package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetRoute recuperar Route
func GetRoute(token string, url string, projeto string, nome string) (resultado int, route model.Route) {
	resultado = 0
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRoute] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, route
		}
		route = model.Route{}
		err = json.Unmarshal(corpo, &route)
		if err != nil {
			fmt.Println("[GetRoute] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, route
}

// GetRouteString recuperar Route
func GetRouteString(token string, url string, projeto string, nome string) (resultado int, routeString string) {
	resultado = 0
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRouteString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routeString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, routeString
}

// ListRoute listar todos Routes
func ListRoute(token string, url string) (resultado int, routes model.Routes) {
	resultado = 0
	endpoint := url + apiRoutes + "routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRoute] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, routes
		}
		routes = model.Routes{}
		err = json.Unmarshal(corpo, &routes)
		if err != nil {
			fmt.Println("[ListRoute] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, routes
}

// ListRouteString listar todos Routes
func ListRouteString(token string, url string) (resultado int, routesString string) {
	resultado = 0
	endpoint := url + apiRoutes + "routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRouteString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, routesString
}

// ListRouteProjeto listar Routes por projetos
func ListRouteProjeto(token string, url string, projeto string) (resultado int, routes model.Routes) {
	resultado = 0
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRouteProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, routes
		}
		routes = model.Routes{}
		err = json.Unmarshal(corpo, &routes)
		if err != nil {
			fmt.Println("[ListRouteProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, routes
}

// ListRouteProjetoString listar Routes por projetos
func ListRouteProjetoString(token string, url string, projeto string) (resultado int, routesString string) {
	resultado = 0
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListRouteString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		routesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, routesString
}

// CriarRoute criar uma rota
func CriarRoute(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiRoutes + "namespaces/" + projeto + "/routes"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarRoute] Erro ao executar o comando no OS.")
		erro = "[CriarRoute] Erro ao executar o comando no OS."
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
