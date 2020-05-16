package variaveis

import (
	"os"
	"time"
)

//DataFormat formato da data
var DataFormat = "02/01/2006 15:04:05"

//DataFormatArquivo formato da data para arquivos
var DataFormatArquivo = "20060102-150405"

//DataHoraAtual a data e hora tual
var DataHoraAtual = time.Now()

//LogDebug
var LogDebug = os.Getenv("LOG_DEBUG")

//RecursosFile
var RecursosFile = os.Getenv("RECURSOS_FILE")

//DirBase
var DirBase string

//GitlabApiURL URL do gitlab
var GitlabApiURL = os.Getenv("GIT_URL")

//GitlabProjectID ID do projeto no gitlab
var GitlabProjectID = os.Getenv("GITLAB_PROJECT_ID")

//GitlabBranch nome do branch do gitlab
var GitlabBranch = os.Getenv("GIT_BRANCH")

//GitlabToken notoken do usuário do gitlab
var GitlabToken = os.Getenv("GITLAB_PRIVATE_KEY")

//OpenshiftApiURL URL do openshift
var OpenshiftApiURL = os.Getenv("OPENSHIFT_URL")

//OpenshiftUsername usuário do openshift
var OpenshiftUsername = os.Getenv("OPENSHIFT_USERNAME")

//OpenshiftPassword URL do openshift
var OpenshiftPassword = os.Getenv("OPENSHIFT_PASSWORD")

//OpenshiftToken token usuário do openshift
var OpenshiftToken string

//OpenshiftApiApps
var OpenshiftApiApps = "/apis/apps.openshift.io/v1/"

//OpenshiftApiV1
var OpenshiftApiV1 = "/api/v1/"

//OpenshiftApiRoutes
var OpenshiftApiRoutes = "/apis/route.openshift.io/v1/"

//OpenshiftApisAppsv1beta1
var OpenshiftApisAppsv1beta1 = "/apis/apps/v1beta1/"

//OpenshiftApisImageV1
var OpenshiftApisImageV1 = "/apis/image.openshift.io/v1/"

//OpenshiftApisAuthorizationOpenshiftV1
var OpenshiftApisAuthorizationOpenshiftV1 = "/apis/authorization.openshift.io/v1/"

// /OpenshiftApisExtensionsV1beta1
var OpenshiftApisExtensionsV1beta1 = "/apis/extensions/v1beta1/"

//OpenshiftApiBuilds
var OpenshiftApiBuilds = "/apis/build.openshift.io/v1/"

//OpenshiftApiTemplates
var OpenshiftApiTemplates = "/apis/template.openshift.io/v1/"

//OpenshiftApiBatchBeta1
var OpenshiftApiBatchBeta1 = "/apis/batch/v1beta1/"

//GitlabApiProjetos
var GitlabApiProjetos = "/api/v4/projects"
