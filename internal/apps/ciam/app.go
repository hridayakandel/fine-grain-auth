package ciam

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/handler"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
)

type App struct {
	storeHandler *handler.StoreHandler
}

func (a *App) Init(ctx context.Context, dbClient *client.SqlClient) error {
	a.storeHandler = handler.NewStoreHandler(dbClient)
	return nil
}

func (a *App) Register(r *mux.Router) {
	storeRoutes := r.PathPrefix("/stores").Subrouter()
	storeRoutes.HandleFunc("", a.storeHandler.CreateStore).Methods("POST")
}
