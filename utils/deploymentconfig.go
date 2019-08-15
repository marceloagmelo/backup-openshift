package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupDcs lista dos dcs do openshift
func BackupDcs(token string, url string) {
	// Listar os dcs
	fmt.Printf("Backup dos recursos deploymentconfig do ambiente %s\n\r", url)
	resultado, dcs := utils.ListDeploymentConfig(token, url)
	if resultado == 0 {

		// Ler os dados dos dcs
		lerDadosDcs(token, dcs)
	} else {
		fmt.Println("[BackupDcs] Dcs não encontrados")
	}
}

func lerDadosDcs(token string, dcs model.DeploymentConfigs) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(dcs.Items); i++ {
		nomeProjeto := dcs.Items[i].Metadata.Namespace
		nomeDc := dcs.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarDc(token, nomeProjeto, nomeDc)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarDc(token string, nomeProjeto string, nomeDc string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirDc := dirProjeto + "/deploymentconfig"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirDc, 0700)

	resultado, dc := utils.GetDeploymentConfigString(token, url, nomeProjeto, nomeDc)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirDc + "/" + nomeDc + ".json"
		resultado = SalvarArquivoJSON(arquivo, dc)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarDc] Dc %s não encontrado no projeto %s ambiente %s", nomeDc, nomeProjeto, url)
	}
	return recursoSalvo
}
