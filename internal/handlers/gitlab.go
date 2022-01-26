package handlers

import (
	"infrastructure-telegram/internal/components/rest/response"
	"io"
	"log"
	"net/http"
)

// GitLab метод, принимающий данные со стороны GitLab.
func GitLab() http.HandlerFunc {
	type payload struct {
		Success bool
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonBytes, err := io.ReadAll(request.Body)
		if err != nil {
			response.WriteError(writer, http.StatusBadRequest, "Ошибка разбора данных запроса: %s", err)
			return
		}
		log.Println(string(jsonBytes))
		response.Write(writer, payload{
			Success: true,
		})
	}
}
