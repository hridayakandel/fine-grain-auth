// internal/apps/ciam/handler/store.go

package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository"
	"net/http"
)

func CreateStoreHandler(w http.ResponseWriter, r *http.Request) {
	var storeReq model.StoreRequest
	if err := json.NewDecoder(r.Body).Decode(&storeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Assign a unique UUID to the store ID
	storeReq.ID = uuid.New().String()
	if err := repository.CreateStoreRepo(storeReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
