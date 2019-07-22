package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarStateFulSets lista dos statefulsets do openshift
func ListarStateFulSets(token string, url string) {
	// Listar os statefulsets
	fmt.Printf("Listando todos os statefulsets do ambiente %s\n\r", url)
	resultado, statefulsets := utils.ListStateFulSet(token, url)
	if resultado == 0 {

		// Ler os dados dos statefulsets
		lerDadosDadosStateFulSets(token, statefulsets)
	} else {
		fmt.Println("[ListarStateFulSet] StateFulSet não encontrados")
	}
}

func lerDadosDadosStateFulSets(token string, statefulsets model.StateFulSets) {
	for i := 0; i < len(statefulsets.Items); i++ {
		nomeProjeto := statefulsets.Items[i].Metadata.Namespace
		nomeStateFulSet := statefulsets.Items[i].Metadata.Name

		lerStateFulSet(token, nomeProjeto, nomeStateFulSet)
	}
}

func lerStateFulSet(token string, nomeProjeto string, nomeStateFulSet string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirStateFulSet := dirProjeto + "/statefulset"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirStateFulSet, 0700)

	resultado, statefulset := utils.GetStateFulSetString(token, url, nomeProjeto, nomeStateFulSet)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirStateFulSet + "/" + nomeStateFulSet + ".json"
		SalvarArquivoJSON(arquivo, statefulset)
	} else {
		fmt.Printf("[lerStateFulSet] StateFulSet %s não encontrado no projeto %s ambiente %s\n\r", nomeStateFulSet, nomeProjeto, url)
	}
}
