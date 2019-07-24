package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupRoleBindings lista dos rolebindings do openshift
func BackupRoleBindings(token string, url string) {
	// Listar os rolebindings
	fmt.Printf("Backup dos recursos rolebindings do ambiente %s\n\r", url)
	resultado, rolebindings := utils.ListRoleBinding(token, url)
	if resultado == 0 {

		// Ler os dados dos rolebindings
		lerDadosRoleBindings(token, rolebindings)
	} else {
		fmt.Println("[ListarRoleBinding] RoleBinding não encontrados")
	}
}

func lerDadosRoleBindings(token string, rolebindings model.RoleBindings) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(rolebindings.Items); i++ {
		nomeProjeto := rolebindings.Items[i].Metadata.Namespace
		nomeRoleBinding := rolebindings.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarRoleBinding(token, nomeProjeto, nomeRoleBinding)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarRoleBinding(token string, nomeProjeto string, nomeRoleBinding string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirRoleBinding := dirProjeto + "/rolebinding"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirRoleBinding, 0700)

	resultado, rolebinding := utils.GetRoleBindingString(token, url, nomeProjeto, nomeRoleBinding)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirRoleBinding + "/" + nomeRoleBinding + ".json"
		resultado = SalvarArquivoJSON(arquivo, rolebinding)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerRoleBinding] RoleBinding %s não encontrado no projeto %s ambiente %s\n\r", nomeRoleBinding, nomeProjeto, url)
	}
	return recursoSalvo
}
