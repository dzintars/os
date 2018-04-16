package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func servicesHandler(r *mux.Router) {
	r.HandleFunc("/services", servicesMainHandler).Methods("GET")
}

// Handlers

func servicesMainHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-services.html", nil)
}
