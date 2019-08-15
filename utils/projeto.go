package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupProjetos lista dos projetos do openshift
func BackupProjetos(token string, url string) {
	// Listar os projetos
	fmt.Printf("Backup dos recursos projetos do ambiente %s\n\r", url)
	resultado, projetos := utils.Namespaces(token, url)
	if resultado == 0 {

		// Ler os dados dos projetos
		lerDadosProjetos(token, projetos)
	} else {
		fmt.Println("[BackupProjetos] Projetos não encontrados")
	}
}

func lerDadosProjetos(token string, projetos model.Projetos) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(projetos.Items); i++ {
		nomeProjeto := projetos.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarProjeto(token, nomeProjeto)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarProjeto(token string, nomeProjeto string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirNamespace := dirProjeto + "/namespace"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirNamespace, 0700)

	resultado, projeto := utils.GetNamespaceString(token, url, nomeProjeto)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirNamespace + "/" + nomeProjeto + ".json"
		resultado = SalvarArquivoJSON(arquivo, projeto)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarProjeto] Projeto %s não encontrado no projeto %s ambiente %s\n\r", nomeProjeto, nomeProjeto, url)
	}
	return recursoSalvo
}
