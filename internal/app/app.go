package app

import (
	"log"
	"net/http"

	"github.com/ReporterP/shorturl/internal/config"
	"github.com/ReporterP/shorturl/internal/logger"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"

	"go.uber.org/zap"
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

	var loglevel string
	
	if loglevel = cfg.EnvLogLevel; cfg.EnvLogLevel == "" {
		loglevel = config.FlagLogLevel
	}

	if errlog := logger.Initialize(loglevel); errlog != nil {
        panic(errlog) 
    }
	
	logger.Log.Info("Running server", zap.String("address", serverAddress))

	r := chi.NewRouter()
	r.Use(logger.RequestLogger)

	r.Post("/", shortingURL)
	r.Get("/{shorturl}", getURL)

	err := http.ListenAndServe(serverAddress, r)
	if err != nil {
		panic(err)
	}
}