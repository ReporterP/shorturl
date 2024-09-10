package main

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/ReporterP/shorturl/cmd/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ShortURL struct {
    URL       string
    shortURL  string
}

var urlMap = make(map[string]ShortURL)

func shortingURL(res http.ResponseWriter, req *http.Request) {
    body, _ := io.ReadAll(req.Body)
    url := string(body)
    hash := sha256.New()
    hashString := base64.StdEncoding.EncodeToString(hash.Sum([]byte(url)))
    hashShortString := string([]rune(hashString)[len(hashString)-10:len(hashString)-2])

    urlMap[hashShortString] = ShortURL{
        URL: url,
        shortURL: baseaddr + "/" + hashShortString,
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

var baseaddr string
func main() {
    var addrandport string

    appConfig := &config.AppConfig{}

    appConfig.ReadYaml("cmd/config/conf.yaml")
    config.ParseFlags()

    if config.FlagRunAddrAndPort != "" { 
        addrandport = config.FlagRunAddrAndPort 
    } else {
        addrandport = appConfig.AddressAndPort
    }

    if config.FlagRunBaseAddr != "" { 
        baseaddr = config.FlagRunBaseAddr 
    } else {
        baseaddr = appConfig.BaseAddress
    }

    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Post("/", shortingURL)
    r.Get("/{shorturl}", getURL)

    err := http.ListenAndServe(addrandport, r)
    if err != nil {
        panic(err)
    }
} 