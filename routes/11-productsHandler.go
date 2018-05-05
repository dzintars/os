package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func productsHandler(r *mux.Router) {
	r.HandleFunc("/products", productsMainHandler).Methods("GET")
}

// Handlers

func productsMainHandler(w http.ResponseWriter, r *http.Request) {

	searchCategories := 1

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
		//Orgs  template.HTML
	}{
		Title: "Oswee.com: Products",
		Apps:  applications,
		//Orgs:  ssi,
	})
}
