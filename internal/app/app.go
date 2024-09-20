package app

import (
	"log"
	"net/http"

	"github.com/ReporterP/shorturl/internal/config"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	var cfg config.Config
	errEnv := env.Parse(&cfg)
	if errEnv != nil {
		log.Fatal(errEnv)
	}
	config.ParseFlags()

	if serverAddress = cfg.ServerAddress; cfg.ServerAddress == "" {
		serverAddress = config.FlagRunAddrAndPort
	}

	if baseURL = cfg.BaseURL; cfg.BaseURL == "" {
		baseURL = config.FlagRunBaseAddr
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/", shortingURL)
	r.Get("/{shorturl}", getURL)

	err := http.ListenAndServe(serverAddress, r)
	if err != nil {
		panic(err)
	}
}
