package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"url-shortener/database"
	"url-shortener/server/response"
	"url-shortener/validation"
)

type CreateLinkRequestBody struct {
	Link string `json:"link"`
}

type ErrorPageData struct {
	IndexLink string
}

func createLinkHandler(w http.ResponseWriter, req *http.Request) {
	var payload CreateLinkRequestBody
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		response.InvalidJSON(w)
		log.Println("Failed to parse JSOn")
		return
	}

	var shortened int

	shortened, err = database.CreateShortenedUrlQuery(payload.Link)

	if err != nil {
		response.InternalError(w)
		log.Println("Error while creating row")
		return
	}

	response.Created(w, fmt.Sprintf("%s/s/%d", os.Getenv("HOST_NAME"), shortened))

}

func getLinkHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	id := strings.TrimPrefix(path, "/s/")

	if id == "" {
		response.Error(w, http.StatusNotFound, "No link provided")
		log.Println("No short link provided")
		return
	}

	url, err := database.GetOriginalUrlQuery(id)
	if url == "" || err != nil || !validation.IsValidURL(url) {
		data := ErrorPageData{IndexLink: os.Getenv("HOST_NAME")}
		response.HTMLResponse(w, data, "templates/error.html")
		log.Printf("Original url for %s not found", id)
		return
	}

	http.Redirect(w, req, url, http.StatusFound)
}

func getIndexPageHandler(w http.ResponseWriter, req *http.Request) {
	response.HTMLResponse(w, "", "templates/index.html")
}

func StartServer() {
	log.Println("Starting http server")
	defer log.Println("http server stopped")

	http.HandleFunc("POST /link/shorten", createLinkHandler)
	http.HandleFunc("GET /s/", getLinkHandler)
	http.HandleFunc("GET /", getIndexPageHandler)

	http.ListenAndServe(":8080", nil)
}
