package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/asonnenschein/technical-assessment-golang/internal/models"
)

func ParseArgs(args []string) models.Meta {
	if len(args) != 2 {
		fmt.Println("Usage: go run cmd/server/main.go https://maps.wikimedia.org/")
		os.Exit(1)
	}

	originUrl, err := url.Parse(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	meta := models.Meta{}

	meta.SetOriginUrl(*originUrl)

	return meta
}
