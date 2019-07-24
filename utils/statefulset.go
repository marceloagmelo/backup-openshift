package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupStateFulSets lista dos statefulsets do openshift
func BackupStateFulSets(token string, url string) {
	// Listar os statefulsets
	fmt.Printf("Backup dos recursos statefulsets do ambiente %s\n\r", url)
	resultado, statefulsets := utils.ListStateFulSet(token, url)
	if resultado == 0 {

		// Ler os dados dos statefulsets
		lerDadosStateFulSets(token, statefulsets)
	} else {
		fmt.Println("[ListarStateFulSet] StateFulSet não encontrados")
	}
}

func lerDadosStateFulSets(token string, statefulsets model.StateFulSets) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(statefulsets.Items); i++ {
		nomeProjeto := statefulsets.Items[i].Metadata.Namespace
		nomeStateFulSet := statefulsets.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarStateFulSet(token, nomeProjeto, nomeStateFulSet)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarStateFulSet(token string, nomeProjeto string, nomeStateFulSet string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirStateFulSet := dirProjeto + "/statefulset"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirStateFulSet, 0700)

	resultado, statefulset := utils.GetStateFulSetString(token, url, nomeProjeto, nomeStateFulSet)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirStateFulSet + "/" + nomeStateFulSet + ".json"
		resultado = SalvarArquivoJSON(arquivo, statefulset)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerStateFulSet] StateFulSet %s não encontrado no projeto %s ambiente %s\n\r", nomeStateFulSet, nomeProjeto, url)
	}
	return recursoSalvo
}
