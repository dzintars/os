package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func appRoutePlannerHandler(r *mux.Router) {
	r.HandleFunc("/apps/router", appRoutePlannerMainGetHandler).Methods("GET")
}

func appRoutePlannerMainGetHandler(w http.ResponseWriter, r *http.Request) {
	visibility := 1

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
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
	utils.ExecuteTemplate(w, "app-route-planner.html", struct {
		Title string
		Apps  []models.Application
		Orgs  []models.Organization
		Mods  []models.Application
	}{
		Title: "Oswee.com: Route Planner",
		Apps:  applications,
		Orgs:  organizations,
		Mods:  modules,
	})
}
