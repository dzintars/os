package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

//Customer Relationship Management endpoints
func appCrmHandler(r *mux.Router) {
	r.HandleFunc("/apps/crm", crmGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/dashboard", crmDashboardGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers", crmCustomersGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/projects", crmProjectsGetHandler).Methods("GET")
}

// Handlers

func crmGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.Application{ID: 1, ShortName: "Customer Relationship Management"}
	utils.ExecuteTemplate(w, "mod-crm.html", d)
}

func crmDashboardGetHandler(w http.ResponseWriter, r *http.Request) {

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

	utils.ExecuteTemplate(w, "mod-crm-dashboard.html", struct {
		Title string
		Apps  []models.Application
		Mods  []models.Application
	}{
		Title: "Oswee.com: CRM Dashboard",
		Apps:  applications,
		Mods:  modules,
	})
}

func crmCustomersGetHandler(w http.ResponseWriter, r *http.Request) {

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

	utils.ExecuteTemplate(w, "mod-crm-customers.html", struct {
		Title string
		Apps  []models.Application
		Mods  []models.Application
	}{
		Title: "Oswee.com: CRM Customers",
		Apps:  applications,
		Mods:  modules,
	})
}

func crmProjectsGetHandler(w http.ResponseWriter, r *http.Request) {

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

	utils.ExecuteTemplate(w, "mod-crm-projects.html", struct {
		Title string
		Apps  []models.Application
		Mods  []models.Application
	}{
		Title: "Oswee.com: CRM Projects",
		Apps:  applications,
		Mods:  modules,
	})
}
