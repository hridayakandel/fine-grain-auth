package router

import (
	"github.com/gorilla/mux"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/handler"
)

// SetupRouter sets up the routes and returns the router instance.
func SetupRouter(storeHandler *handler.StoreHandler) *mux.Router {
	r := mux.NewRouter()

	// Setting up the stores endpoints
	storeSubRouter := r.PathPrefix("/stores").Subrouter()
	storeSubRouter.HandleFunc("", storeHandler.CreateStore).Methods("POST")

	// You can add more routes and sub-routers as needed

	return r
}
