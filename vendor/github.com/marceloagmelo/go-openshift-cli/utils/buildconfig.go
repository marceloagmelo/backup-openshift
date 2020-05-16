package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetBuildConfig recuperar BC
func GetBuildConfig(token string, url string, projeto string, nome string) (resultado int, bc model.BuildConfig) {
	resultado = 0
	endpoint := url + apiBuilds + "namespaces/" + projeto + "/buildconfigs/" + nome

	fmt.Println("[endpoint] : ", endpoint)

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetBuildConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1

		}
		bc = model.BuildConfig{}
		err = json.Unmarshal(corpo, &bc)
		if err != nil {
			fmt.Println("[GetBuildConfig] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, bc
}

// GetBuildConfigString recuperar BC
func GetBuildConfigString(token string, url string, projeto string, nome string) (resultado int, bcString string) {
	resultado = 0
	endpoint := url + apiBuilds + "namespaces/" + projeto + "/buildconfigs/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetBuildConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		bcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, bcString
}

// ListBuildConfig listar todos buildconfig
func ListBuildConfig(token string, url string) (resultado int, bcs model.BuildConfigs) {
	resultado = 0
	endpoint := url + apiBuilds + "buildconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListBuildConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, bcs
		}
		bcs = model.BuildConfigs{}
		err = json.Unmarshal(corpo, &bcs)
		if err != nil {
			fmt.Println("[ListBuildConfig] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, bcs
}

// ListBuildConfigString listar todos buildconfig
func ListBuildConfigString(token string, url string) (resultado int, dcsString string) {
	resultado = 0
	endpoint := url + apiBuilds + "buildconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListBuildConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcsString
}

// ListBuildConfigProjeto listar buildconfig por projetos
func ListBuildConfigProjeto(token string, url string, projeto string) (resultado int, bcs model.BuildConfigs) {
	resultado = 0
	endpoint := url + apiBuilds + "namespaces/" + projeto + "/buildconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListBuildConfigProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, bcs
		}
		bcs = model.BuildConfigs{}
		err = json.Unmarshal(corpo, &bcs)
		if err != nil {
			fmt.Println("[ListBuildConfigProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, bcs
}

// CriarBuildConfig criar um buildconfig
func CriarBuildConfig(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0
	endpoint := url + apiBuilds + "namespaces/" + projeto + "/buildconfigs"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarBuildConfig] Erro ao executar o comando no OS.")
		erro = "[CriarBuildConfig] Erro ao executar o comando no OS."
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
