package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func appSrmHandler(r *mux.Router) {
	r.HandleFunc("/apps/srm", srmGetHandler).Methods("GET")
	r.HandleFunc("/apps/srm/services", srmServicesGetHandler).Methods("GET")
	r.HandleFunc("/apps/srm/suppliers", suppliersGetHandler).Methods("GET")
}

func srmGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-srm.html", nil)
}
func srmServicesGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-srm-services.html", nil)
}

func suppliersGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-srm-suppliers.html", nil)
}
