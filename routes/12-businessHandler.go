package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func businessHandler(r *mux.Router) {
	r.HandleFunc("/business", businessMainHandler).Methods("GET")
}

// Handlers

func businessMainHandler(w http.ResponseWriter, r *http.Request) {
	data, err := models.ListOrganizations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}
	utils.ExecuteTemplate(w, "mod-business.html", data)
}
