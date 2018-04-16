package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func businessHandler(r *mux.Router) {
	r.HandleFunc("/business", businessMainHandler).Methods("GET")
}

// Handlers

func businessMainHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-business.html", nil)
}
