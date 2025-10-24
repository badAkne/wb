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

	if err := s.repo.SaveOrder(ctx, order); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) GetOrder(ctx context.Context, uid string) (model.Order, bool, error) {
	order, found, err := s.repo.GetOrderByID(ctx, uid)

	if err != nil {
		return model.Order{}, false, err
	}

	return order, found, err
}
