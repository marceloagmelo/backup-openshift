package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetResourceQuota recuperar ResourceQuota
func GetResourceQuota(token string, url string, projeto string, nome string) (resultado int, resourceQuota model.ResourceQuota) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/resourcequotas/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getResourceQuota] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, resourceQuota
		}
		resourceQuota = model.ResourceQuota{}
		err = json.Unmarshal(corpo, &resourceQuota)
		if err != nil {
			fmt.Println("[getResourceQuota] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, resourceQuota
}

// GetResourceQuotaString recuperar ResourceQuota
func GetResourceQuotaString(token string, url string, projeto string, nome string) (resultado int, resourceQuotaString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/resourcequotas/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[getResourceQuotaString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		resourceQuotaString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, resourceQuotaString
}

// ListResourceQuota listar todos resourcequotas
func ListResourceQuota(token string, url string) (resultado int, resourceQuotas model.ResourceQuotas) {
	resultado = 0
	endpoint := url + apiV1 + "resourcequotas"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListResourceQuota] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, resourceQuotas
		}
		resourceQuotas = model.ResourceQuotas{}
		err = json.Unmarshal(corpo, &resourceQuotas)
		if err != nil {
			fmt.Println("[ListResourceQuota] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, resourceQuotas
}

// ListresourceQuotaString listar todos resourcequotas
func ListresourceQuotaString(token string, url string) (resultado int, resourceQuotaString string) {
	resultado = 0
	endpoint := url + apiV1 + "resourcequotas"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListresourceQuotaString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		resourceQuotaString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, resourceQuotaString
}

// ListResourceQuotaProjeto listar resourcequotas por projetos
func ListResourceQuotaProjeto(token string, url string, projeto string) (resultado int, resourceQuota model.ResourceQuota) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/resourcequotas"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListResourceQuotaProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, resourceQuota
		}
		resourceQuota = model.ResourceQuota{}
		err = json.Unmarshal(corpo, &resourceQuota)
		if err != nil {
			fmt.Println("[ListResourceQuotaProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, resourceQuota
}

// ListResourceQuotaProjetoString listar resourcequotas por projetos
func ListResourceQuotaProjetoString(token string, url string, projeto string) (resultado int, resourceQuotaString string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/resourcequotas"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListresourceQuotaString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		resourceQuotaString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, resourceQuotaString
}

// CriarResourceQuota criar um resourcequotas
func CriarResourceQuota(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiV1 + "namespaces/" + projeto + "/resourcequotas"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarResourceQuota] Erro ao executar o comando no OS.")
		erro = "[CriarResourceQuota] Erro ao executar o comando no OS."
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
