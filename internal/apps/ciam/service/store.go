// internal/apps/ciam/service/store.go

package service

import (
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository"
)

func CreateStoreService(storeReq model.StoreRequest) error {
	return repository.CreateStoreRepo(storeReq)
}
