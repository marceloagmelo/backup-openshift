package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarSecrets lista dos dcs do openshift
func ListarSecrets(token string, url string) {
	// Listar os secrets
	fmt.Printf("Listando todos as secrets do ambiente %s\n\r", url)
	resultado, secrets := utils.ListSecret(token, url)
	if resultado == 0 {

		// Ler os dados dos secrets
		lerDadosDadosSecrets(token, secrets)
	} else {
		fmt.Println("[ListarSecrets] Secrets não encontrados")
	}
}

func lerDadosDadosSecrets(token string, secrets model.Secrets) {
	for i := 0; i < len(secrets.Items); i++ {
		nomeProjeto := secrets.Items[i].Metadata.Namespace
		nomeSecret := secrets.Items[i].Metadata.Name

		lerSecret(token, nomeProjeto, nomeSecret)
	}
}

func lerSecret(token string, nomeProjeto string, nomeSecret string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirSecret := dirProjeto + "/secret"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirSecret, 0700)

	resultado, secret := utils.GetSecretString(token, url, nomeProjeto, nomeSecret)
	if resultado == 0 {
		// Salvar o arquivo de secret
		arquivo := dirSecret + "/" + nomeSecret + ".json"
		SalvarArquivoJSON(arquivo, secret)
	} else {
		fmt.Printf("[lerSecret] secret %s não encontrada no projeto %s ambiente %s\n\r", nomeSecret, nomeProjeto, url)
	}
}
