package routes

import (
	"github.com/gorilla/mux"
)

//NewRouter is main routing entry point
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	//There we are calling handlers from package routes
	mainHandler(r)
	authHandler(r)
	srmHandler(r)
	appCrmHandler(r)
	return r
}
