package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupDaemonSets lista dos daemonsets do openshift
func BackupDaemonSets(token string, url string) {
	// Listar os daemonsets
	fmt.Printf("Backup dos recursos daemonset do ambiente %s\n\r", url)
	resultado, daemonsets := utils.ListDaemonSet(token, url)
	if resultado == 0 {

		// Ler os dados dos daemonsets
		lerDadosDaemonSets(token, daemonsets)
	} else {
		fmt.Println("[ListarDaemonSet] DaemonSet não encontrados")
	}
}

func lerDadosDaemonSets(token string, daemonsets model.DaemonSets) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(daemonsets.Items); i++ {
		nomeProjeto := daemonsets.Items[i].Metadata.Namespace
		nomeDaemonSet := daemonsets.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarDaemonSet(token, nomeProjeto, nomeDaemonSet)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarDaemonSet(token string, nomeProjeto string, nomeDaemonSet string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirDaemonSet := dirProjeto + "/daemonset"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirDaemonSet, 0700)

	resultado, daemonset := utils.GetDaemonSetString(token, url, nomeProjeto, nomeDaemonSet)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirDaemonSet + "/" + nomeDaemonSet + ".json"
		resultado = SalvarArquivoJSON(arquivo, daemonset)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerDaemonSet] DaemonSet %s não encontrado no projeto %s ambiente %s\n\r", nomeDaemonSet, nomeProjeto, url)
	}
	return recursoSalvo
}
