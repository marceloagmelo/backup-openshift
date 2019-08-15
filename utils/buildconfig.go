package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupBcs lista dos bcs do openshift
func BackupBcs(token string, url string) {
	// Listar os dcs
	fmt.Printf("Backup dos recursos buildconfig do ambiente %s\n\r", url)
	resultado, bcs := utils.ListBuildConfig(token, url)
	if resultado == 0 {
		// Ler os dados dos bcs
		lerDadosBcs(token, bcs)
	} else {
		fmt.Println("[BackupBcs] Bcs não encontrados")
	}
}

func lerDadosBcs(token string, bcs model.BuildConfigs) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(bcs.Items); i++ {
		nomeProjeto := bcs.Items[i].Metadata.Namespace
		nomeDc := bcs.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarBc(token, nomeProjeto, nomeDc)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarBc(token string, nomeProjeto string, nomeBc string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirBc := dirProjeto + "/buildconfig"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirBc, 0700)

	resultado, bc := utils.GetBuildConfigString(token, url, nomeProjeto, nomeBc)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirBc + "/" + nomeBc + ".json"
		resultado = SalvarArquivoJSON(arquivo, bc)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarBc] Bc %s não encontrado no projeto %s ambiente %s", nomeBc, nomeProjeto, url)
	}
	return recursoSalvo
}
