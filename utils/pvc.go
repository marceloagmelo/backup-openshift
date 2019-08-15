package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupPvcs lista dos pvcs do openshift
func BackupPvcs(token string, url string) {
	// Listar os pvcs
	fmt.Printf("Backup dos recursos pvcs do ambiente %s\n\r", url)
	resultado, pvcs := utils.ListPvc(token, url)
	if resultado == 0 {

		// Ler os dados dos pvcs
		lerDadosPvcs(token, pvcs)
	} else {
		fmt.Println("[BackupPvcs] Pvcs não encontrados")
	}
}

func lerDadosPvcs(token string, pvcs model.Pvcs) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(pvcs.Items); i++ {
		nomeProjeto := pvcs.Items[i].Metadata.Namespace
		nomePvc := pvcs.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarPvc(token, nomeProjeto, nomePvc)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarPvc(token string, nomeProjeto string, nomePvc string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirPvc := dirProjeto + "/pvc"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirPvc, 0700)

	resultado, pvc := utils.GetPvcString(token, url, nomeProjeto, nomePvc)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirPvc + "/" + nomePvc + ".json"
		resultado = SalvarArquivoJSON(arquivo, pvc)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarPvc] Pvc %s não encontrado no projeto %s ambiente %s\n\r", nomePvc, nomeProjeto, url)
	}
	return recursoSalvo
}
