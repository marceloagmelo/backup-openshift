package utils

import (
	"bufio"
	"fmt"
	"os"
)

//SalvarArquivoJSON salvar o arquivo
func SalvarArquivoJSON(arquivo string, texto string) (resultado int) {
	resultado = 0
	arquivoJSON, err := os.Create(arquivo)
	if err != nil {
		fmt.Printf("[SalvarArquivo] Houve um erro ao criar o arquivo %s. Erro: %s\n\r", arquivo, err.Error())
		resultado = 1
	}
	defer arquivoJSON.Close()
	escritorArquivo := bufio.NewWriter(arquivoJSON)
	escritorArquivo.WriteString(texto)
	escritorArquivo.Flush()
	return resultado
}
