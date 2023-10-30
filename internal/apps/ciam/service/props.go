package service

import (
	"context"
	"github.com/hridayakandel/fine-grain-auth/internal/apps/ciam/model"
)

type CreateStoreProps struct {
	Context context.Context
	Body    model.StoreRequest
}
