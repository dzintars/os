package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func frontServicesHandler(r *mux.Router) {
	r.HandleFunc("/services", servicesMainHandler).Methods("GET")
}

// Handlers

func servicesMainHandler(w http.ResponseWriter, r *http.Request) {

	searchCategories := 1

	services, err := models.ListServices()
	if err != nil {
		fmt.Println("Some error in ervicesHandler")
		return
	}
	applications, err := models.ListApplications(searchCategories)
	if err != nil {
		fmt.Println("Some error in osHandler")
		return
	}
	utils.ExecuteTemplate(w, "mod-services.html", struct {
		Title string
		Apps  []models.Application
		Srv   []models.Service
	}{
		Title: "Oswee.com: Services",
		Apps:  applications,
		Srv:   services,
	})
}
