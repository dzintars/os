package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

// Customer Relationship Management endpoints
func appHandler(r *mux.Router) {
	r.HandleFunc("/apps", appGetHandler).Methods("GET")
}

// Handlers

func appGetHandler(w http.ResponseWriter, r *http.Request) {

	apps, err := models.ListChildApplications(7)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	visibleModules := 3

	modules, err := models.ListApplications(visibleModules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	utils.ExecuteTemplate(w, "mod-apps.html", struct {
		Title     string
		Mods      []models.Application
		Apps      []models.Application
		Shortcuts []models.Shortcut
	}{
		Title:     "Oswee Applications",
		Mods:      modules,
		Apps:      apps,
		Shortcuts: shortcuts,
	})
}
