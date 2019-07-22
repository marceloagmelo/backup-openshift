package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarRoutes lista dos routes do openshift
func ListarRoutes(token string, url string) {
	// Listar os routes
	fmt.Printf("Listando todos os routes do ambiente %s\n\r", url)
	resultado, routes := utils.ListRoute(token, url)
	if resultado == 0 {

		// Ler os dados dos routes
		lerDadosDadosRoutes(token, routes)
	} else {
		fmt.Println("[ListarRoutes] Routes não encontrados")
	}
}

func lerDadosDadosRoutes(token string, routes model.Routes) {
	for i := 0; i < len(routes.Items); i++ {
		nomeProjeto := routes.Items[i].Metadata.Namespace
		nomeRoute := routes.Items[i].Metadata.Name

		lerRoute(token, nomeProjeto, nomeRoute)
	}
}

func lerRoute(token string, nomeProjeto string, nomeRoute string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirRoute := dirProjeto + "/route"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirRoute, 0700)

	resultado, route := utils.GetRouteString(token, url, nomeProjeto, nomeRoute)
	if resultado == 0 {
		// Salvar o arquivo de Route
		arquivo := dirRoute + "/" + nomeRoute + ".json"
		SalvarArquivoJSON(arquivo, route)
	} else {
		fmt.Printf("[lerRoute] route %s não encontrado no projeto %s ambiente %s\n\r", nomeRoute, nomeProjeto, url)
	}
}
