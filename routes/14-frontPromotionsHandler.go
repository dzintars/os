package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func frontPromotionsHandler(r *mux.Router) {
	r.HandleFunc("/promotions", promotionsMainHandler).Methods("GET")
}

// Handlers

func promotionsMainHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-promotions.html", nil)
}
