package service

import (
	"context"
	"fmt"
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
	if order.OrderUID == "" {
		return fmt.Errorf("invalid customer id")
	}

	if order.TrackNumber == "" {
		return fmt.Errorf("invalid track number")
	}

	if order.Entry == "" {
		return fmt.Errorf("invalid entry")
	}

	if order.Delivery.Name == "" {
		return fmt.Errorf("invalid name")
	}

	if order.Delivery.Phone == "" {
		return fmt.Errorf("invalid phone")
	}

	if order.Delivery.Zip == "" {
		return fmt.Errorf("invalid zip")
	}

	if order.Delivery.City == "" {
		return fmt.Errorf("invalid city")
	}

	if order.Delivery.Address == "" {
		return fmt.Errorf("invalid address")
	}

	if order.Delivery.Region == "" {
		return fmt.Errorf("invalid region")
	}

	if order.Delivery.Email == "" {
		return fmt.Errorf("invalid email")
	}

	if order.Payment.Transaction == "" {
		return fmt.Errorf("invalid transaction")
	}

	if order.Payment.RequestId == "" {
		return fmt.Errorf("invalid request id")
	}

	if order.Payment.RequestId == "" {
		return fmt.Errorf("invalid request id")
	}

	if order.Payment.Currency == "" {
		return fmt.Errorf("invalid currency")
	}

	if order.Payment.Provider == "" {
		return fmt.Errorf("invalid provider")
	}

	if order.Payment.Amount == 0 {
		return fmt.Errorf("invalid amount")
	}

	if order.Payment.Paymentdt == 0 {
		return fmt.Errorf("invalid payment dt")
	}

	if order.Payment.Bank == "" {
		return fmt.Errorf("invalid bank")
	}

	if order.Payment.DeliveryCost == 0 {
		return fmt.Errorf("invalid delivery cost")
	}

	if order.Payment.GoodsTotal == 0 {
		return fmt.Errorf("invalid goods total")
	}

	if order.Payment.CustomFee == 0 {
		return fmt.Errorf("invalid custom fee")
	}

	for i := 0; i < len(order.Items); i++ {
		if order.Items[i].ChrtId == 0 {
			return fmt.Errorf("invalid chrt id")
		}

		if order.Items[i].TrackNumber == "" {
			return fmt.Errorf("invalid track number")
		}

		if order.Items[i].Price == 0 {
			return fmt.Errorf("invalid price")
		}

		if order.Items[i].Rid == "" {
			return fmt.Errorf("invalid rid")
		}

		if order.Items[i].Name == "" {
			return fmt.Errorf("invalid item name")
		}

		if order.Items[i].Sale == 0 {
			return fmt.Errorf("invalid sale amount")
		}

		if order.Items[i].Size == "" {
			return fmt.Errorf("invalid size")
		}

		if order.Items[i].TotalPrice == 0 {
			return fmt.Errorf("invalid total price")
		}

		if order.Items[i].NmId == 0 {
			return fmt.Errorf("invalid nmid")
		}

		if order.Items[i].Brand == "" {
			return fmt.Errorf("invalid brand name")
		}

		if order.Items[i].Status == 0 {
			return fmt.Errorf("invalid status")
		}
	}

	if order.Locale == "" {
		return fmt.Errorf("invalid locale")
	}

	if order.InternalSignature == "" {
		return fmt.Errorf("invalid internal signature")
	}

	if order.CustomerId == "" {
		return fmt.Errorf("invalid customer id")
	}

	if order.DeliveryService == "" {
		return fmt.Errorf("invalid delivery service")
	}

	if order.Sharkey == "" {
		return fmt.Errorf("invalid shar key")
	}

	if order.SmId == 0 {
		return fmt.Errorf("invalid smid")
	}

	if order.DateCreated.String() == "" {
		return fmt.Errorf("invalid creation date")
	}

	if order.OofShard == "" {
		return fmt.Errorf("invalid oof shard")
	}

	return s.repo.SaveOrder(ctx, order)
}

func (s *OrderService) GetOrder(ctx context.Context, uid string) (model.Order, bool, error) {

	return s.repo.GetOrderByID(ctx, uid)
}
