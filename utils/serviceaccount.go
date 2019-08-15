package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupServiceAccounts lista dos serviceaccounts do openshift
func BackupServiceAccounts(token string, url string) {
	// Listar os serviceaccounts
	fmt.Printf("Backup dos recursos serviceaccounts do ambiente %s\n\r", url)
	resultado, serviceaccounts := utils.ListServiceAccount(token, url)
	if resultado == 0 {

		// Ler os dados dos serviceaccounts
		lerDadosServiceAccounts(token, serviceaccounts)
	} else {
		fmt.Println("[BackupServiceAccounts] ServiceAccount não encontrados")
	}
}

func lerDadosServiceAccounts(token string, serviceaccounts model.ServiceAccounts) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(serviceaccounts.Items); i++ {
		nomeProjeto := serviceaccounts.Items[i].Metadata.Namespace
		nomeServiceAccount := serviceaccounts.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarServiceAccount(token, nomeProjeto, nomeServiceAccount)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarServiceAccount(token string, nomeProjeto string, nomeServiceAccount string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirServiceAccount := dirProjeto + "/serviceaccount"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirServiceAccount, 0700)

	resultado, serviceaccount := utils.GetServiceAccountString(token, url, nomeProjeto, nomeServiceAccount)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirServiceAccount + "/" + nomeServiceAccount + ".json"
		resultado = SalvarArquivoJSON(arquivo, serviceaccount)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarServiceAccount] ServiceAccount %s não encontrado no projeto %s ambiente %s\n\r", nomeServiceAccount, nomeProjeto, url)
	}
	return recursoSalvo
}
