package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func frontOrganizationsHandler(r *mux.Router) {
	r.HandleFunc("/organizations", organizationsMainHandler).Methods("GET")
}

// Handlers

func organizationsMainHandler(w http.ResponseWriter, r *http.Request) {

	visibility := 1

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	organizations, err := models.ListOrganizations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-organizations.html", struct {
		Title string
		Apps  []models.Application
		Orgs  []models.Organization
	}{
		Title: "Oswee.com: Organizations",
		Apps:  applications,
		Orgs:  organizations,
	})
}
