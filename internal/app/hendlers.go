package app

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

type ShortURL struct {
	URL      string
	shortURL string
}

var serverAddress string
var baseURL string

type URLMap struct {
    mx sync.Mutex
    m map[string]ShortURL
}

func NewURLMap() *URLMap {
    return &URLMap{
        m: make(map[string]ShortURL),
    }
}

func (c *URLMap) Load(key string) (ShortURL, bool) {
    c.mx.Lock()
    defer c.mx.Unlock()
    val, ok := c.m[key]
    return val, ok
}

func (c *URLMap) Store(key string, value ShortURL) {
    c.mx.Lock()
    defer c.mx.Unlock()
    c.m[key] = value
}

var storeURLMap = NewURLMap()


func shortingURL(res http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	url := string(body)
	hash := sha256.New()
	hashString := base64.StdEncoding.EncodeToString(hash.Sum([]byte(url)))
	hashShortString := string([]rune(hashString)[len(hashString)-10 : len(hashString)-2])
	storeURLMap.Store(hashShortString, ShortURL{
		URL:      url,
		shortURL: baseURL + "/" + hashShortString,
	})

	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	value, _ := storeURLMap.Load(hashShortString)
	res.Write([]byte(value.shortURL))
}

func getURL(res http.ResponseWriter, req *http.Request) {
	shorturl, isExist := storeURLMap.Load(chi.URLParam(req, "shorturl"))
	if isExist {
		res.Header().Add("location", shorturl.URL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}