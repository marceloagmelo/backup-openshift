package utils

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/backup-openshift/variaveis"
	"github.com/marceloagmelo/go-openshift-cli/model"
	"github.com/marceloagmelo/go-openshift-cli/utils"
)

//BackupImageStreams lista dos ImageStreams do openshift
func BackupImageStreams(token string, url string) {
	// Listar os ImageStreams
	fmt.Printf("Backup dos recursos imagestreams do ambiente %s\n\r", url)
	resultado, ImageStreams := utils.ListImageStream(token, url)
	if resultado == 0 {

		// Ler os dados dos ImageStreams
		lerDadosImageStreams(token, ImageStreams)
	} else {
		fmt.Println("[ListarImageStream] ImageStream não encontradas")
	}
}

func lerDadosImageStreams(token string, ImageStreams model.ImageStreams) {
	quantidadeRecursoSalvo := 0
	for i := 0; i < len(ImageStreams.Items); i++ {
		nomeProjeto := ImageStreams.Items[i].Metadata.Namespace
		nomeImageStream := ImageStreams.Items[i].Metadata.Name

		quantidadeRecursoSalvo = quantidadeRecursoSalvo + lerSalvarImageStream(token, nomeProjeto, nomeImageStream)
	}
	fmt.Printf("Quantidade dos recursos salvos = %d\n\r", quantidadeRecursoSalvo)
}

func lerSalvarImageStream(token string, nomeProjeto string, nomeImageStream string) (recursoSalvo int) {
	recursoSalvo = 0
	url := os.Getenv("OPENSHIFT_URL")
	dirProjeto := variaveis.DirBase + "/" + nomeProjeto
	dirImageStream := dirProjeto + "/imagestream"

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirImageStream, 0700)

	resultado, ImageStream := utils.GetImageStreamString(token, url, nomeProjeto, nomeImageStream)
	if resultado == 0 {
		// Salvar o arquivo de DC
		arquivo := dirImageStream + "/" + nomeImageStream + ".json"
		resultado = SalvarArquivoJSON(arquivo, ImageStream)
		if resultado == 0 {
			recursoSalvo = 1
		}
	} else {
		fmt.Printf("[lerImageStream] ImageStream %s não encontrada no projeto %s ambiente %s\n\r", nomeImageStream, nomeProjeto, url)
	}
	return recursoSalvo
}
