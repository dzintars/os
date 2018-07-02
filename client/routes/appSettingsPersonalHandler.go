package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/client/models"
	"github.com/oswee/os/client/utils"
)

func appSettingsPersonalHandler(r *mux.Router) {
	r.HandleFunc("/settings/personal", appSettingsPersonalMainGetHandler).Methods("GET")
}

func appSettingsPersonalMainGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	// Get Launcher modules
	visibility := 3

	launcherModules, err := models.ListApplications(visibility)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}

	// Get Personal Settings sub-modules
	parentModule := 27

	personalSettingsModules, err := models.ListChildApplications(parentModule)
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

	utils.ExecuteTemplate(w, "app-personal-settings.html", struct {
		Title                   string
		Mods                    []models.Application
		PersonalSettingsModules []models.Application
		Shortcuts               []models.Shortcut
	}{
		Title: "Personal Account Settings",
		Mods:  launcherModules,
		PersonalSettingsModules: personalSettingsModules,
		Shortcuts:               shortcuts,
	})
}
