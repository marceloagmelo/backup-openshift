package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

var data = `
username: marceloagmelo@gmail.com
password: magm0101
`

type generalConfig struct {
	Usuario string `yaml:"username"`
	Senha   string `yaml:"password"`
}

// generalConfigLoad load general configuration from conf/config.yaml
func generalConfigLoad() (*generalConfig, error) {
	var config generalConfig

	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

//GetUsuarioSenha
func GetUsuarioSenha() (usuario string, senha string) {
	config, err := generalConfigLoad()
	if err != nil {
		fmt.Println("GetUsuarioSenha:", err)
		os.Exit(1)
	}

	usuario = config.Usuario
	senha = config.Senha

	return usuario, senha
}

// ExecCurl execuctar CURL.
func ExecCurl(strCurl string) (resultado int, resposta string) {
	resultado = 0
	cmd := exec.Command("/bin/bash", "-c", strCurl)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ExecCurl:", err)
		resultado = 1
	} else {
		resposta = string(out)
	}

	return resultado, resposta
}

//SalvarArquivoJSON salvar o arquivo JSON
func SalvarArquivoJSON(arquivo string, texto string) {
	arquivoJSON, err := os.Create(arquivo)
	if err != nil {
		fmt.Println("[salvarDc] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
	}
	defer arquivoJSON.Close()
	escritorDc := bufio.NewWriter(arquivoJSON)
	escritorDc.WriteString(texto)
	escritorDc.Flush()
}
