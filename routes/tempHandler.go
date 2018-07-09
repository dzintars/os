package routes

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/os/utils"
)

// Handler to test stuff.
func temporaryHandler(r *mux.Router) {
	r.HandleFunc("/temp/vika", vikaHandler).Methods("GET")
}

// Handlers

func vikaHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "temp-vika.html", nil)
}

func getHTML(url string) string {
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
