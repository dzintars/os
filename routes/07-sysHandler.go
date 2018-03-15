package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

// System Management Modules
func appSysHandler(r *mux.Router) {
	r.HandleFunc("/apps/sys", sysGetHandler).Methods("GET")
	r.HandleFunc("/apps/sys/users", usersGetHandler).Methods("GET")
}

func sysGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-sys.html", nil)
}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {
	data, err := models.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}
	utils.ExecuteTemplate(w, "mod-sys-users.html", data)
}
