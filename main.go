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
	url := os.Getenv("OPENSHIFT_URL")
	username := os.Getenv("OPENSHIFT_USERNAME")
	password := os.Getenv("OPENSHIFT_PASSWORD")

	token := openshiftutils.GetToken(url, username, password)
	gitRepositorio := os.Getenv("GIT_REPOSITORIO")

	// Atribuir valores as variáveis
	dataFormatada := variaveis.DataHoraAtual.Format(variaveis.DataFormatArquivo)
	variaveis.DirBase = "/tmp/backup-openshift-" + dataFormatada

	// Criar diretórios
	os.Mkdir(variaveis.DirBase, 0700)

	// Clonar o respositório de backup
	gitutils.GitClone(gitRepositorio, variaveis.DirBase)

	// Recuperar os RoleBindings
	utils.BackupRoleBindings(token, url)

	// Recuperar os Services
	utils.BackupServices(token, url)

	// Recuperar os ServiceAccounts
	utils.BackupServiceAccounts(token, url)

	// Recuperar os Routes
	utils.BackupRoutes(token, url)

	// Recuperar os Secrets
	utils.BackupSecrets(token, url)

	// Recuperar os Pvcs
	utils.BackupPvcs(token, url)

	// Recuperar os ConfigMaps
	utils.BackupConfigMaps(token, url)

	// Recuperar as ImagesStreams
	utils.BackupImageStreams(token, url)

	// Recuperar os Bcs
	utils.BackupBcs(token, url)

	// Recuperar os Dcs
	utils.BackupDcs(token, url)

	// Recuperar os StateFulSets
	utils.BackupStateFulSets(token, url)

	// Recuperar os DaemonSets
	utils.BackupDaemonSets(token, url)

	// Commit e push no git
	//gitutils.GitCommitPush(variaveis.DirBase, "Backup-"+dataFormatada)

	// Criar branch no git
	//gitutils.GitCriarBranch(variaveis.DirBase, "Backup-"+dataFormatada)

	// Remover o diretório base
	fmt.Printf("removendo diretorio %s\n\r", variaveis.DirBase)
	os.RemoveAll(variaveis.DirBase)

	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Fim -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}
