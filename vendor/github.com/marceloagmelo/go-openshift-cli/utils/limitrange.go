package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetLimitRange recuperar LimitRange
func GetLimitRange(token string, url string, projeto string, nome string) (resultado int, LimitRange model.LimitRange) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/limitranges/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getLimitRange] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, LimitRange
		}
		LimitRange = model.LimitRange{}
		err = json.Unmarshal(corpo, &LimitRange)
		if err != nil {
			fmt.Println("[getLimitRange] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, LimitRange
}

// GetLimitRangeString recuperar LimitRange
func GetLimitRangeString(token string, url string, projeto string, nome string) (resultado int, LimitRangeString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/limitranges/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getLimitRangeString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		LimitRangeString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, LimitRangeString
}

// ListLimitRange listar todos LimitRanges
func ListLimitRange(token string, url string) (resultado int, LimitRanges model.LimitRanges) {
	resultado = 0
	endpoint := url + apiV1 + "limitranges"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListLimitRange] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, LimitRanges
		}
		LimitRanges = model.LimitRanges{}
		err = json.Unmarshal(corpo, &LimitRanges)
		if err != nil {
			fmt.Println("[ListLimitRange] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, LimitRanges
}

// ListLimitRangeString listar todos LimitRanges
func ListLimitRangeString(token string, url string) (resultado int, LimitRangeString string) {
	resultado = 0
	endpoint := url + apiV1 + "limitranges"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListLimitRangeString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		LimitRangeString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, LimitRangeString
}

// ListLimitRangeProjeto listar LimitRanges por projetos
func ListLimitRangeProjeto(token string, url string, projeto string) (resultado int, LimitRange model.LimitRange) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/limitranges"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListLimitRangeProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, LimitRange
		}
		LimitRange = model.LimitRange{}
		err = json.Unmarshal(corpo, &LimitRange)
		if err != nil {
			fmt.Println("[ListLimitRangeProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, LimitRange
}

// ListLimitRangeProjetoString listar LimitRanges por projetos
func ListLimitRangeProjetoString(token string, url string, projeto string) (resultado int, LimitRangeString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/limitranges"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListLimitRangeString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		LimitRangeString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, LimitRangeString
}

// CriarLimitRange criar um LimitRanges
func CriarLimitRange(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/limitranges"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarLimitRange] Erro ao executar o comando no OS.")
		erro = "[CriarLimitRange] Erro ao executar o comando no OS."
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
