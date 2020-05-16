package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/marceloagmelo/go-openshift-cli/model"
)

//var configfile = "/home/marcelo/go/src/github.com/marceloagmelo/go-openshift-cli/conf/config.yaml"
var apiProjeto = "/apis/project.openshift.io/v1/projects/"
var apiApps = "/apis/apps.openshift.io/v1/"
var apiPods = "/api/v1/pods/"
var apiV1 = "/api/v1/"
var apiRoutes = "/apis/route.openshift.io/v1/"
var apisAppsv1beta1 = "/apis/apps/v1beta1/"
var apisImageV1 = "/apis/image.openshift.io/v1/"
var apisAuthorizationOpenshiftV1 = "/apis/authorization.openshift.io/v1/"
var apisExtensionsV1beta1 = "/apis/extensions/v1beta1/"
var apiBuilds = "/apis/build.openshift.io/v1/"
var apiTemplates = "/apis/template.openshift.io/v1/"

// GetToken recuperar Token do usuário.
func GetToken(url string, username string, password string) (resultado int, resposta string) {
	endpoint := url + "/oauth/authorize?client_id=openshift-challenging-client&response_type=token"

	cmdCurl := "curl -s -u " + username + ":" + password + " -kI '" + endpoint + "' | grep -oP 'access_token=\\K[^&]*'"

	resultado, resposta = ExecCmd(cmdCurl)

	if resultado > 0 {
		fmt.Println("[GetToken] Erro ao executar o CURL.")
	}
	return resultado, resposta
}

// ExecCmd execuctar comando no OS.
func ExecCmd(strCurl string) (resultado int, resposta string) {
	resultado = 0
	cmd := exec.Command("/bin/bash", "-c", strCurl)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ExecCmd:", err)
		resultado = 1
	} else {
		resposta = string(out)
	}

	return resultado, resposta
}

// GetRequest recuperar a requisição
func GetRequest(token string, endpoint string) (resultado int, resposta *http.Response) {
	resultado = 0
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("[GetRequest] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
		return resultado, resposta
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[GetRequest] Erro ao abrir a pagina. Erro: ", err.Error())
		resultado = 1
	}
	return resultado, resposta
}

// GetRequestJSON recuperar a requisição
func GetRequestJSON(token string, endpoint string) (resultado int, dc model.Dc) {
	resultado = 0

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("[GetRequestJSON] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
		return resultado, dc
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[GetRequestJSON] Erro ao abrir a pagina. Erro: ", err.Error())
		resultado = 1
		return resultado, dc
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRequestJSON] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
			return resultado, dc
		}
		//dc := model.Dc{}
		err = json.Unmarshal(corpo, &dc)
		if err != nil {
			fmt.Println("[GetDc] Erro ao converter o retorno JSON do Servidor. Erro: ", err.Error())
			resultado = 1
		}
	} else {
		resultado = 1
	}

	return resultado, dc
}

// GetRequestString recuperar a requisição
func GetRequestString(token string, endpoint string) (resultado int, dcString string) {
	resultado = 0

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("[GetRequestString] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
		return resultado, dcString
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")

	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[GetRequestString] Erro ao abrir a pagina. Erro: ", err.Error())
		resultado = 1
		return resultado, dcString
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[GetRequestString] Erro ao ler o conteudo da pagina. Erro: ", err.Error())
			resultado = 1
		}
		dcString = string(corpo)
	} else {
		resultado = 1
	}
	return resultado, dcString
}

// PostRequest recuperar a requisição
func PostRequest(token string, endpoint string, conteudoJSON string) (resultado int, resposta *http.Response) {
	resultado = 0
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}
	var service = model.Service{}
	err := json.Unmarshal([]byte(conteudoJSON), &service)
	if err != nil {
		fmt.Println("[PostRequest] Erro ao gerar o objeto serviço com o JSON lido. Erro. Erro: ", err.Error())
		resultado = 1
		return resultado, resposta
	}

	conteudoEnviar, err := json.Marshal(&service)
	if err != nil {
		fmt.Println("[PostRequest] Erro ao gerar o JSON do objeto serviço. Erro. Erro: ", err.Error())
		resultado = 1
		return resultado, resposta
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		fmt.Println("[PostRequest] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
		return resultado, resposta
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")
	request.Header.Set("Content-Type", "application/json")

	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[PostRequest] Erro ao executar o post. Erro: ", err.Error())
		resultado = 1
	}
	return resultado, resposta
}

// PostRequestFile recuperar a requisição
func PostRequestFile(token string, endpoint string, arquivo string) (resultado int, resposta *http.Response) {
	resultado = 0
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}

	f, err := os.Open(arquivo)
	if err != nil {
		fmt.Println("[PostRequest] Erro ao ler o arquivo temporaário. Erro: ", err.Error())
		resultado = 1
		return resultado, resposta
	}

	request, err := http.NewRequest("POST", endpoint, f)
	if err != nil {
		fmt.Println("[PostRequest] Erro ao criar um request. Erro: ", err.Error())
		resultado = 1
		return resultado, resposta
	}

	var bearer = "Bearer " + strings.TrimSuffix(token, "\n")
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Accept", "application/json; charset=utf-8")
	request.Header.Set("Content-Type", "application/json")

	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[PostRequest] Erro ao executar o post. Erro: ", err.Error())
		resultado = 1
	}
	return resultado, resposta
}
