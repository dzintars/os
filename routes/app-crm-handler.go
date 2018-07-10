package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

// Customer Relationship Management endpoints
func appCrmHandler(r *mux.Router) {
	r.HandleFunc("/apps/crm", crmGetHandler).Methods("GET")                        // Show CRM App modules, extensions and other
	r.HandleFunc("/apps/crm/dashboard", crmDashboardGetHandler).Methods("GET")     // Show main (default) Dashboard of CRM
	r.HandleFunc("/apps/crm/customers", crmCustomersGetHandler).Methods("GET")     // Show list of all customers (Partner Accounts) and management tools
	r.HandleFunc("/apps/crm/customers/new", crmCustomersNewHandler).Methods("GET") // Show form to create new customer (Partner Account) entry
	r.HandleFunc("/apps/crm/customers/new", crmCustomerCreateHandler).Methods("POST")
	r.HandleFunc("/apps/crm/customers/{customerID:[0-9]+}/update", crmCustomerUpdateGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers/{customerID:[0-9]+}/update", crmCustomerUpdatePostHandler).Methods("POST")
	r.HandleFunc("/apps/crm/customers/{id:[0-9]+}", crmCustomerGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers/{customerID:[0-9]+}/dashboard", crmCustomerDashboardGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers/{customerID:[0-9]+}/profile", crmCustomerProfileHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers/{customerID:[0-9]+}/projects", crmCustomerProjectsGetHandler).Methods("GET")
	r.HandleFunc("/apps/crm/customers/delete/{customerID:[0-9]+}", crmCustomerDeleteHandler).Methods("POST")
	r.HandleFunc("/apps/crm/projects", crmProjectsGetHandler).Methods("GET")
}

// Handlers

func crmGetHandler(w http.ResponseWriter, r *http.Request) {

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm.html", struct {
		Title     string
		Mods      []models.Application
		Shortcuts []models.Shortcut
	}{
		Title:     "Oswee Applications",
		Mods:      modules,
		Shortcuts: shortcuts,
	})
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

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	customers, err := models.ListCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customers.html", struct {
		Title     string
		Mods      []models.Application
		Customers []models.Customer
		Shortcuts []models.Shortcut
	}{
		Title:     "Oswee.com: CRM Customers",
		Mods:      modules,
		Customers: customers,
		Shortcuts: shortcuts,
	})
}

func crmProjectsGetHandler(w http.ResponseWriter, r *http.Request) {

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	projects, err := models.ListProjects()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-projects.html", struct {
		Title     string
		Mods      []models.Application
		Projects  []models.Project
		Shortcuts []models.Shortcut
	}{
		Title:     "Oswee.com: CRM Projects",
		Mods:      modules,
		Projects:  projects,
		Shortcuts: shortcuts,
	})
}

func crmCustomersNewHandler(w http.ResponseWriter, r *http.Request) {

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customers-new.html", struct {
		Title     string
		Customer  models.Customer
		Mods      []models.Application
		Projects  []models.Project
		Shortcuts []models.Shortcut
	}{
		Title:     "CRM Customers New",
		Mods:      modules,
		Shortcuts: shortcuts,
	})
}

func crmCustomerCreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	//name := r.PostForm.Get("account-name")
	accName := r.FormValue("account-name")
	// Database function goes there
	lastID, err := models.CustomerCreate(accName)

	fmt.Println("New customer created: ", lastID)
	http.Redirect(w, r, "/apps/crm/customers", 302)
}

func crmCustomerUpdateGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	customer := models.GetCustomer(customerID)

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customers-update.html", struct {
		Title     string
		Customer  models.Customer
		Mods      []models.Application
		Shortcuts []models.Shortcut
	}{
		Title:     "CRM Customers New",
		Customer:  customer,
		Mods:      modules,
		Shortcuts: shortcuts,
	})
}

func crmCustomerUpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	accName := r.PostForm.Get("account-name")
	//accName := r.FormValue("account-name")
	// Database function goes there
	err = models.CustomerUpdate(accName, customerID)
	if err != nil {
		log.Println(err.Error)
	}

	fmt.Println("New customer updated: ")
	http.Redirect(w, r, "/apps/crm/customers", 302)
}

func crmCustomerDashboardGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	application := models.GetApplication(43) // Get CRM Customer Dashboard App

	customer := models.GetCustomer(customerID)

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	tabs, err := models.ListChildApplications(15)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customer.html", struct {
		Title     string
		App       models.Application
		Customer  models.Customer
		Mods      []models.Application
		Tabs      []models.Application
		Shortcuts []models.Shortcut
	}{
		Title:     "CRM Customer",
		App:       application,
		Customer:  customer,
		Mods:      modules,
		Tabs:      tabs,
		Shortcuts: shortcuts,
	})
}

func crmCustomerGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	application := models.GetApplication(15) // Get CRM App

	customer := models.GetCustomer(customerID)

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	tabs, err := models.ListChildApplications(15)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customer.html", struct {
		Title     string
		App       models.Application
		Customer  models.Customer
		Mods      []models.Application
		Tabs      []models.Application
		Shortcuts []models.Shortcut
	}{
		Title:     "CRM Customer",
		App:       application,
		Customer:  customer,
		Mods:      modules,
		Tabs:      tabs,
		Shortcuts: shortcuts,
	})
}

func crmCustomerProjectsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	application := models.GetApplication(46) // Get CRM Customer Project

	customer := models.GetCustomer(customerID)

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	tabs, err := models.ListChildApplications(15)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	projects, err := models.ListCustomerProjects(customerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customer.html", struct {
		Title     string
		App       models.Application
		Customer  models.Customer
		Mods      []models.Application
		Tabs      []models.Application
		Projects  []models.Project
		Shortcuts []models.Shortcut
	}{
		Title:     "Oswee.com: CRM Projects",
		App:       application,
		Customer:  customer,
		Mods:      modules,
		Tabs:      tabs,
		Projects:  projects,
		Shortcuts: shortcuts,
	})
}

func crmCustomerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	err := models.CustomerDelete(customerID)
	if err != nil {
		log.Println(err.Error)
	}
	http.Redirect(w, r, "/apps/crm/customers", 302)
}

func crmCustomerProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]

	customer := models.GetCustomer(customerID)

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	tabs, err := models.ListChildApplications(14)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-crm-customer-profile.html", struct {
		Title     string
		Section   string
		Customer  models.Customer
		Mods      []models.Application
		Shortcuts []models.Shortcut
		Tabs      []models.Application
	}{
		Title:     "CRM Customer Profile",
		Section:   "customer-profile",
		Customer:  customer,
		Mods:      modules,
		Shortcuts: shortcuts,
		Tabs:      tabs,
	})
}
