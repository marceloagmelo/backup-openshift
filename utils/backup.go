package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/marceloagmelo/go-backup-openshift/logger"
	"github.com/marceloagmelo/go-backup-openshift/model"
	"github.com/marceloagmelo/go-backup-openshift/openshift/resource"
	"github.com/marceloagmelo/go-backup-openshift/variaveis"
)

//ExecutarBackup o backup
func ExecutarBackup(recursosValidos model.RecursosValidos) error {

	// Listar namespaces
	recursos, _, err := listarNamespaces()
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao ler os namespaces", err)
		logger.Erro.Println(mensagem)
		return err
	}

	// Se existir namespaces
	if len(recursos.Items) > 0 {
		// Ler todos os namespaces
		for i := 0; i < len(recursos.Items); i++ {
			recurso := "namespace"
			namespace := recursos.Items[i].Metadata.Name
			nomeRecurso := recursos.Items[i].Metadata.Name

			mensagem := fmt.Sprintf("Executando o backup do namespace [%s]", namespace)
			logger.Info.Println(mensagem)

			// Salvar os dados do namespace
			err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
			if err != nil {
				mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
				logger.Erro.Println(mensagem)
			}

			// Ler todos os recursos válidos
			for i := 0; i < len(recursosValidos.Recursos); i++ {
				recurso := recursosValidos.Recursos[i].Nome

				interf, _, err := listarRecursos(namespace, recurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]: %s", "Erro ao ler os recursos", namespace, recurso, err)
					logger.Erro.Println(mensagem)
					continue
				}

				err = lerDados(namespace, recurso, interf)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]: %s", "Erro ao ler os dados do recurso", namespace, recurso, err)
					logger.Erro.Println(mensagem)
					continue
				}
			}
		}
	}

	return nil
}

//listarNamespaces
func listarNamespaces() (model.Namespaces, int, error) {
	recursos := model.Namespaces{}

	var openshift resource.Openshift
	openshift.Metodo = "ListarNamespaces"

	interf, statusCode, err := resource.Executar(openshift)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao listar os namespaces", err)
		logger.Erro.Println(mensagem)
		return recursos, statusCode, err
	}

	reqBody, err := json.Marshal(interf)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao recuperar o conteúdo dos namespaces", err)
		logger.Erro.Println(mensagem)
		return recursos, http.StatusInternalServerError, err
	}

	err = json.Unmarshal(reqBody, &recursos)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct do namesapce", err)
		logger.Erro.Println(mensagem)
		return recursos, http.StatusInternalServerError, err
	}

	return recursos, statusCode, nil
}

//listarRecursos
func listarRecursos(namespace, recurso string) (interface{}, int, error) {
	var openshift resource.Openshift
	openshift.Metodo = getMetodoListaRecurso(recurso)
	openshift.Namespace = namespace

	interf, statusCode, err := resource.Executar(openshift)
	if err != nil {
		mensagem := fmt.Sprintf("%s [%s]: %s", "Erro ao listar os recursos", recurso, err)
		logger.Erro.Println(mensagem)
		return interf, statusCode, err
	}

	return interf, statusCode, nil
}

//getRecurso
func getRecurso(namespace, recurso, nomeRecurso string) (interface{}, int, error) {
	var openshift resource.Openshift
	openshift.Namespace = namespace
	openshift.NomeRecurso = nomeRecurso
	openshift.Metodo = getMetodoGetRecurso(recurso)

	interf, statusCode, err := resource.Executar(openshift)
	if err != nil {
		mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao executar o backup do recurso", namespace, recurso, nomeRecurso, err)
		logger.Erro.Println(mensagem)
		return interf, statusCode, err
	}

	return interf, statusCode, nil
}

//getMetodoListaRecurso
func getMetodoListaRecurso(recurso string) string {

	switch recurso {
	case "namespace":
		return "ListarNamespaces"
	case "service":
		return "ListarServices"
	case "deploymentconfig":
		return "ListarDeploymentConfigs"
	case "deployment":
		return "ListarDeployments"
	case "secret":
		return "ListarSecrets"
	case "configmap":
		return "ListarConfigMaps"
	case "rolebinding":
		return "ListarRoleBindings"
	case "role":
		return "ListarRoles"
	case "route":
		return "ListarRoutes"
	case "statefulset":
		return "ListarStateFulSets"
	case "buildconfig":
		return "ListarBuildConfigs"
	case "serviceaccount":
		return "ListarServiceAccounts"
	case "replicaset":
		return "ListarReplicaSets"
	case "imagestream":
		return "ListarImageStreams"
	case "resourcequota":
		return "ListarResourceQuotas"
	case "limitrange":
		return "ListarLimitRanges"
	case "cronjob":
		return "ListarCronJobs"
	}
	return ""
}

