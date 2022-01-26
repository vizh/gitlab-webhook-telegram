package middleware

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"infrastructure-telegram/internal/components/utils"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if _, err := fmt.Printf("%s: %s\n", request.Method, request.URL.Path); err != nil {
			utils.CaptureFatalEvent(sentry.Event{
				Message: "Не удалось вывести сообщение в STDOUT, терминал недоступен?",
				Extra: map[string]interface{}{
					"RequestURL": request.Method + ": " + request.URL.String(),
					"RemoteAddr": request.RemoteAddr,
				},
			})
		}
		next.ServeHTTP(writer, request)
	})
}
