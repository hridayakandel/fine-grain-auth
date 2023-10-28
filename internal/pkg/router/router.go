// pkg/router/router.go

package router

import (
	"github.com/gorilla/mux"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/handler"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Setting up the stores endpoints
	s := r.PathPrefix("/stores").Subrouter()
	s.HandleFunc("", handler.CreateStoreHandler).Methods("POST")

	// ... You can expand this with more routes and middlewares

	return r
}
