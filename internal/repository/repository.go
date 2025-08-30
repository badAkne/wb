package repository

import (
	"context"
	"wb/internal/model"
)

type OrderRepository interface {
	LoadCacheFromDB(ctx context.Context) error
	SaveOrder(ctx context.Context,
		order model.Order) error
	GetOrderByID(ctx context.Context, id string) (model.Order, bool, error)
}
