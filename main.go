package main

import (
	"fmt"
	"os"
	"time"

	"github.com/marceloagmelo/backup-openshift/utils"
	"github.com/marceloagmelo/backup-openshift/variaveis"
	openshiftutils "github.com/marceloagmelo/go-openshift-cli/utils"
)

func init() {
	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Inicio -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}

func main() {
	url := openshiftutils.URLGen(variaveis.Ambiente)
	token := openshiftutils.GetToken(url)

	// Atribuir valores as variáveis
	dataFormatada := variaveis.DataHoraAtual.Format(variaveis.DataFormatArquivo)
	variaveis.DirBase = "/tmp/backup-" + variaveis.Ambiente + "-" + dataFormatada

	// Criar diretórios
	os.Mkdir(variaveis.DirBase, 0700)

	// Clonar o respositório de backup
	utils.GitClone(variaveis.GitRepositorio, variaveis.DirBase)

	// Recuperar os Services
	utils.ListarServices(token, url)

	// Recuperar os Routes
	utils.ListarRoutes(token, url)

	// Recuperar os Secrets
	utils.ListarSecrets(token, url)

	// Recuperar os Pvcs
	utils.ListarPvcs(token, url)

	// Recuperar os ConfigMaps
	utils.ListarConfigMaps(token, url)

	// Recuperar os Dcs
	utils.ListarDcs(token, url)

	// Recuperar os StateFulSets
	utils.ListarStateFulSets(token, url)

	// Commit e push no git
	utils.GitCommitPush(variaveis.DirBase, dataFormatada)

	// Criar branch no git
	utils.GitCriarBranch(variaveis.DirBase, dataFormatada)

	// Remover o diretório base
	os.Remove(variaveis.DirBase)

	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Fim -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}
