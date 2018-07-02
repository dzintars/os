package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/client/models"
	"github.com/oswee/os/client/utils"
)

func frontSearchHandler(r *mux.Router) {
	r.HandleFunc("/search", searchMainHandler).Methods("GET")
}

// Handlers

func searchMainHandler(w http.ResponseWriter, r *http.Request) {
	visibility := 1

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}
	utils.ExecuteTemplate(w, "mod-search.html", struct {
		Title string
		Apps  []models.Application
	}{
		Title: "Oswee.com: Search",
		Apps:  applications,
	})
}
