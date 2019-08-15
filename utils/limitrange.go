package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupLimitRanges lista dos limitranges do openshift
func BackupLimitRanges(token string, url string) {
	// Listar os limitranges
	fmt.Printf("Backup dos recursos limitrange do ambiente %s\n\r", url)
	resultado, limitranges := utils.ListLimitRange(token, url)
	if resultado == 0 {

		// Ler os dados dos limitranges
		lerDadosLimitRanges(token, limitranges)
	} else {
		fmt.Println("[BackupLimitRanges] LimitRange não encontrados")
	}
}

func lerDadosLimitRanges(token string, limitranges model.LimitRanges) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(limitranges.Items); i++ {
		nomeProjeto := limitranges.Items[i].Metadata.Namespace
		nomeLimitRange := limitranges.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarLimitRange(token, nomeProjeto, nomeLimitRange)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarLimitRange(token string, nomeProjeto string, nomeLimitRange string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirLimitRange := dirProjeto + "/limitrange"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirLimitRange, 0700)

	resultado, limitrange := utils.GetLimitRangeString(token, url, nomeProjeto, nomeLimitRange)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirLimitRange + "/" + nomeLimitRange + ".json"
		resultado = SalvarArquivoJSON(arquivo, limitrange)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarLimitRange] LimitRange %s não encontrado no projeto %s ambiente %s\n\r", nomeLimitRange, nomeProjeto, url)
	}
	return recursoSalvo
}
