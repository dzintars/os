package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

func setupHandler(r *mux.Router) {
	r.HandleFunc("/setup/accounts", sysAccountsGetHandler).Methods("GET")
}

func sysAccountsGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-accounts.html", nil)
}
