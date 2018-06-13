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

	frontSearchHandler(r)        // Front page
	frontServicesHandler(r)      // Front page
	frontProductsHandler(r)      // Front page
	frontOrganizationsHandler(r) // Front page
	frontPromotionsHandler(r)    // Front page
	frontApplicationsHandler(r)  // Front page list of available applications

	desktopHandler(r)

	appSettingsBusinessHandler(r)
	appSettingsPersonalHandler(r)
	appSettingsSystemHandler(r)
	appRoutePlannerHandler(r)
	appCrmHandler(r) // CRM application and modules
	appSrmHandler(r) // SRM application and modules
	appMarketplaceHandler(r)

	temporaryHandler(r) // Some random stuff testing
	return r
}
