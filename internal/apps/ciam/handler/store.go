package handler

import (
	"encoding/json"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
	"net/http"
)

type StoreHandler struct {
	repo *repository.StoreRepository
}

func NewStoreHandler(dbClient *client.SqlClient) *StoreHandler {
	return &StoreHandler{
		repo: repository.NewStoreRepository(dbClient),
	}
}

func (h *StoreHandler) CreateStore(w http.ResponseWriter, r *http.Request) {
	var storeReq model.StoreRequest
	if err := json.NewDecoder(r.Body).Decode(&storeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(storeReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
