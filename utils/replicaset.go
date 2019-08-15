package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupReplicaSets lista dos replicaSets do openshift
func BackupReplicaSets(token string, url string) {
	// Listar os replicaSets
	fmt.Printf("Backup dos recursos replicaSets do ambiente %s\n\r", url)
	resultado, replicaSets := utils.ListReplicaSet(token, url)
	if resultado == 0 {

		// Ler os dados dos replicaSets
		lerDadosReplicaSets(token, replicaSets)
	} else {
		fmt.Println("[BackupReplicaSets] ReplicaSet não encontrados")
	}
}

func lerDadosReplicaSets(token string, replicaSets model.ReplicaSets) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(replicaSets.Items); i++ {
		nomeProjeto := replicaSets.Items[i].Metadata.Namespace
		nomeReplicaSet := replicaSets.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarReplicaSet(token, nomeProjeto, nomeReplicaSet)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarReplicaSet(token string, nomeProjeto string, nomeReplicaSet string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirReplicaSet := dirProjeto + "/replicaSet"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirReplicaSet, 0700)

	resultado, replicaSet := utils.GetReplicaSetString(token, url, nomeProjeto, nomeReplicaSet)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirReplicaSet + "/" + nomeReplicaSet + ".json"
		resultado = SalvarArquivoJSON(arquivo, replicaSet)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarReplicaSet] ReplicaSet %s não encontrado no projeto %s ambiente %s\n\r", nomeReplicaSet, nomeProjeto, url)
	}
	return recursoSalvo
}
