package main

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
)

type ShortURL struct {
    URL       string
    shortURL  string
}

var urlMap = make(map[string]ShortURL)

func shortingURL(res http.ResponseWriter, req *http.Request) {    
    if req.Method == http.MethodPost {
        body, _ := io.ReadAll(req.Body)
        
        url := string(body)

        hash := sha256.New()
        hashString := base64.StdEncoding.EncodeToString(hash.Sum([]byte(url)))
        hashShortString := string([]rune(hashString)[len(hashString)-10:len(hashString)-2])
        urlMap[hashShortString] = ShortURL{
            URL: url,
            shortURL: "http://localhost:8080/"+hashShortString,
        }
        
        res.Header().Set("content-type", "text/plain")
        res.WriteHeader(http.StatusCreated)

        res.Write([]byte(urlMap[hashShortString].shortURL))
    } else { 
        res.WriteHeader(http.StatusBadRequest)
    }
}

func getUrl(res http.ResponseWriter, req *http.Request) { 
    if req.Method == http.MethodGet {
        shorturl := urlMap[req.PathValue("shorturl")].URL
        res.Header().Add("location", shorturl)
        res.WriteHeader(http.StatusTemporaryRedirect)
    } else { 
        res.WriteHeader(http.StatusBadRequest)
    }

}

func main() {

    mux := http.NewServeMux()
    mux.HandleFunc(`/`, shortingURL)
    mux.HandleFunc(`/{shorturl}`, getUrl)

    err := http.ListenAndServe(`:8080`, mux)
    if err != nil {
        panic(err)
    }
} 