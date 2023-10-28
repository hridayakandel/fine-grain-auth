package repository

import (
	"github.com/google/uuid"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
)

type StoreRepository struct {
	dbClient *client.SqlClient
}

func NewStoreRepository(dbClient *client.SqlClient) *StoreRepository {
	return &StoreRepository{
		dbClient: dbClient,
	}
}

func (r *StoreRepository) Create(storeReq model.StoreRequest) error {
	// Generate a UUID for the new store
	storeID := uuid.New().String()
	query := `
		INSERT INTO store (id, name, created_at)
		VALUES ($1, $2, NOW())
	`
	_, err := r.dbClient.Conn.Exec(query, storeID, storeReq.Name)
	return err
}
