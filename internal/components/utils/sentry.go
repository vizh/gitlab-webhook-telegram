package utils

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"os"
	"time"
)

func RegisterSentryOptions(options sentry.ClientOptions) {
	if err := sentry.Init(options); err != nil {
		log.Fatalf("register sentry options: %s", err)
	}
}

func CaptureEvent(event sentry.Event) {
	var err error
	// Попытаемся определить имя сервера, на котором работаем. Будет сподручнее в Kubernetes.
	if event.ServerName, err = os.Hostname(); err != nil {
		sentry.CurrentHub().CaptureEvent(&sentry.Event{
			Message: "Ошибка обработки ошибки: не удалось определить имя хоста сервера, на котором работаем",
			Extra: map[string]interface{}{
				"Error": err.Error(),
			},
		})
	}
	// Собственно, отправляем событие в Sentry.
	eventID := sentry.CurrentHub().CaptureEvent(&event)
	// На всякий пожарный, маякнем в консоль. Мало ли, кто-то наблюдает.
	if event.Message == "" {
		event.Message = event.Exception[0].Type + ": " + event.Exception[0].Value
	}
	if _, err = fmt.Fprintf(os.Stderr, "[%s] %s\n", *eventID, event.Message); err != nil {
		sentry.CurrentHub().CaptureEvent(&sentry.Event{
			Message:    "Ошибка обработки ошибки: не удалось вывести сообщение в STDERR, терминал недоступен?",
			ServerName: event.ServerName,
			Extra: map[string]interface{}{
				"Error": err.Error(),
			},
		})
	}
}

func CaptureFatalEvent(event sentry.Event) {
	event.Level = sentry.LevelFatal
	CaptureEvent(event)
	// Сразу выходить нельзя, так как сообщения, в реальности, могут не успеть добраться до серверов Sentry.
	sentry.CurrentHub().Flush(30 * time.Second)
	os.Exit(1)
}
