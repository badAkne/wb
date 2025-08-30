package service

import (
	"context"
	"wb/internal/model"
)

type OrderService interface {
	ProcessOrder(ctx context.Context, order model.Order) error
	GetOrder(ctx context.Context, uid string) (model.Order, bool, error)
}
