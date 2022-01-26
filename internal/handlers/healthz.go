package handlers

import (
	"infrastructure-telegram/internal/components/rest/response"
	"net/http"
)

// Healthz метод, используемый оркестраторами для определения готовности запускаемого экземпляра приложения к работе.
func Healthz() http.HandlerFunc {
	type payload struct {
		Success bool `json:"success"`
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		response.Write(writer, payload{
			Success: true,
		})
	}
}
