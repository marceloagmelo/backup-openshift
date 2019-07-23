package main

import (
	"fmt"
	"os"
	"time"

	"github.com/marceloagmelo/backup-openshift/utils"
	"github.com/marceloagmelo/backup-openshift/variaveis"
	gitutils "github.com/marceloagmelo/go-git-cli/utils"
	openshiftutils "github.com/marceloagmelo/go-openshift-cli/utils"
)

func init() {
	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Inicio -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}

func main() {
	url := openshiftutils.URLGen(variaveis.Ambiente)
	token := openshiftutils.GetToken(url)
	gitRepositorio := os.Getenv("GIT_REPOSITORIO")

	// Atribuir valores as variáveis
	dataFormatada := variaveis.DataHoraAtual.Format(variaveis.DataFormatArquivo)
	variaveis.DirBase = "/tmp/backup-" + variaveis.Ambiente + "-" + dataFormatada

	// Criar diretórios
	os.Mkdir(variaveis.DirBase, 0700)

	// Clonar o respositório de backup
	gitutils.GitClone(gitRepositorio, variaveis.DirBase)

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
	gitutils.GitCommitPush(variaveis.DirBase, "Backup-"+dataFormatada)

	// Criar branch no git
	gitutils.GitCriarBranch(variaveis.DirBase, "Backup-"+dataFormatada)

	// Remover o diretório base
	fmt.Printf("removendo diretorio %s\n\r", variaveis.DirBase)
	os.RemoveAll(variaveis.DirBase)

	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Fim -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}
