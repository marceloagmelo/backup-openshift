package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupTemplates lista dos templates do openshift
func BackupTemplates(token string, url string) {
	// Listar os templates
	fmt.Printf("Backup dos recursos templates do ambiente %s\n\r", url)
	resultado, templates := utils.ListTemplate(token, url)
	if resultado == 0 {

		// Ler os dados dos templates
		lerDadosTemplates(token, templates)
	} else {
		fmt.Println("[BackupTemplates] Templates não encontrados")
	}
}

func lerDadosTemplates(token string, templates model.Templates) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(templates.Items); i++ {
		nomeProjeto := templates.Items[i].Metadata.Namespace
		nomeTemplate := templates.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarTemplate(token, nomeProjeto, nomeTemplate)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarTemplate(token string, nomeProjeto string, nomeTemplate string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirTemplate := dirProjeto + "/template"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirTemplate, 0700)

	resultado, template := utils.GetTemplateString(token, url, nomeProjeto, nomeTemplate)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirTemplate + "/" + nomeTemplate + ".json"
		resultado = SalvarArquivoJSON(arquivo, template)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarTemplate] Template %s não encontrado no projeto %s ambiente %s\n\r", nomeTemplate, nomeProjeto, url)
	}
	return recursoSalvo
}
