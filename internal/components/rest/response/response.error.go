package response

import (
	"github.com/getsentry/sentry-go"
	"infrastructure-telegram/internal/components/utils"
	"net/http"
)

type errorResponse struct {
	Code    sentry.EventID `json:",omitempty"`
	Message string
}

func WriteError(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	Write(writer, errorResponse{
		Code:    *utils.CaptureEvent(sentry.Event{Message: message}),
		Message: message,
	})
}
