package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/marceloagmelo/go-backup-openshift/logger"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

//SalvarArquivoJSON salvar o arquivo
func SalvarArquivoJSON(arquivo string, texto string) error {
	arquivoJSON, err := os.Create(arquivo)
	if err != nil {
		mensagem := fmt.Sprintf("Houve um erro ao criar o arquivo %s. Erro: %s", arquivo, err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	defer arquivoJSON.Close()
	escritorArquivo := bufio.NewWriter(arquivoJSON)
	escritorArquivo.WriteString(texto)
	escritorArquivo.Flush()
	return nil
}

//CopyFile
func CopyFile(src, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		mensagem := fmt.Sprintf("Houve um erro ao copiar o arquivo origem %s para destino %s. Erro: %s", src, dst, err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	if !sfi.Mode().IsRegular() {
		mensagem := fmt.Sprintf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			mensagem := fmt.Sprintf("Houve um erro o caminho destino %s não encontrado. Erro: %s", dst, err)
			logger.Erro.Println(mensagem)
			err := errors.New(mensagem)
			return err
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			mensagem := fmt.Sprintf("CopyFile: non-regular destination file %s (%q)", sfi.Name(), sfi.Mode().String())
			logger.Erro.Println(mensagem)
			err := errors.New(mensagem)
			return err
		}
		if os.SameFile(sfi, dfi) {
			return nil
		}
	}
	if err = os.Link(src, dst); err == nil {
		return nil
	}
	err = copyFileContents(src, dst)
	if err != nil {
		return err
	}
	return nil
}

func copyFileContents(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		mensagem := fmt.Sprintf("Houve um erro ao abrir o arquivo origem %s. Erro: %s", src, err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		mensagem := fmt.Sprintf("Houve um erro ao abrir o arquivo destino %s. Erro: %s", dst, err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
			mensagem := fmt.Sprintf("Erro: %s", err)
			logger.Erro.Println(mensagem)
			err = errors.New(mensagem)
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		mensagem := fmt.Sprintf("Houve um erro ao copiar arquivo. Erro: %s", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	err = out.Sync()
	if err != nil {
		mensagem := fmt.Sprintf("Houve um erro ao sincronizar arquivo. Erro: %s", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	return nil
}

//CopiarPastaGit copiar a pasta do git
func CopiarPastaGit(dataFormatada string) error {
	// Copiando pasta .git para pasta temporária
	dirTemp := "/tmp/backup-openshift-temp-" + dataFormatada

	mensagem := fmt.Sprintf("Copiando a pasta .git do diretorio %s para o diretório %s", variaveis.DirBase, dirTemp)
	logger.Info.Println(mensagem)

	os.Mkdir(dirTemp, 0700)

	cpCmd := exec.Command("cp", "-rf", variaveis.DirBase+"/.git/", dirTemp)
	err := cpCmd.Run()
	if err != nil {
		mensagem := fmt.Sprintf("Falha na copia da pasta .git %q", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}
	mensagem = fmt.Sprintf("Pasta copiada com sucesso.")
	logger.Info.Println(mensagem)

	// Remover o diretório base
	mensagem = fmt.Sprintf("Removendo diretorio %s", variaveis.DirBase)
	logger.Info.Println(mensagem)
	os.RemoveAll(variaveis.DirBase)

	// Criar diretórios
	mensagem = fmt.Sprintf("recriando diretorio %s", variaveis.DirBase)
	logger.Info.Println(mensagem)
	os.Mkdir(variaveis.DirBase, 0700)

	// Copiando pasta .git para pasta temporária
	mensagem = fmt.Sprintf("Copiando a pasta .git do diretorio %s para o diretório %s", dirTemp, variaveis.DirBase)
	logger.Info.Println(mensagem)
	cpCmd = exec.Command("cp", "-rf", dirTemp+"/.git/", variaveis.DirBase)
	err = cpCmd.Run()
	if err != nil {
		mensagem := fmt.Sprintf("Falha na copia da pasta .git %q", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return err
	}

	mensagem = fmt.Sprintf("Pasta copiada com sucesso.")
	logger.Info.Println(mensagem)

	// Remover o diretório temp
	mensagem = fmt.Sprintf("removendo diretorio %s", dirTemp)
	logger.Info.Println(mensagem)
	os.RemoveAll(dirTemp)

	return nil
}

//GetToken recuperar Token do usuário.
func GetToken() (string, error) {
	resposta := ""
	endpoint := variaveis.OpenshiftApiURL + "/oauth/authorize?client_id=openshift-challenging-client&response_type=token"

	cmdCurl := "curl -s -u " + variaveis.OpenshiftUsername + ":" + variaveis.OpenshiftPassword + " -kI '" + endpoint + "' | grep -oP 'access_token=\\K[^&]*'"

	logger.Info.Println(cmdCurl)
	resposta, err := ExecCmd(cmdCurl)

	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o CURL", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return resposta, err
	}
	return resposta, nil
}

//ExecCmd execuctar comando no OS.
func ExecCmd(strCurl string) (string, error) {
	resposta := ""
	cmd := exec.Command("/bin/bash", "-c", strCurl)

	out, err := cmd.CombinedOutput()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o comando no OS", err)
		logger.Erro.Println(mensagem)
		err := errors.New(mensagem)
		return resposta, err
	}
	resposta = string(out)

	return resposta, nil
}
