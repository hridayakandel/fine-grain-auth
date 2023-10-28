package repository

import (
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
)

func CreateStoreRepo(storeReq model.StoreRequest) error {
	// Insert ID as well
	query := "INSERT INTO store (id, name, created_at) VALUES ($1, $2, NOW())"
	_, err := client.DatabaseInstance.Conn.Exec(query, storeReq.ID, storeReq.Name)
	return err
}
