package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/client/models"
	"github.com/oswee/os/client/utils"
)

func frontApplicationsHandler(r *mux.Router) {
	r.HandleFunc("/applications", frontApplicationsMainHandler).Methods("GET")
}

// Handlers

func frontApplicationsMainHandler(w http.ResponseWriter, r *http.Request) {

	visibility := 1

	services, err := models.ListServices()
	if err != nil {
		fmt.Println("Some error in ervicesHandler")
		return
	}

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-applications.html", struct {
		Title string
		Apps  []models.Application
		Srv   []models.Service
	}{
		Title: "Oswee.com: Applications",
		Apps:  applications,
		Srv:   services,
	})
}
