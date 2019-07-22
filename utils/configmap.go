package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//ListarConfigMaps lista dos configmaps do openshift
func ListarConfigMaps(token string, url string) {
	// Listar os configmaps
	fmt.Printf("Listando todos os configmaps do ambiente %s\n\r", url)
	resultado, configmaps := utils.ListConfigMap(token, url)
	if resultado == 0 {

		// Ler os dados dos configmaps
		lerDadosDadosConfigMaps(token, configmaps)
	} else {
		fmt.Println("[ListarConfigMap] ConfigMap não encontrados")
	}
}

func lerDadosDadosConfigMaps(token string, configmaps model.ConfigMaps) {
	for i := 0; i < len(configmaps.Items); i++ {
		nomeProjeto := configmaps.Items[i].Metadata.Namespace
		nomeConfigMap := configmaps.Items[i].Metadata.Name

		lerConfigMap(token, nomeProjeto, nomeConfigMap)
	}
}

func lerConfigMap(token string, nomeProjeto string, nomeConfigMap string) {
	url := utils.URLGen(variaveis.Ambiente)
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirConfigMap := dirProjeto + "/configmap"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirConfigMap, 0700)

	resultado, configmap := utils.GetConfigMapString(token, url, nomeProjeto, nomeConfigMap)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirConfigMap + "/" + nomeConfigMap + ".json"
		SalvarArquivoJSON(arquivo, configmap)
	} else {
		fmt.Printf("[lerConfigMap] ConfigMap %s não encontrado no projeto %s ambiente %s\n\r", nomeConfigMap, nomeProjeto, url)
	}
}