//getMetodoGetRecurso
func getMetodoGetRecurso(recurso string) string {

	switch recurso {
	case "namespace":
		return "GetNamespace"
	case "service":
		return "GetService"
	case "deploymentconfig":
		return "GetDeploymentConfig"
	case "deployment":
		return "GetDeployment"
	case "secret":
		return "GetSecret"
	case "configmap":
		return "GetConfigMap"
	case "pvc":
		return "GetPVC"
	case "rolebinding":
		return "GetRoleBinding"
	case "role":
		return "GetRoles"
	case "route":
		return "GetRoutes"
	case "statefulset":
		return "GetStateFulSet"
	case "buildconfig":
		return "GetBuildConfig"
	case "serviceaccount":
		return "GetServiceAccount"
	case "replicaset":
		return "GetReplicaSet"
	case "imagestream":
		return "GetImageStream"
	case "resourcequota":
		return "GetResourceQuota"
	case "limitrange":
		return "GetLimitRange"
	case "cronjob":
		return "GetCronJob"
	}
	return ""
}

//lerDados
func lerDados(namespace, recurso string, interf interface{}) error {

	reqBody, err := json.Marshal(interf)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao recuperar o conteúdo do recurso", err)
		logger.Erro.Println(mensagem)
		return err
	}

	switch recurso {
	case "namespace":
		recursos := model.Namespaces{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Name
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "service":
		recursos := model.Services{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "deploymentconfig":
		recursos := model.DeploymentConfigs{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "deployment":
		recursos := model.Deployments{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}
			mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
			logger.Info.Println(mensagem)

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "secret":
		recursos := model.Secrets{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "configmap":
		recursos := model.ConfigMaps{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "rolebinding":
		recursos := model.RoleBindings{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "role":
		recursos := model.Roles{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "route":
		recursos := model.Routes{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "statefulset":
		recursos := model.StateFulSets{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "buildconfig":
		recursos := model.BuildConfigs{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}

		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "serviceaccount":
		recursos := model.ServiceAccounts{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "replicaset":
		recursos := model.ReplicaSets{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "imagestream":
		recursos := model.ImageStreams{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "resourcequota":
		recursos := model.ResourceQuotas{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "limitrange":
		recursos := model.LimitRanges{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	case "cronjob":
		recursos := model.CronJobs{}
		err = json.Unmarshal(reqBody, &recursos)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter JSON para o struct", err)
			logger.Erro.Println(mensagem)
			return err
		}
		if len(recursos.Items) > 0 {
			if variaveis.LogDebug == "S" {
				mensagem := fmt.Sprintf("Executando o backup dos recursos [%s] do namespace [%s]", recurso, namespace)
				logger.Info.Println(mensagem)
			}

			for i := 0; i < len(recursos.Items); i++ {
				namespace := recursos.Items[i].Metadata.Namespace
				nomeRecurso := recursos.Items[i].Metadata.Name

				err := recuperarSalvarDados(namespace, recurso, nomeRecurso)
				if err != nil {
					mensagem := fmt.Sprintf("%s [%s]-[%s]-[%s]: %s", "Erro ao recuperar e salvar o recurso", namespace, recurso, nomeRecurso, err)
					logger.Erro.Println(mensagem)
				}
			}
		}
	}
	return nil
}

//recuperarSalvarDados
func recuperarSalvarDados(namespace, recurso, nomeRecurso string) error {
	if variaveis.LogDebug == "S" {
		mensagem := fmt.Sprintf("Executando o backup do recurso [%s] do namespace [%s]", nomeRecurso, namespace)
		logger.Info.Println(mensagem)
	}

	interf, _, err := getRecurso(namespace, recurso, nomeRecurso)
	if err != nil {
		mensagem := fmt.Sprintf("%s [%s]: %s", "Erro ao recuperar os dados do recurso", nomeRecurso, err)
		logger.Erro.Println(mensagem)
		return err
	}

	err = salvarDados(namespace, recurso, nomeRecurso, interf)
	if err != nil {
		mensagem := fmt.Sprintf("%s [%s]: %s", "Erro ao salvar o recurso  ", nomeRecurso, err)
		logger.Erro.Println(mensagem)
		return err
	}
	return nil
}

//salvarDados
func salvarDados(namespace, recurso, nomeRecurso string, interf interface{}) error {

	dirProjeto := variaveis.DirBase + "/" + namespace
	dirRecurso := dirProjeto + "/" + recurso

	// Criar diretórios
	os.Mkdir(dirProjeto, 0700)
	os.Mkdir(dirRecurso, 0700)

	arquivo := dirRecurso + "/" + nomeRecurso + ".json"

	reqBody, err := json.Marshal(interf)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o recurso", err)
		logger.Erro.Println(mensagem)
		return err
	}

	err = SalvarArquivoJSON(arquivo, string(reqBody))
	if err != nil {
		mensagem := fmt.Sprintf("%s [%s]: %s", "Erro ao salvar o arquivo ", arquivo, err)
		logger.Erro.Println(mensagem)
		return err
	}

	return nil
}
