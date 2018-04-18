package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func searchHandler(r *mux.Router) {
	r.HandleFunc("/search", searchMainHandler).Methods("GET")
}

// Handlers

func searchMainHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-search.html", nil)
}
