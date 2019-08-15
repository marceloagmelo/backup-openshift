package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupRoles lista dos roles do openshift
func BackupRoles(token string, url string) {
	// Listar os roles
	fmt.Printf("Backup dos recursos roles do ambiente %s\n\r", url)
	resultado, roles := utils.ListRole(token, url)
	if resultado == 0 {

		// Ler os dados dos roles
		lerDadosRoles(token, roles)
	} else {
		fmt.Println("[BackupRoles] Role não encontrados")
	}
}

func lerDadosRoles(token string, roles model.Roles) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(roles.Items); i++ {
		nomeProjeto := roles.Items[i].Metadata.Namespace
		nomeRole := roles.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarRole(token, nomeProjeto, nomeRole)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarRole(token string, nomeProjeto string, nomeRole string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirRole := dirProjeto + "/role"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirRole, 0700)

	resultado, role := utils.GetRoleString(token, url, nomeProjeto, nomeRole)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirRole + "/" + nomeRole + ".json"
		resultado = SalvarArquivoJSON(arquivo, role)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarRole] Role %s não encontrado no projeto %s ambiente %s\n\r", nomeRole, nomeProjeto, url)
	}
	return recursoSalvo
}
