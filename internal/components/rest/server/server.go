package server

import (
	pkgConfig "infrastructure-telegram/config"
	"infrastructure-telegram/internal/handlers"
	"net/http"
)

type Server interface {
	Start() error
}

type server struct {
	listen string
}

func New(config pkgConfig.Config) Server {
	return &server{
		listen: config.Listen,
	}
}

func (s server) Start() error {
	http.HandleFunc("/healthz", handlers.Healthz())
	return http.ListenAndServe(s.listen, nil)
}
