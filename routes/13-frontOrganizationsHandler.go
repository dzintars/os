package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func frontOrganizationsHandler(r *mux.Router) {
	r.HandleFunc("/organizations", organizationsMainHandler).Methods("GET")
	r.HandleFunc("/organization/{id:[0-9]+}", getOrganization).Methods("GET")
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

func getOrganization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgID := vars["id"]
	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")
	fmt.Println("Org. ID:", orgID, "IP:", r.Header.Get("X-Forwarded-For"), "Time:", ts)

	visibility := 1

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	organizations, err := models.GetOrganization(orgID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-organization-profile.html", struct {
		Title string
		Apps  []models.Application
		Orgs  []models.Organization
	}{
		Title: "Oswee.com: Organization",
		Apps:  applications,
		Orgs:  organizations,
	})
}
