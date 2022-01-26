package response

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"infrastructure-telegram/internal/components/utils"
	"net/http"
)

func Write(responseWriter http.ResponseWriter, content interface{}) {
	if body, err := json.Marshal(content); err != nil {
		utils.CaptureFatalEvent(sentry.Event{
			Message: "Ошибка сериализации результата в JSON: " + err.Error(),
		})
	} else {
		if _, err := responseWriter.Write(body); err != nil {
			utils.CaptureFatalEvent(sentry.Event{
				Message: "Ошибка записи результата работы: " + err.Error(),
			})
		}
	}
}
