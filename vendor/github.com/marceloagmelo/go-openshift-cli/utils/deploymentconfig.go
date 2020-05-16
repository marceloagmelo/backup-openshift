package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

// GetDeploymentConfig recuperar DC
func GetDeploymentConfig(token string, url string, projeto string, nome string) (resultado int, dc model.DeploymentConfig) {
	resultado = 0
	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, dc
		}
		dc = model.DeploymentConfig{}
		err = json.Unmarshal(corpo, &dc)
		if err != nil {
			fmt.Println("[GetDeploymentConfig] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dc
}

// GetDeploymentConfigString recuperar DC
func GetDeploymentConfigString(token string, url string, projeto string, nome string) (resultado int, dcString string) {
	resultado = 0
	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs/" + nome

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcString
}

// ListDeploymentConfig listar todos deploymentconfig
func ListDeploymentConfig(token string, url string) (resultado int, dcs model.DeploymentConfigs) {
	resultado = 0
	endpoint := url + apiApps + "deploymentconfigs"

	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, dcs
		}
		dcs = model.DeploymentConfigs{}
		err = json.Unmarshal(corpo, &dcs)
		if err != nil {
			fmt.Println("[ListDeploymentConfig] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dcs
}

// ListDeploymentConfigString listar todos deploymentconfig
func ListDeploymentConfigString(token string, url string) (resultado int, dcsString string) {
	resultado = 0
	endpoint := url + apiApps + "deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentConfig] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcsString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcsString
}

// ListDeploymentConfigProjeto listar deploymentconfig por projetos
func ListDeploymentConfigProjeto(token string, url string, projeto string) (resultado int, dcs model.DeploymentConfigs) {
	resultado = 0
	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs"
	resultado, resposta := GetRequest(token, endpoint)
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[ListDeploymentConfigProjeto] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, dcs
		}
		dcs = model.DeploymentConfigs{}
		err = json.Unmarshal(corpo, &dcs)
		if err != nil {
			fmt.Println("[ListDeploymentConfigProjeto] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}
	return resultado, dcs
}

// GetDeploymentConfigImage recuperar DC
func GetDeploymentConfigImage(dc model.DeploymentConfig) string {
	return dc.Spec.Template.Spec.Containers[0].Image
}

// CriarDeploymentConfig criar um deployment config
func CriarDeploymentConfig(token string, url string, projeto string, conteudoJSON string) (resultado int, erro string) {
	resultado = 0

	endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs"

	cmd := "sed -i s/\\\"resourceVersion[^,]*,//g " + conteudoJSON
	resultado, _ = ExecCmd(cmd)

	if resultado > 0 {
		fmt.Println("[CriarDeploymentConfig] Erro ao executar o comando no OS.")
		erro = "[CriarDeploymentConfig] Erro ao executar o comando no OS."
		return resultado, erro
	}

	resultado, resposta := PostRequestFile(token, endpoint, conteudoJSON)
	//resultado, resposta := PostRequest(token, endpoint, conteudoJSON)
	defer resposta.Body.Close()
	if resposta.StatusCode != 201 {
		erro = resposta.Status
		resultado = 1
	}
	/*
		token = strings.TrimRight(token, "\r\n")
		bearer := `"Authorization: Bearer ` + token + `"`
		header := "-k -X POST -d @- -H " + bearer + " -H 'Accept: application/json' -H 'Content-Type: application/json'"
		endpoint := url + apiApps + "namespaces/" + projeto + "/deploymentconfigs"
		cmdCurl := "curl " + header + " " + endpoint + " <<'EOF'\n\r" + conteudoJSON + "EOF "

		resultado, erro = utils.ExecCurl(cmdCurl)

		if resultado > 0 {
			fmt.Println("[CriarDeploymentConfig] Erro ao executar o CURL.")
			erro = "[CriarDeploymentConfig] Erro ao executar o CURL."
			resultado = 1
		}
	*/
	return resultado, erro
}
