package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func frontApplicationsHandler(r *mux.Router) {
	r.HandleFunc("/applications", frontApplicationsMainHandler).Methods("GET")
}

// Handlers

func frontApplicationsMainHandler(w http.ResponseWriter, r *http.Request) {

	visibility := 1

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-applications.html", struct {
		Title string
		Apps  []models.Application
	}{
		Title: "Oswee.com: Applications",
		Apps:  applications,
	})
}
