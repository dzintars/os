package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func appAppHandler(r *mux.Router) {
	r.HandleFunc("/apps", appsHandler).Methods("GET")
	r.HandleFunc("/apps/dashboard", appsDashboardHandler).Methods("GET")
}

// Handlers

func appsHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-apps.html", nil)
}

func appsDashboardHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-apps-dashboard.html", nil)
}
