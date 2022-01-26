package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"infrastructure-telegram/config"
	"infrastructure-telegram/internal/components/rest/server"
	"infrastructure-telegram/internal/components/utils"
	"log"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	// Запуск веб-сервера...
	log.Printf("Слушаем подключения к %s...\n", cfg.Listen)
	if err := srv.Start(); err != nil {
		utils.CaptureFatalEvent(sentry.Event{
			Message: fmt.Sprintf("Ошибка запуска приложения: %s", err),
		})
	}
}
