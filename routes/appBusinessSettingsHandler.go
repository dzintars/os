package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func appBusinessSettingsHandler(r *mux.Router) {
	r.HandleFunc("/apps/business-settings", appBusinessSettingsMainGetHandler).Methods("GET")
}

func appBusinessSettingsMainGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// Get Launcher modules
	visibility := 3

	launcherModules, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	// Get Business Settings modules
	parentModule := 26

	businessSettingsModules, err := models.ListChildApplications(parentModule)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: Cant't list sections"))
		return
	}

	shortcuts, err := models.ListShortcuts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: Can't list shortcuts"))
		return
	}

	utils.ExecuteTemplate(w, "app-business-settings.html", struct {
		Title                   string
		Mods                    []models.Application
		BusinessSettingsModules []models.Application
		Shortcuts               []models.Shortcut
	}{
		Title: "Business Account Settings",
		Mods:  launcherModules,
		BusinessSettingsModules: businessSettingsModules,
		Shortcuts:               shortcuts,
	})
}
