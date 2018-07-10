package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func indexHandler(r *mux.Router) {

	r.HandleFunc("/favicon.ico", faviconHandler).Methods("GET")
	r.HandleFunc("/non-supported-browser", unsupportedBrowserHandler).Methods("GET")
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/search", searchGetHandler).Methods("GET")
	r.HandleFunc("/search", searchPostHandler).Methods("POST")
	r.HandleFunc("/about", aboutGetHandler).Methods("GET")
	r.HandleFunc("/about", aboutPostHandler).Methods("POST")
}

// Handlers

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "/static/img/favicon.png", nil)
}

func unsupportedBrowserHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-unsupported-browser.html", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(r.Header.Get("X-Forwarded-For"), t)

	searchCategories := 1

	applications, err := models.ListApplications(searchCategories)
	if err != nil {
		fmt.Println("Some error in osHandler")
		return
	}
	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		Apps  []models.Application
	}{
		Title: "Oswee.com: Shared Resource Planning platform & more",
		Apps:  applications,
	})
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func searchGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-search.html", struct {
		Title string
	}{
		Title: "Oswee.com: Search",
	})
}

func searchPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/mod-search.html", 302)
}

func aboutGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "mod-about.html", struct {
		Title string
	}{
		Title: "Oswee.com: About",
	})
}

func aboutPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	//name := r.PostForm.Get("account-name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	// Database function goes there
	lastID, err := models.MessageCreate(email, message)

	fmt.Println("New message received: ", lastID)
	http.Redirect(w, r, "/about", 302)
}
