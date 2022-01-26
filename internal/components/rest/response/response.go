package response

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"infrastructure-telegram/internal/components/utils"
	"net/http"
)

func Write(writer http.ResponseWriter, content interface{}) {
	writer.WriteHeader(200)
	if body, err := json.Marshal(content); err != nil {
		utils.CaptureFatalEvent(sentry.Event{
			Message: "Ошибка сериализации результата в JSON: " + err.Error(),
		})
	} else {
		if _, err := writer.Write(body); err != nil {
			utils.CaptureFatalEvent(sentry.Event{
				Message: "Ошибка записи результата работы: " + err.Error(),
			})
		}
	}
}
