package service

import (
	"context"
	"wb/internal/model"
	"wb/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return OrderService{repo: repo}
}

func (s *OrderService) ProcessOrder(ctx context.Context, order model.Order) error {
	return s.repo.SaveOrder(ctx, order)
}

func (s *OrderService) GetOrder(ctx context.Context, uid string) (model.Order, bool, error) {
	return s.repo.GetOrderByID(ctx, uid)
}
