package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupServices lista dos services do openshift
func BackupServices(token string, url string) {
	// Listar os services
	fmt.Printf("Backup dos recursos services do ambiente %s\n\r", url)
	resultado, services := utils.ListService(token, url)
	if resultado == 0 {

		// Ler os dados dos services
		lerDadosServices(token, services)
	} else {
		fmt.Println("[ListarServices] Services não encontrados")
	}
}

func lerDadosServices(token string, services model.Services) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(services.Items); i++ {
		nomeProjeto := services.Items[i].Metadata.Namespace
		nomeService := services.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarService(token, nomeProjeto, nomeService)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarService(token string, nomeProjeto string, nomeService string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirService := dirProjeto + "/service"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirService, 0700)

	resultado, service := utils.GetServiceString(token, url, nomeProjeto, nomeService)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirService + "/" + nomeService + ".json"
		resultado = SalvarArquivoJSON(arquivo, service)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerService] Service %s não encontrado no projeto %s ambiente %s\n\r", nomeService, nomeProjeto, url)
	}
	return recursoSalvo
}
