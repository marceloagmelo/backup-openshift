# Backup dos Recursos do Openshift

Aplicação que realiza o backup de recursos do **Openshift**, este recursos serão armazenados no repositório do [GITLAB](https://gitlab.com/marceloagmelo/openshift-backup.git) **ID 18811050** serão feitos o backup dos seguintes recursos:

- buildconfig
- configmap
- cronjob
- daemonset
- deployment
- deploymentconfig
- namespace
- imagestream
- replicaSet
- rolebinding
- route
- secret
- service
- serviceaccount
- statefulset
- role
- rolebinding
- template

----

# Instalação

```
go get -v github.com/marceloagmelo/go-backup-openshift
```
```
cd go-backup-openshift
```

## Build da Aplicação

```
./image-build.sh
```

# Instalação no Openshift


Importar o [Template](https://github.com/marceloagmelo/go-backup-openshift/blob/master/openshift/template/go-backup-openshift-template.json) no projeto do openshift e preencher as seguintes informações:

```
Application Name: apagar-templates-default-openshift
Openshift URL: https://console.openshift.lab:8443
Openshift Username: 
Openshift Password:
Gitlab URL: https://gitlab.com
Gitlab Repositório: https://gitlab.com/marceloagmelo/openshift-backup.git
Gitlab Branch: master
Gitlab Username: 
Gitlab Password:
Gitlab Private Key:
Gitlab Project ID: 11027
Caminho do arquivo de recursos: /go/bin/recursosValidos.json
Limpar Recursos: (S ou N)
Execucação do Job Ativada: (S ou N)
Log DEBUG: (S ou N)
Schedule: ( 0 2 * * * ) -> Todos os dias as 2hs da manhã
```
