package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/client/models"
	"github.com/oswee/os/client/utils"
)

// System Management Modules
func appMarketplaceHandler(r *mux.Router) {
	r.HandleFunc("/marketplace", appMarketplaceMainGetHandler).Methods("GET")
}

func appMarketplaceMainGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// Get Launcher modules
	visibility := 3

	launcherModules, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: Can't list shortcuts"))
		return
	}

	utils.ExecuteTemplate(w, "app-marketplace.html", struct {
		Title     string
		Mods      []models.Application
		Shortcuts []models.Shortcut
	}{
		Title:     "Oswee Marketplace",
		Mods:      launcherModules,
		Shortcuts: shortcuts,
	})
}
