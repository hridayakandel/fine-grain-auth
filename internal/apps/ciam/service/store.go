package service

import (
	"github.com/google/uuid"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository/types"
	"time"
)

// StoreServiceConfig houses the dependencies for our store service.
type storeServiceConfig struct {
	repo repository.StoreRepo
}

// StoreService defines the methods our service should implement.
type StoreService interface {
	CreateStore(props CreateStoreProps) (model.StoreResponse, error)
}

// NewStoreService creates a new store service with the given repository.
func NewStoreService(repo repository.StoreRepo) StoreService {
	return &storeServiceConfig{
		repo: repo,
	}
}

// CreateStore implements the CreateStore method of the StoreService interface.
func (s *storeServiceConfig) CreateStore(props CreateStoreProps) (model.StoreResponse, error) {
	store, err := s.repo.CreateStore(repository.CreateStoreProps{
		Context: props.Context,
		Args: types.StoreCreateArg{
			ID:        uuid.NewString(),
			Name:      props.Body.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
	if err != nil {
		return model.StoreResponse{}, err
	}

	return store, nil
}
