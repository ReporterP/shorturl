package main

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/ReporterP/shorturl/cmd/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/caarlos0/env/v6"
)

type ShortURL struct {
    URL       string
    shortURL  string
}

var SERVER_ADDRESS string
var BASE_URL string
var urlMap = make(map[string]ShortURL)

func shortingURL(res http.ResponseWriter, req *http.Request) {
    body, _ := io.ReadAll(req.Body)
    url := string(body)
    hash := sha256.New()
    hashString := base64.StdEncoding.EncodeToString(hash.Sum([]byte(url)))
    hashShortString := string([]rune(hashString)[len(hashString)-10:len(hashString)-2])

    urlMap[hashShortString] = ShortURL{
        URL: url,
        shortURL: BASE_URL + "/" + hashShortString,
    }
    
    res.Header().Set("content-type", "text/plain")
    res.WriteHeader(http.StatusCreated)
    res.Write([]byte(urlMap[hashShortString].shortURL))
}

func getURL(res http.ResponseWriter, req *http.Request) { 
    shorturl := urlMap[chi.URLParam(req, "shorturl")].URL
    res.Header().Add("location", shorturl)
    res.WriteHeader(http.StatusTemporaryRedirect)
}


func main() {
    var cfg config.Config
    errEnv := env.Parse(&cfg)
    if errEnv != nil {
        log.Fatal(errEnv)
    }
    config.ParseFlags()

    if SERVER_ADDRESS = cfg.ServerAddress; cfg.ServerAddress == "" {
        SERVER_ADDRESS = config.FlagRunAddrAndPort
    }

    if BASE_URL = cfg.BaseURL; cfg.BaseURL == "" { 
        BASE_URL = config.FlagRunBaseAddr
    }

    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Post("/", shortingURL)
    r.Get("/{shorturl}", getURL)

    err := http.ListenAndServe(SERVER_ADDRESS, r)
    if err != nil {
        panic(err)
    }
} 