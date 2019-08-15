package utils

import (
	"fmt"
	"os"

	"gitlab.produbanbr.corp/paas-brasil/go-backup-openshift/variaveis"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/model"
	"gitlab.produbanbr.corp/paas-brasil/go-openshift-cli/utils"
)

//BackupRoutes lista dos routes do openshift
func BackupRoutes(token string, url string) {
	// Listar os routes
	fmt.Printf("Backup dos recursos routes do ambiente %s\n\r", url)
	resultado, routes := utils.ListRoute(token, url)
	if resultado == 0 {

		// Ler os dados dos routes
		lerDadosRoutes(token, routes)
	} else {
		fmt.Println("[BackupRoutes] Routes não encontrados")
	}
}

func lerDadosRoutes(token string, routes model.Routes) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(routes.Items); i++ {
		nomeProjeto := routes.Items[i].Metadata.Namespace
		nomeRoute := routes.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarRoute(token, nomeProjeto, nomeRoute)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarRoute(token string, nomeProjeto string, nomeRoute string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirRoute := dirProjeto + "/route"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirRoute, 0700)

	resultado, route := utils.GetRouteString(token, url, nomeProjeto, nomeRoute)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirRoute + "/" + nomeRoute + ".json"
		resultado = SalvarArquivoJSON(arquivo, route)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarRoute] route %s não encontrado no projeto %s ambiente %s\n\r", nomeRoute, nomeProjeto, url)
	}
	return recursoSalvo
}
