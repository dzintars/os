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
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management"}
	utils.ExecuteTemplate(w, "mod-crm.html", d)
}

func crmDashboardGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management", ModuleTitle: "Dashboard"}
	utils.ExecuteTemplate(w, "mod-crm-dashboard.html", d)
}

func crmCustomersGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management", ModuleTitle: "Customers"}
	utils.ExecuteTemplate(w, "mod-crm-customers.html", d)
}

func crmProjectsGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management", ModuleTitle: "Projects"}
	utils.ExecuteTemplate(w, "mod-crm-projects.html", d)
}