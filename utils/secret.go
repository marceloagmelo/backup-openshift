package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupSecrets lista dos dcs do openshift
func BackupSecrets(token string, url string) {
	// Listar os secrets
	fmt.Printf("Backup dos recursos secrets do ambiente %s\n\r", url)
	resultado, secrets := utils.ListSecret(token, url)
	if resultado == 0 {

		// Ler os dados dos secrets
		lerDadosSecrets(token, secrets)
	} else {
		fmt.Println("[BackupSecrets] Secrets não encontrados")
	}
}

func lerDadosSecrets(token string, secrets model.Secrets) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(secrets.Items); i++ {
		nomeProjeto := secrets.Items[i].Metadata.Namespace
		nomeSecret := secrets.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarSecret(token, nomeProjeto, nomeSecret)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarSecret(token string, nomeProjeto string, nomeSecret string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirSecret := dirProjeto + "/secret"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirSecret, 0700)

	resultado, secret := utils.GetSecretString(token, url, nomeProjeto, nomeSecret)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirSecret + "/" + nomeSecret + ".json"
		resultado = SalvarArquivoJSON(arquivo, secret)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarSecret] secret %s não encontrada no projeto %s ambiente %s\n\r", nomeSecret, nomeProjeto, url)
	}
	return recursoSalvo
}
