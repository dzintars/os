package routes

import (
	"github.com/gorilla/mux"
)

//NewRouter is main routing entry point
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	//There we are calling handlers from package routes
	osHandler(r)  // Root handler
	fileServer(r) // Fileserver to serve static files

	authHandler(r)  // Authorization handlers
	setupHandler(r) // Accounts setup handlers

	servicesHandler(r)     // Front page
	productsHandler(r)     // Front page
	businessHandler(r)     // Front page
	promotionsHandler(r)   // Front page
	applicationsHandler(r) // Front page

	appAppHandler(r) // Application and Module management
	appSysHandler(r) // System level modules.
	appCrmHandler(r) // CRM application and modules
	appSrmHandler(r) // SRM application and modules

	temporaryHandler(r) // Some random stuff testing
	return r
}
