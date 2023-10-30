package ciam

import (
	"github.com/gorilla/mux"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/handler"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
)

type App struct {
	store  handler.StoreHandler
	router *mux.Router
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init(sql client.SqlClient) error {
	storeHandler := handler.NewStoreHandlerConfig()
	if err := storeHandler.Init(&sql); err != nil {
		return err
	}
	a.store = storeHandler

	// Initialize the router and register routes
	a.router = mux.NewRouter()
	a.store.Register("/stores", a.router)

	return nil
}

func (a *App) GetRouter() *mux.Router {
	return a.router
}
