package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupResourceQuotas lista dos resourceQuotas do openshift
func BackupResourceQuotas(token string, url string) {
	// Listar os resourceQuotas
	fmt.Printf("Backup dos recursos resourcequota do ambiente %s\n\r", url)
	resultado, resourceQuotas := utils.ListResourceQuota(token, url)
	if resultado == 0 {

		// Ler os dados dos resourcequotas
		lerDadosResourceQuotas(token, resourceQuotas)
	} else {
		fmt.Println("[BackupResourceQuotas] ResourceQuota não encontrados")
	}
}

func lerDadosResourceQuotas(token string, resourceQuotas model.ResourceQuotas) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(resourceQuotas.Items); i++ {
		nomeProjeto := resourceQuotas.Items[i].Metadata.Namespace
		nomeResourceQuota := resourceQuotas.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarResourceQuota(token, nomeProjeto, nomeResourceQuota)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarResourceQuota(token string, nomeProjeto string, nomeResourceQuota string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirResourceQuota := dirProjeto + "/resourcequota"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirResourceQuota, 0700)

	resultado, resourceQuota := utils.GetResourceQuotaString(token, url, nomeProjeto, nomeResourceQuota)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirResourceQuota + "/" + nomeResourceQuota + ".json"
		resultado = SalvarArquivoJSON(arquivo, resourceQuota)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarResourceQuota] ResourceQuota %s não encontrado no projeto %s ambiente %s\n\r", nomeResourceQuota, nomeProjeto, url)
	}
	return recursoSalvo
}
