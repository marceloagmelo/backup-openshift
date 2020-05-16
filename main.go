package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/marceloagmelo/go-backup-openshift/api"
	"github.com/marceloagmelo/go-backup-openshift/logger"
	"github.com/marceloagmelo/go-backup-openshift/model"
	"github.com/marceloagmelo/go-backup-openshift/utils"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

var err error
var gitRepositorio = os.Getenv("GIT_REPOSITORIO")
var gitUsername = os.Getenv("GIT_USERNAME")
var gitPassword = os.Getenv("GIT_PASSWORD")
var limparRecursos = os.Getenv("LIMPAR_RECURSOS")
var execucaoAtivada = os.Getenv("EXECUCAO_ATIVADA")

func init() {
	logger.Info.Println("=== Início ===")

	if execucaoAtivada == "S" {
		variaveis.OpenshiftToken, err = utils.GetToken()
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao recuperar o token", err.Error())
			logger.Erro.Fatalln(mensagem)
			return
		}

		if runtime.GOOS == "windows" {
			limparRecursos = strings.TrimRight(limparRecursos, "\r\n")
		} else {
			limparRecursos = strings.TrimRight(limparRecursos, "\n")
		}
	}
}

func main() {
	if execucaoAtivada == "S" {
		dataFormatada := variaveis.DataHoraAtual.Format(variaveis.DataFormatArquivo)
		variaveis.DirBase = "/tmp/backup-openshift-" + dataFormatada

		// Criar diretórios
		mensagem := fmt.Sprintf("Criando diretorio %s", variaveis.DirBase)
		logger.Info.Println(mensagem)

		os.Mkdir(variaveis.DirBase, 0700)

		// Clonar o respositório de backup
		err := api.GitClone(gitRepositorio, variaveis.DirBase, gitUsername, gitPassword, variaveis.GitlabBranch)
		if err != nil {
			logger.Erro.Fatalln(err)
			return
		}

		// Verficiar se precisa limpar recursos antes
		if len(limparRecursos) > 0 {
			limparRecursos = strings.ToUpper(limparRecursos)
			if limparRecursos == "S" {
				// Copiar a pasta .git para pasta temporária
				utils.CopiarPastaGit(dataFormatada)
			}
		}

		// Recuperar recursos válidos
		recursosValidos, err := model.GetRecursosValidos(variaveis.RecursosFile)
		if err != nil {
			logger.Erro.Fatalln(err)
			return
		}

		mensagem = fmt.Sprintf("Executando o backup dos recursos")
		logger.Info.Println(mensagem)

		err = utils.ExecutarBackup(recursosValidos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o backup dos recursos", err)
			logger.Erro.Println(mensagem)
		}

		dataFormatada = variaveis.DataHoraAtual.Format(variaveis.DataFormatArquivo)

		// Commit e push no git
		err = api.GitCommitPush(variaveis.DirBase, "Backup "+dataFormatada, gitUsername, gitPassword)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o commit no git", err)
			logger.Erro.Fatalln(mensagem)
			return
		}

		// Criar a tag
		mensagem = fmt.Sprintf("Criando a tag %s", dataFormatada)
		logger.Info.Println(mensagem)
		_, _, err = api.GitCriarTag(dataFormatada)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao executar a criação da tag no gitlab", err)
			logger.Erro.Println(mensagem)
		}

		// Remover o diretório base
		mensagem = fmt.Sprintf("Removendo diretorio %s", variaveis.DirBase)
		logger.Info.Println(mensagem)
		os.RemoveAll(variaveis.DirBase)
	}

	logger.Info.Println("=== Fim ===")
}
