package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func fileServer(r *mux.Router) {
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
