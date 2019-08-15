package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupConfigMaps lista dos configmaps do openshift
func BackupConfigMaps(token string, url string) {
	// Listar os configmaps
	fmt.Printf("Backup dos recursos configmap do ambiente %s\n\r", url)
	resultado, configmaps := utils.ListConfigMap(token, url)
	if resultado == 0 {

		// Ler os dados dos configmaps
		lerDadosConfigMaps(token, configmaps)
	} else {
		fmt.Println("[BackupConfigMaps] ConfigMap não encontrados")
	}
}

func lerDadosConfigMaps(token string, configmaps model.ConfigMaps) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(configmaps.Items); i++ {
		nomeProjeto := configmaps.Items[i].Metadata.Namespace
		nomeConfigMap := configmaps.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarConfigMap(token, nomeProjeto, nomeConfigMap)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarConfigMap(token string, nomeProjeto string, nomeConfigMap string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirConfigMap := dirProjeto + "/configmap"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirConfigMap, 0700)

	resultado, configmap := utils.GetConfigMapString(token, url, nomeProjeto, nomeConfigMap)
	if resultado == 0 {
		// Salvar o arquivo
		arquivo := dirConfigMap + "/" + nomeConfigMap + ".json"
		resultado = SalvarArquivoJSON(arquivo, configmap)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerSalvarConfigMap] ConfigMap %s não encontrado no projeto %s ambiente %s\n\r", nomeConfigMap, nomeProjeto, url)
	}
	return recursoSalvo
}
