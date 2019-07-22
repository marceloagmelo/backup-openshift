package variaveis

import "time"

//DataFormat formato da data
var DataFormat = "02/01/2006 15:04:05"

//DataFormatArquivo formato da data para arquivos
var DataFormatArquivo = "20060102-150405"

//var imagens = getSlicesImagensOpenshift()
var imagens []string

//GitRepositorio nome do repositório do git
var GitRepositorio = "https://github.com/marceloagmelo/backup-recursos-openshift.git"

//DataHoraAtual a data e hora tual
var DataHoraAtual = time.Now()

//Ambiente que será executado
var Ambiente = "pre"

//DirBase
var DirBase string

//DirDc
var DirDc string
