package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarDcs lista dos dcs do openshift
func ListarDcs(token string, url string) {
	// Listar os dcs
	fmt.Printf("Listando todos os dcs do ambiente %s\n\r", url)
	resultado, dcs := utils.ListDc(token, url)
	if resultado == 0 {

		// Ler os dados dos dcs
		lerDadosDadosDcs(token, dcs)
	} else {
		fmt.Println("[main] Dcs não encontrados")
	}
}

func lerDadosDadosDcs(token string, dcs model.Dcs) {
	for i := 0; i < len(dcs.Items); i++ {
		nomeProjeto := dcs.Items[i].Metadata.Namespace
		nomeDc := dcs.Items[i].Metadata.Name

		lerDc(token, nomeProjeto, nomeDc)
	}
}

func lerDc(token string, nomeProjeto string, nomeDc string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirDc := dirProjeto + "/dc"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirDc, 0700)

	resultado, dc := utils.GetDcString(token, url, nomeProjeto, nomeDc)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirDc + "/" + nomeDc + ".json"
		SalvarArquivoJSON(arquivo, dc)
	} else {
		fmt.Printf("[lerDc] Dc %s não encontrado no projeto %s ambiente %s", nomeDc, nomeProjeto, url)
	}
}
