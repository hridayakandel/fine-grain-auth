package repository

import (
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository/queries"
	"github.com/hridayakandel/fine-grain-auth/internal/pkg/db/sql/client"
)

type storeRepoConfig struct {
	dbClient *client.SqlClient
}

// StoreRepo defines the methods our repository should implement.
type StoreRepo interface {
	CreateStore(props CreateStoreProps) (model.StoreResponse, error)
}

// NewStoreRepo creates a new store repository with the given database client.
func NewStoreRepo(dbClient *client.SqlClient) StoreRepo {
	return &storeRepoConfig{
		dbClient: dbClient,
	}
}

// CreateStore implements the CreateStore method of the StoreRepo interface.
func (r *storeRepoConfig) CreateStore(props CreateStoreProps) (model.StoreResponse, error) {
	//storeID := props.Args.ID
	//name := props.Args.Name
	//query := queries.CreateStore
	//var createdAt time.Time
	//err := r.dbClient.Conn.QueryRow(query, storeID, name).Scan(&createdAt)
	//if err != nil {
	//	return model.StoreResponse{}, err
	//}
	//
	//return model.StoreResponse{
	//	ID:        storeID,
	//	Name:      props.Args.Name,
	//	CreatedAt: createdAt,
	//}, nil
	var res model.StoreResponse

	err := r.dbClient.NamedQueryExecution(&client.PrepareNamedWithContextProps{
		Ctx:   props.Context,
		Query: queries.CreateStore,
		Dest:  &res,
		Args:  props.Args,
	})

	if err != nil {
		return model.StoreResponse{}, err
	}

	return res, nil
}
