package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/client/models"
	"github.com/oswee/os/client/utils"
)

func frontProductsHandler(r *mux.Router) {
	r.HandleFunc("/products", productsMainHandler).Methods("GET")
}

// Handlers

func productsMainHandler(w http.ResponseWriter, r *http.Request) {

	searchCategories := 1

	services, err := models.ListServices()
	if err != nil {
		fmt.Println("Some error in ervicesHandler")
		return
	}

	applications, err := models.ListApplications(searchCategories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	//ssi := template.HTML(getHTML("http://vid.gov.lv"))

	utils.ExecuteTemplate(w, "mod-products.html", struct {
		Title string
		Apps  []models.Application
		Srv   []models.Service
		//Orgs  template.HTML
	}{
		Title: "Oswee.com: Products",
		Apps:  applications,
		Srv:   services,
		//Orgs:  ssi,
	})
}
