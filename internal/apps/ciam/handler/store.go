package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/service"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
	"log"
	"net/http"
)

type storeHandlerConfig struct {
	service service.StoreService
}

type StoreHandler interface {
	Init(sql *client.SqlClient) error
	Register(pathPrefix string, router *mux.Router)
}

func NewStoreHandlerConfig() StoreHandler {
	return &storeHandlerConfig{}
}

func (s *storeHandlerConfig) Init(sql *client.SqlClient) error {
	log.Println("Starting CIAM store service initialization")

	// Initializing the service layer inside the handler's Init method.
	s.service = service.NewStoreService(repository.NewStoreRepo(sql))
	return nil
}

func (s *storeHandlerConfig) Register(pathPrefix string, router *mux.Router) {
	// Here you register your endpoints to the router
	router.HandleFunc(pathPrefix, s.createStore).Methods("POST")
}

func (s *storeHandlerConfig) createStore(w http.ResponseWriter, req *http.Request) {
	var reqBody model.StoreRequest

	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		log.Printf("Error while parsing request body: %s", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.service.CreateStore(service.CreateStoreProps{
		Context: req.Context(),
		Body:    reqBody,
	})
	if err != nil {
		log.Printf("Error while creating store: %s", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
