package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarPvcs lista dos pvcs do openshift
func ListarPvcs(token string, url string) {
	// Listar os pvcs
	fmt.Printf("Listando todos os pvcs do ambiente %s\n\r", url)
	resultado, pvcs := utils.ListPvc(token, url)
	if resultado == 0 {

		// Ler os dados dos pvcs
		lerDadosDadosPvcs(token, pvcs)
	} else {
		fmt.Println("[ListarPvcs] Pvcs não encontrados")
	}
}

func lerDadosDadosPvcs(token string, pvcs model.Pvcs) {
	for i := 0; i < len(pvcs.Items); i++ {
		nomeProjeto := pvcs.Items[i].Metadata.Namespace
		nomePvc := pvcs.Items[i].Metadata.Name

		lerPvc(token, nomeProjeto, nomePvc)
	}
}

func lerPvc(token string, nomeProjeto string, nomePvc string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirPvc := dirProjeto + "/pvc"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirPvc, 0700)

	resultado, pvc := utils.GetPvcString(token, url, nomeProjeto, nomePvc)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirPvc + "/" + nomePvc + ".json"
		SalvarArquivoJSON(arquivo, pvc)
	} else {
		fmt.Printf("[lerPvc] Pvc %s não encontrado no projeto %s ambiente %s\n\r", nomePvc, nomeProjeto, url)
	}
}
