package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetTemplate recuperar Template
func GetTemplate(token string, url string, projeto string, nome string) (resultado int, template model.Template) {
	resultado = 0
	endpoint := url + apiTemplates + "namespaces/" + projeto + "/templates/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetTemplate] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, template
		}
		template = model.Template{}
		err = json.Unmarshal(corpo, &template)
		if err != nil {
			fmt.Println("[GetTemplate] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, template
}

// GetTemplateString recuperar Template
func GetTemplateString(token string, url string, projeto string, nome string) (resultado int, templateString string) {
	resultado = 0
	endpoint := url + apiTemplates + "namespaces/" + projeto + "/templates/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetTemplateString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		templateString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, templateString
}

// ListTemplate listar todos Templates
func ListTemplate(token string, url string) (resultado int, templates model.Templates) {
	resultado = 0
	endpoint := url + apiTemplates + "templates"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListTemplate] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, templates
		}
		templates = model.Templates{}
		err = json.Unmarshal(corpo, &templates)
		if err != nil {
			fmt.Println("[ListTemplate] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, templates
}

// ListTemplateString listar todos Templates
func ListTemplateString(token string, url string) (resultado int, templatesString string) {
	resultado = 0
	endpoint := url + apiTemplates + "templates"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListTemplateString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		templatesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, templatesString
}

// ListTemplateProjeto listar Templates por projetos
func ListTemplateProjeto(token string, url string, projeto string) (resultado int, templates model.Templates) {
	resultado = 0
	endpoint := url + apiTemplates + "namespaces/" + projeto + "/templates"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListTemplateProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, templates
		}
		templates = model.Templates{}
		err = json.Unmarshal(corpo, &templates)
		if err != nil {
			fmt.Println("[ListTemplateProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, templates
}

// ListTemplateProjetoString listar Templates por projetos
func ListTemplateProjetoString(token string, url string, projeto string) (resultado int, templatesString string) {
	resultado = 0
	endpoint := url + apiTemplates + "namespaces/" + projeto + "/templates"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListTemplateString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		templatesString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, templatesString
}

// CriarTemplate criar um template
func CriarTemplate(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiTemplates + "namespaces/" + projeto + "/templates"

	resultado, resposta := PostRequest(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
