package middleware

import (
	"infrastructure-telegram/config"
	"infrastructure-telegram/internal/components/rest/response"
	"net/http"
)

func Security(cfg config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/healthz" {
			token := request.Header.Get("X-Gitlab-Token")

			if token == "" {
				response.WriteError(writer, http.StatusUnauthorized, "Доступ запрещён: отсутствует обязательный заголовок запроса X-Gitlab-Token")
				return
			}

			if token != cfg.GitLabToken {
				response.WriteError(writer, http.StatusUnauthorized, "Доступ запрещён: неверное значение заголовка X-Gitlab-Token")
				return
			}
		}
		next.ServeHTTP(writer, request)
	})
}
