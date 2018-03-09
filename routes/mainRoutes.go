package routes

import (
	"net/http"

	"io/ioutil"
	"log"

	"github.com/gorilla/mux"
	"github.com/oswee/os/models"
	"github.com/oswee/os/utils"
)

func mainHandler(r *mux.Router) {
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/search", searchGetHandler).Methods("GET")
	r.HandleFunc("/search", searchPostHandler).Methods("POST")
	r.HandleFunc("/about", aboutGetHandler).Methods("GET")
	r.HandleFunc("/test", testGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "OS", Title: "Collaborative Resource Planning"}
	utils.ExecuteTemplate(w, "index.html", d)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func searchGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "OS", Title: "Search"}
	utils.ExecuteTemplate(w, "mod-search.html", d)
}

func searchPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/mod-search.html", 302)
}
func aboutGetHandler(w http.ResponseWriter, r *http.Request) {
	d := models.App{Abr: "OS", Title: "About"}
	utils.ExecuteTemplate(w, "mod-about.html", d)
}

func testGetHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	data, err := models.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: 001, Internal Server Error"))
		return
	}
	//d := getHTML()
	// utils.ExecuteTemplate(w, "mod-test.html", template.HTML(d))
	utils.ExecuteTemplate(w, "mod-test.html", data)
}

func getHTML() string {
	url := "http://google.com"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}
