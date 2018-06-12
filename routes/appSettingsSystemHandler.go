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
	r.HandleFunc("/apps/sys/users", usersGetHandler).Methods("GET")
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

	// Get System Settings modules
	parentModule := 25

	systemSettingsModules, err := models.ListChildApplications(parentModule)
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

	utils.ExecuteTemplate(w, "app-system-settings.html", struct {
		Title                 string
		Mods                  []models.Application
		SystemSettingsModules []models.Application
		Shortcuts             []models.Shortcut
	}{
		Title: "System Settings",
		Mods:  launcherModules,
		SystemSettingsModules: systemSettingsModules,
		Shortcuts:             shortcuts,
	})
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
