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

// WildCard is an HTTP service handler that accepts wildcard parameters, proxies requests to an
// origin domain supplied by user at server startup, and transforms color PNG/JPEG images to
// grayscale if an image is returned in the proxy response.  This method uses a closure to pass
// global settings (i.e. origin domain) to the request handler via Meta struct.  It is optimistic
// in nature, in that it does not currently implement any graceful failure mechanisms.
func WildCard(meta models.Meta) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		// Pull wildcard path parameters from router via Gorilla/Mux library
		pathVars := mux.Vars(req)
		// Get user supplied origin URL from server startup
		originUrl := meta.GetOriginUrl()

		// Create the full proxy URL (origin domain + wildcard path) and sanity check it for integrity
		proxyUrl, err := url.Parse(originUrl.String() + pathVars["wildcard"])
		if err != nil {
			log.Fatalln(err)
		}

		// Override default http client w/ explicit 5 second timeout
		httpClient := &http.Client{
			Timeout: time.Second * 5,
		}

		// Perform HTTP GET request for proxy URL
		resp, err := httpClient.Get(proxyUrl.String())
		if err != nil {
			log.Fatalln(err)
		}

		// Forward proxied request Content-Type header to http.ResponsWriter
		contentType := resp.Header.Get("Content-Type")
		w.Header().Set("Content-Type", contentType)

		// Decode proxied response body into image format from buffer
		img, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// Convert decoded color image to grayscale
		grayImage := &models.GrayScale{Image: img}

		// Encode grayscale image according to requested Content-Type and add to http.ResponseWriter
		if contentType == "image/png" {
			png.Encode(w, grayImage)
		} else if contentType == "image/jpeg" {
			jpeg.Encode(w, grayImage, nil)
		}
	}
}
