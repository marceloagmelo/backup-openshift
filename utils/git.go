package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

//GitClone do Git
func GitClone(repositorio string, diretorio string) {

	username, password := GetUsuarioSenha()

	fmt.Printf("git clone %s %s --recursive\n\r", repositorio, diretorio)

	r, err := git.PlainClone(diretorio, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:      repositorio,
		Progress: os.Stdout,
	})

	// Recuperando o branch apontado por HEAD
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	// Recuperando o objeto commit
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commit)
}

//GitCommitPush Commit e Push no git
func GitCommitPush(diretorio string, dataFormatada string) {
	username, password := GetUsuarioSenha()

	// Abrir um repositório já existente
	r, err := git.PlainOpen(diretorio)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// Adds the new file to the staging area.
	fmt.Println("git add .")
	_, err = w.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	// We can verify the current status of the worktree using the method Status.
	fmt.Println("git status --porcelain")
	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status)

	// Commit dos arquivos
	fmt.Printf("git commit -m %s\n\r", "Backup-"+dataFormatada)
	commit, err := w.Commit("Backup-"+dataFormatada, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Marcelo Melo",
			Email: "marceloagmelo@gmail.com",
			When:  time.Now(),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Prints the current HEAD to verify that all worked well.
	fmt.Println("git show -s")
	obj, err := r.CommitObject(commit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj)

	fmt.Println("git push")
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		}})
	if err != nil {
		log.Fatal(err)
	}
}

//GitCriarBranch Criar um branch
func GitCriarBranch(diretorio string, dataFormatada string) {
	username, password := GetUsuarioSenha()
	nomeBranch := "backup-" + dataFormatada

	// Abrir um repositório já existente
	r, err := git.PlainOpen(diretorio)
	if err != nil {
		log.Fatal(err)
	}

	// Criar um branch do HEAD
	fmt.Printf("git branch %s\n\r", nomeBranch)

	headRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	// Referenciar um nome para o branch
	refName := plumbing.NewBranchReferenceName(nomeBranch)
	ref := plumbing.NewHashReference(refName, headRef.Hash())

	// A referencia crida e salva e armazenada
	err = r.Storer.SetReference(ref)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("git push")
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		}})
	if err != nil {
		log.Fatal(err)
	}
}
