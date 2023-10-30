package repository

import (
	"context"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/repository/types"
)

type CreateStoreProps struct {
	Context context.Context
	Args    types.StoreCreateArg
}
