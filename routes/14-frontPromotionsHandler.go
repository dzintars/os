package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func frontPromotionsHandler(r *mux.Router) {
	r.HandleFunc("/promotions", frontPromotionsMainHandler).Methods("GET")
}

// Handlers

func frontPromotionsMainHandler(w http.ResponseWriter, r *http.Request) {

	visibility := 1

	applications, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-promotions.html", struct {
		Title string
		Apps  []models.Application
	}{
		Title: "Oswee.com: Promotions",
		Apps:  applications,
	})
}
