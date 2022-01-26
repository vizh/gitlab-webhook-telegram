package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/getsentry/sentry-go"
	"infrastructure-telegram/internal/components/utils"
	"log"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Release     string `env:"RELEASE" envDefault:"edge"`
	Listen      string `env:"LISTEN" envDefault:"127.0.0.1:8080"` // В целях безопасности, приложение не слушает на внешних интерфейсах, по-умолчанию.
	SentryDSN   string `env:"SENTRY_DSN,required"`
}

func Load() (config Config) {
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Ошибка конфигурации приложения: %s", err)
	}
	utils.RegisterSentryOptions(sentry.ClientOptions{
		Dsn:         config.SentryDSN,
		Environment: config.Environment,
		Release:     config.Release,
	})
	return
}
