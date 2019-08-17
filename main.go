package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/marceloagmelo/go-backup-openshift/utils"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	gitutils "github.com/marceloagmelo/go-git-cli/utils"
	openshiftutils "github.com/marceloagmelo/go-openshift-cli/utils"
)

func init() {
	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Inicio -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}

func main() {
	url := os.Getenv("OPENSHIFT_URL")
	openshiftUsername := os.Getenv("OPENSHIFT_USERNAME")
	openshiftPassword := os.Getenv("OPENSHIFT_PASSWORD")
	gitUsername := os.Getenv("GIT_USERNAME")
	gitPassword := os.Getenv("GIT_PASSWORD")
	gitRepositorio := os.Getenv("GIT_REPOSITORIO")
	limparRecursos := os.Getenv("LIMPAR_RECURSOS")
	if runtime.GOOS == "windows" {
		limparRecursos = strings.TrimRight(limparRecursos, "\r\n")
	} else {
		limparRecursos = strings.TrimRight(limparRecursos, "\n")
	}

	resultado, token := openshiftutils.GetToken(url, openshiftUsername, openshiftPassword)
	if resultado > 0 {
		fmt.Println("[main] Token não recuperado.")
	}

	if len(strings.TrimSpace(token)) > 0 {
		// Atribuir valores as variáveis
		dataFormatada := variaveis.DataHoraAtual.Format(variaveis.DataFormatArquivo)
		variaveis.DirBase = "/tmp/backup-openshift-" + dataFormatada

		// Criar diretórios
		fmt.Printf("criando diretorio %s\n\r", variaveis.DirBase)
		os.Mkdir(variaveis.DirBase, 0700)

		// Clonar o respositório de backup
		gitutils.GitClone(gitRepositorio, variaveis.DirBase, gitUsername, gitPassword)

		// Verficiar se precisa limpar recursos antes
		if len(limparRecursos) > 0 {
			limparRecursos = strings.ToUpper(limparRecursos)
			if limparRecursos == "S" {
				// Copiar a pasta .git para pasta temporária
				utils.CopiarPastaGit(dataFormatada)
			}
		}

		// Backup dos Projetos
		utils.BackupProjetos(token, url)

		// Backup dos Templates
		utils.BackupTemplates(token, url)

		// Backup dos RoleBindings
		utils.BackupRoleBindings(token, url)

		// Backup dos Roles
		utils.BackupRoles(token, url)

		// Backup dos Services
		utils.BackupServices(token, url)

		// Backup dos ServiceAccounts
		utils.BackupServiceAccounts(token, url)

		// Backup dos Routes
		utils.BackupRoutes(token, url)

		// Backup dos Secrets
		utils.BackupSecrets(token, url)

		// Backup dos Pvcs
		utils.BackupPvcs(token, url)

		// Backup dos ConfigMaps
		utils.BackupConfigMaps(token, url)

		// Recuperar as ImagesStreams
		utils.BackupImageStreams(token, url)

		// Backup dos Bcs
		utils.BackupBcs(token, url)

		// Backup dos Dcs
		utils.BackupDcs(token, url)

		// Backup dos StateFulSets
		utils.BackupStateFulSets(token, url)

		// Backup dos DaemonSets
		utils.BackupDaemonSets(token, url)

		// Backup dos ReplicaSets
		utils.BackupReplicaSets(token, url)

		// Backup dos LimitRanges
		utils.BackupLimitRanges(token, url)

		// Backup dos ResourceQuotas
		utils.BackupResourceQuotas(token, url)

		// Commit e push no git
		gitutils.GitCommitPush(variaveis.DirBase, "Backup "+dataFormatada, gitUsername, gitPassword)
		// Criar branch no git
		gitutils.GitCriarBranch(variaveis.DirBase, dataFormatada, gitUsername, gitPassword)

		// Remover o diretório base
		fmt.Printf("removendo diretorio %s\n\r", variaveis.DirBase)
		os.RemoveAll(variaveis.DirBase)
	}

	variaveis.DataHoraAtual = time.Now()
	fmt.Println("Fim -->> ", variaveis.DataHoraAtual.Format(variaveis.DataFormat))
}
