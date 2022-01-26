package handlers

import (
	"encoding/json"
	"infrastructure-telegram/internal/components/rest/response"
	"io"
	"log"
	"net/http"
)

// GitLab метод, принимающий данные со стороны GitLab.
func GitLab() http.HandlerFunc {
	type kind string

	const (
		kindDeployment kind = "deployment"
		kindBuild           = "build"
	)

	type user struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type project struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		Path      string `json:"path_with_namespace"`
	}

	type webhook struct {
		Kind kind `json:"object_kind"`
	}

	type webhookDeployment struct {
		webhook
		Status  string  `json:"status"`
		User    user    `json:"user"`
		Project project `json:"project"`
		// Информация о коммите.
		CommitSHA   string `json:"short_sha"`
		CommitTitle string `json:"commit_title"`
	}

	type payload struct {
		Success bool
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		jsonBytes, err := io.ReadAll(request.Body)
		if err != nil {
			response.WriteError(writer, http.StatusBadRequest, "Ошибка разбора данных запроса: %s", err)
			return
		}
		// Определяем тип полученных данных.
		webhook := webhook{}
		if err = json.Unmarshal(jsonBytes, &webhook); err != nil {
			response.WriteError(writer, http.StatusBadRequest, "Ошибка определения типа полученных данных: %s", err)
			return
		}
		switch webhook.Kind {
		case kindDeployment:
			data := webhookDeployment{}
			if err = json.Unmarshal(jsonBytes, &data); err != nil {
				response.WriteError(writer, http.StatusBadRequest, "Ошибка разбора данных %s: %s", webhook.Kind, err)
				return
			}
			log.Printf("%+v\n", data)
		default:
			log.Println(string(jsonBytes))
		}
		response.Write(writer, payload{
			Success: true,
		})
	}
}
