package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetImageStream recuperar ImageStream
func GetImageStream(token string, url string, projeto string, nome string) (resultado int, imageStream model.ImageStream) {
	resultado = 0
	endpoint := url + apisImageV1 + "namespaces/" + projeto + "/imagestreams/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetImageStream] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, imageStream
		}
		imageStream = model.ImageStream{}
		err = json.Unmarshal(corpo, &imageStream)
		if err != nil {
			fmt.Println("[GetImageStream] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, imageStream
}

// GetImageStreamString recuperar ImageStream
func GetImageStreamString(token string, url string, projeto string, nome string) (resultado int, imageStreamString string) {
	resultado = 0
	endpoint := url + apisImageV1 + "namespaces/" + projeto + "/imagestreams/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetImageStreamString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		imageStreamString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, imageStreamString
}

// ListImageStream listar todos imagestreams
func ListImageStream(token string, url string) (resultado int, imageStream model.ImageStreams) {
	resultado = 0
	endpoint := url + apisImageV1 + "imagestreams"

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListImageStream] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, imageStream
		}
		imageStream = model.ImageStreams{}
		err = json.Unmarshal(corpo, &imageStream)
		if err != nil {
			fmt.Println("[ListImageStream] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, imageStream
}

// ListImageStreamString listar todos imagestreams
func ListImageStreamString(token string, url string) (resultado int, imageStreamString string) {
	resultado = 0
	endpoint := url + apisImageV1 + "imagestreams"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListImageStreamString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		imageStreamString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, imageStreamString
}

// ListImageStreamProjeto listar imagestreams por projetos
func ListImageStreamProjeto(token string, url string, projeto string) (resultado int, imageStream model.ImageStreams) {
	resultado = 0
	endpoint := url + apisImageV1 + "namespaces/" + projeto + "/imagestreams"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListImageStreamProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, imageStream
		}
		imageStream = model.ImageStreams{}
		err = json.Unmarshal(corpo, &imageStream)
		if err != nil {
			fmt.Println("[ListImageStreamProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, imageStream
}

// ListImageStreamProjetoString listar imagestreams por projetos
func ListImageStreamProjetoString(token string, url string, projeto string) (resultado int, imageStreamString string) {
	resultado = 0
	endpoint := url + apisImageV1 + "namespaces/" + projeto + "/imagestreams"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListImageStreamString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		imageStreamString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, imageStreamString
}

// CriarImageStream criar um image stream
func CriarImageStream(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apisImageV1 + "namespaces/" + projeto + "/imagestreams"

	resultado, resposta := PostRequest(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	return resultado, erro
}
