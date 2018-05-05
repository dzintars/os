package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func osHandler(r *mux.Router) {

	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/search", searchGetHandler).Methods("GET")
	r.HandleFunc("/search", searchPostHandler).Methods("POST")
	r.HandleFunc("/about", aboutGetHandler).Methods("GET")
}

// Handlers

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	searchCategories := 1

	applications, err := models.ListApplications(searchCategories)
	if err != nil {
		fmt.Println("Some error in osHandler")
		return
	}
	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		Apps  []models.Application
	}{
		Title: "Oswee.com: Shared Resource Planning platform & more",
		Apps:  applications,
	})
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func searchGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.Application{ID: 1, Title: "Search"}
	utils.ExecuteTemplate(w, "mod-search.html", d)
}

func searchPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/mod-search.html", 302)
}

func aboutGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.Application{ID: 1, Title: "About"}
	utils.ExecuteTemplate(w, "mod-about.html", d)
}
