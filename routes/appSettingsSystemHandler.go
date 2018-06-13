package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

// System Management Modules
func appSettingsSystemHandler(r *mux.Router) {
	r.HandleFunc("/apps/settings/system", appSettingsSystemMainGetHandler).Methods("GET")
	r.HandleFunc("/apps/settings/system/users", usersGetHandler).Methods("GET")
}

func appSettingsSystemMainGetHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get System Settings modules
	parentModule := 25

	systemSettingsModules, err := models.ListChildApplications(parentModule)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: Cant't list sections"))
		return
	}

	utils.ExecuteTemplate(w, "app-system-settings.html", struct {
		Title                 string
		Mods                  []models.Application
		Shortcuts             []models.Shortcut
		SystemSettingsModules []models.Application
	}{
		Title:                 "System Settings",
		Mods:                  launcherModules,
		Shortcuts:             shortcuts,
		SystemSettingsModules: systemSettingsModules,
	})
}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {
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

	users, err := models.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}
	utils.ExecuteTemplate(w, "mod-system-users.html", struct {
		Title     string
		Mods      []models.Application
		Shortcuts []models.Shortcut
		Users     []models.User
	}{
		Title:     "System Users",
		Mods:      launcherModules,
		Shortcuts: shortcuts,
		Users:     users,
	})
}
