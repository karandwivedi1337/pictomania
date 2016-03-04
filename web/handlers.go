package web

import (
	"net/http"

	"github.com/flosch/pongo2"
)

var tplIndex = pongo2.Must(pongo2.FromFile("web/templates/index.html"))

//Index handler handles the landing page of the UI
func Index(w http.ResponseWriter, r *http.Request) {
	err := tplIndex.ExecuteWriter(pongo2.Context{"testValue": "Hello World"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
