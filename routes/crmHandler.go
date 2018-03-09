package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

//Customer Relationship Management
func appCrmHandler(r *mux.Router) {
	r.HandleFunc("/apps/crm", crmHandler).Methods("GET")
	r.HandleFunc("/apps/crm/dashboard", dashboardHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers", customersHandler).Methods("GET")
	r.HandleFunc("/apps/crm/projects", projectsHandler).Methods("GET")
}

func crmHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management"}
	utils.ExecuteTemplate(w, "mod-crm.html", d)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management", ModuleTitle: "Dashboard"}
	utils.ExecuteTemplate(w, "mod-crm-dashboard.html", d)
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management", ModuleTitle: "Customers"}
	utils.ExecuteTemplate(w, "mod-crm-customers.html", d)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "CRM", Title: "Customer Relationship Management", ModuleTitle: "Projects"}
	utils.ExecuteTemplate(w, "mod-crm-projects.html", d)
}
