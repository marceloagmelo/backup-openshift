package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/marceloagmelo/go-backup-openshift/logger"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

//GitClone do Git
func GitClone(repositorio string, diretorio string, username string, password string, branch string) error {

	mensagem := fmt.Sprintf("git clone %s %s branch %s --recursive", repositorio, diretorio, branch)
	logger.Info.Println(mensagem)

	// Create a custom http(s) client with your config
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	customClient := &http.Client{
		// accept any certificate (might be useful for testing)
		Transport: tr,

		// 15 second timeout
		Timeout: 120 * time.Second,

		// don't follow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Override http(s) default protocol to use our custom client
	client.InstallProtocol("https", githttp.NewClient(customClient))

	r, err := git.PlainClone(diretorio, false, &git.CloneOptions{
		Auth: &githttp.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:      repositorio,
		Progress: os.Stdout,
		//ReferenceName: plumbing.ReferenceName(branch),
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		SingleBranch:  true,
	})
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "git.PlainClone", err)
		logger.Erro.Println(mensagem)
		return err
	}

	// Recuperando o branch apontado por HEAD
	ref, err := r.Head()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "r.Head()", err)
		logger.Erro.Println(mensagem)
		return err
	}
	// Recuperando o objeto commit
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "r.CommitObject", err)
		logger.Erro.Println(mensagem)
		return err
	}

	logger.Info.Println(commit)
	return nil
}

//GitCommitPush Commit e Push no git
func GitCommitPush(diretorio string, mensagemCommit string, username string, password string) error {

	// Abrir um repositório já existente
	r, err := git.PlainOpen(diretorio)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "git.PlainOpen", err)
		logger.Erro.Println(mensagem)
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "r.Worktree", err)
		logger.Erro.Println(mensagem)
		return err
	}

	// Adds the new file to the staging area.
	logger.Info.Println("git add .")
	_, err = w.Add(".")
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "w.Add", err)
		logger.Erro.Println(mensagem)
		return err
	}

	// We can verify the current status of the worktree using the method Status.
	logger.Info.Println("git status --porcelain")
	status, err := w.Status()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "w.Status", err)
		logger.Erro.Println(mensagem)
		return err
	}

	logger.Info.Println(status)

	// Commit dos arquivos
	logger.Info.Printf("git commit -m %s\n\r", mensagemCommit)
	commit, err := w.Commit(mensagemCommit, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Marcelo Melo",
			Email: "marceloagmelo@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "w.Commit", err)
		logger.Erro.Println(mensagem)
		return err
	}

	// Prints the current HEAD to verify that all worked well.
	logger.Info.Println("git show -s")
	obj, err := r.CommitObject(commit)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "r.CommitObject", err)
		logger.Erro.Println(mensagem)
		return err
	}
	logger.Info.Println(obj)

	logger.Info.Println("git push")
	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		Auth: &githttp.BasicAuth{
			Username: username,
			Password: password,
		}})
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "r.Push", err)
		logger.Erro.Println(mensagem)
		return err
	}
	return nil
}

//GitCriarTag
func GitCriarTag(nome string) (interface{}, int, error) {
	var interf interface{}

	var apiRequest ApiRequest
	apiRequest.Headers = montarHeaderGitlab()
	apiRequest.EndPoint = variaveis.GitlabApiURL + variaveis.GitlabApiProjetos + "/" + variaveis.GitlabProjectID + "/repository/tags?tag_name=" + nome + "&ref=" + variaveis.GitlabBranch
	apiRequest.Metodo = "POST"

	interf, statusCode, err := ExecutarRequest(apiRequest)
	if err != nil {
		return interf, statusCode, err
	}
	return interf, statusCode, nil
}

//getHeader
func montarHeaderGitlab() []Header {
	var headers []Header
	header := Header{}

	header.Chave = "PRIVATE-TOKEN"
	header.Valor = strings.TrimSuffix(variaveis.GitlabToken, "\n")
	headers = append(headers, header)

	return headers
}
