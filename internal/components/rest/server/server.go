package server

import (
	"infrastructure-telegram/config"
	"infrastructure-telegram/internal/components/rest/middleware"
	"infrastructure-telegram/internal/handlers"
	"net/http"
)

type Server interface {
	Start() error
}

type server struct {
	cfg config.Config
}

func New(config config.Config) Server {
	return &server{
		cfg: config,
	}
}

func (s server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.Healthz())
	mux.HandleFunc("/", handlers.GitLab())
	return http.ListenAndServe(s.cfg.Listen, middleware.Logger(middleware.Security(s.cfg, mux)))
}
