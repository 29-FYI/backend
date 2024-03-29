package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/29-FYI/twentynine"
)

type twentyninefyi struct {
	lr     linkring
	logger *log.Logger
}

func (tnfyi *twentyninefyi) Links(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tnfyi.lr.Links())
}

func (tnfyi *twentyninefyi) CreateLink(w http.ResponseWriter, r *http.Request) {
	var link twentynine.Link
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	nlr, ok := tnfyi.lr.LinkLink(link)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	tnfyi.lr = nlr
}

func (tnfyi *twentyninefyi) DeleteLinks(w http.ResponseWriter, r *http.Request) {
	tnfyi.lr = linkring{}
}

func (tnfyi *twentyninefyi) Handler() http.Handler {
	rtr := chi.NewRouter()
	rtr.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	}))
	rtr.Get("/", tnfyi.Links)
	rtr.Post("/", tnfyi.CreateLink)
	rtr.Delete("/", tnfyi.DeleteLinks)
	return rtr
}

func main() {
	// logger
	logger := log.New(os.Stderr, "", log.Lshortfile)

	// handler
	tnfyi := twentyninefyi{
		logger: logger,
	}
	hndlr := tnfyi.Handler()

	// HTTP
	if err := http.ListenAndServe(":http", hndlr); err != nil {
		logger.Fatalf("HTTP server: %s", err)
	}
}
