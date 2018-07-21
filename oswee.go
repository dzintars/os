package main

import (
	"net/http"

	"github.com/oswee/os/routes"
	"github.com/oswee/os/utils"
)

func main() {
	utils.LoadTemplates("templates/*.html")
	//Load routes package
	r := routes.NewRouter()
	//Pass routes package to Handler
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}
