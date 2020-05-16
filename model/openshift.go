package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/logger"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

//OpenshiftErro retorno de erro
type OpenshiftErro struct {
	APIVersion string `json:"apiVersion"`
	Code       int    `json:"code"`
	Details    struct {
	} `json:"details"`
	Kind     string `json:"kind"`
	Message  string `json:"message"`
	Metadata struct {
	} `json:"metadata"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}

//OpenshiftSuccess retorno de sucesso
type OpenshiftSuccess struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
	} `json:"metadata"`
	Status  string `json:"status"`
	Details struct {
		Name string `json:"name"`
		Kind string `json:"kind"`
		UID  string `json:"uid"`
	} `json:"details"`
}

//RecursosValidos
type RecursosValidos struct {
	Recursos []struct {
		Nome string `json:"nome"`
	} `json:"recursos"`
}

//GetRecursosValidos
func GetRecursosValidos(arquivo string) (RecursosValidos, error) {
	recursosValidos := RecursosValidos{}

	jsonFile, err := os.Open(arquivo)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao ler o arquivo de recursos válidos [%s]: %s", arquivo, err.Error())
		logger.Erro.Println(mensagem)
		return recursosValidos, err
	}
	defer jsonFile.Close()

	// Ler o json como um array de bytes
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao converter o arquivo [%s] para bytes: %s", variaveis.RecursosFile, err.Error())
		logger.Erro.Println(mensagem)
		return recursosValidos, err
	}

	err = json.Unmarshal(byteValue, &recursosValidos)
	if err != nil {
		mensagem := fmt.Sprintf("Erro ao converter o arquivo [%s] para o struct de recursos válidos: %s", variaveis.RecursosFile, err.Error())
		logger.Erro.Println(mensagem)
		return recursosValidos, err
	}

	return recursosValidos, nil
}
