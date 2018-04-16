package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func productsHandler(r *mux.Router) {
	r.HandleFunc("/products", productsMainHandler).Methods("GET")
}

// Handlers

func productsMainHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-products.html", nil)
}
