package main

import (
	"log"
	"net/http"
	"os"

	"github.com/asonnenschein/technical-assessment-golang/internal/handlers"
	"github.com/asonnenschein/technical-assessment-golang/internal/models"
	"github.com/asonnenschein/technical-assessment-golang/internal/utils"
	"github.com/gorilla/mux"
)

var meta models.Meta

func main() {
	meta = utils.ParseArgs(os.Args)

	router := mux.NewRouter()
	router.HandleFunc("/{wildcard:[a-zA-Z0-9/.-]+}", handlers.WildCard(meta))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Failed to start HTTP server", err)
	}
}
