package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

//SalvarArquivoJSON salvar o arquivo
func SalvarArquivoJSON(arquivo string, texto string) (resultado int) {
	resultado = 0
	arquivoJSON, err := os.Create(arquivo)
	if err != nil {
		fmt.Printf("[SalvarArquivoJSON] Houve um erro ao criar o arquivo %s. Erro: %s\n\r", arquivo, err.Error())
		resultado = 1
	}
	defer arquivoJSON.Close()
	escritorArquivo := bufio.NewWriter(arquivoJSON)
	escritorArquivo.WriteString(texto)
	escritorArquivo.Flush()
	return resultado
}

//CopyFile
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

//CopiarPastaGit copiar a pasta do git
func CopiarPastaGit(dataFormatada string) {
	// Copiando pasta .git para pasta temporária
	dirTemp := "/tmp/backup-openshift-temp-" + dataFormatada
	fmt.Printf("Copiando a pasta .git do diretorio %s para o diretório %s\n\r", variaveis.DirBase, dirTemp)
	os.Mkdir(dirTemp, 0700)
	cpCmd := exec.Command("cp", "-rf", variaveis.DirBase+"/.git/", dirTemp)
	err := cpCmd.Run()
	if err != nil {
		fmt.Printf("Falha na copia da pasta .git %q\n", err)
	} else {
		fmt.Printf("Pasta copiada com sucesso.\n")
	}

	// Remover o diretório base
	fmt.Printf("removendo diretorio %s\n\r", variaveis.DirBase)
	os.RemoveAll(variaveis.DirBase)

	// Criar diretórios
	fmt.Printf("recriando diretorio %s\n\r", variaveis.DirBase)
	os.Mkdir(variaveis.DirBase, 0700)

	// Copiando pasta .git para pasta temporária
	fmt.Printf("Copiando a pasta .git do diretorio %s para o diretório %s\n\r", dirTemp, variaveis.DirBase)
	cpCmd = exec.Command("cp", "-rf", dirTemp+"/.git/", variaveis.DirBase)
	err = cpCmd.Run()
	if err != nil {
		fmt.Printf("Falha na copia da pasta .git %q\n", err)
	} else {
		fmt.Printf("Pasta copiada com sucesso.\n")
	}

	// Remover o diretório temp
	fmt.Printf("removendo diretorio %s\n\r", dirTemp)
	os.RemoveAll(dirTemp)
}
