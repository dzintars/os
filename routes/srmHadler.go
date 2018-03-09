package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func srmHandler(r *mux.Router) {
	r.HandleFunc("/srm", servicesHandler).Methods("GET")
	r.HandleFunc("/srm/suppliers", suppliersHandler).Methods("GET")
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "services.html", nil)
}

func suppliersHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "services.html", nil)
}
