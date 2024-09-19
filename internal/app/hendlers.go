package app

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ShortURL struct {
	URL      string
	shortURL string
}

var serverAddress string
var baseURL string
var urlMap = make(map[string]ShortURL)

func shortingURL(res http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	url := string(body)
	hash := sha256.New()
	hashString := base64.StdEncoding.EncodeToString(hash.Sum([]byte(url)))
	hashShortString := string([]rune(hashString)[len(hashString)-10 : len(hashString)-2])

	urlMap[hashShortString] = ShortURL{
		URL:      url,
		shortURL: baseURL + "/" + hashShortString,
	}

	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(urlMap[hashShortString].shortURL))
}

func getURL(res http.ResponseWriter, req *http.Request) {
	shorturl, isExist := urlMap[chi.URLParam(req, "shorturl")]
	if isExist {
		res.Header().Add("location", shorturl.URL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}