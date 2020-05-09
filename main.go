package main

import (
	"encoding/json"
	"net/http"

	"github.com/29-FYI/twentynine"
)

var lr linkring

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var link twentynine.Link
		if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		nlr, ok := lr.LinkLink(link)
		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		lr = nlr

		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(lr.Links())
	case http.MethodDelete:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6969", nil)
}
