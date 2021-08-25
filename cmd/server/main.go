package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Meta struct {
	originUrl url.URL
}

func (m *Meta) SetOriginUrl(originUrl url.URL) {
	m.originUrl = originUrl
}

func (m *Meta) OriginUrl() url.URL {
	return m.originUrl
}

var meta Meta

func init() {
	meta = Meta{}
}

func handler(w http.ResponseWriter, req *http.Request) {
	pathVars := mux.Vars(req)
	originUrl := meta.originUrl
	proxyUrl, err := url.Parse(originUrl.String() + pathVars["wildcard"])
	if err != nil {
		log.Fatalln(err)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := httpClient.Get(proxyUrl.String())
	if err != nil {
		log.Fatalln(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(dump)
}

func parseArgs(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: go run cmd/server/main.go https://maps.wikimedia.org/")
		os.Exit(1)
	}

	originUrl, err := url.Parse(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	meta.SetOriginUrl(*originUrl)
}

func main() {
	parseArgs(os.Args)

	router := mux.NewRouter()

	router.HandleFunc("/{wildcard:[a-zA-Z0-9/.-]+}", handler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Failed to start HTTP server", err)
	}
}
