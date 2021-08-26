package handlers

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/asonnenschein/technical-assessment-golang/internal/models"
	"github.com/gorilla/mux"
)

func WildCard(meta models.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		pathVars := mux.Vars(req)
		originUrl := meta.GetOriginUrl()

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

		contentType := resp.Header.Get("Content-Type")

		w.Header().Set("Content-Type", contentType)

		img, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		grayImage := &models.GrayScale{Image: img}

		if contentType == "image/png" {
			png.Encode(w, grayImage)
		} else if contentType == "image/jpeg" {
			jpeg.Encode(w, grayImage, nil)
		}
	}
}
