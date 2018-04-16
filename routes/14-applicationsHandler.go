package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func applicationsHandler(r *mux.Router) {
	r.HandleFunc("/apps", applicationsMainHandler).Methods("GET")
}

// Handlers

func applicationsMainHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-apps.html", nil)
}
