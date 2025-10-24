package service

import (
	"context"
	"fmt"
	"testing"

	"wb/internal/model"
	"wb/internal/repository"
	service "wb/internal/service/order"
	testdata "wb/internal/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrder(t *testing.T) {

	mockRepo := repository.NewMockOrderRepository(t)

	usecase := service.NewOrderService(mockRepo)

	ctx := context.Background()
	uid := "b563feb7b2b84b6test"
	expResp := testdata.Good

	mockRepo.On("GetOrderByID", ctx, uid).Return(expResp, true, nil).Times(1)

	resp, found, err := usecase.GetOrder(ctx, uid)

	require.NoError(t, err)
	assert.Equal(t, expResp, resp)
	assert.Equal(t, found, true)
}
func TestGetOrderError(t *testing.T) {
	mockRepo := repository.NewMockOrderRepository(t)

	usecase := service.NewOrderService(mockRepo)

	ctx := context.Background()
	uid := "testuid"
	expResp := model.Order{}
	repoErr := fmt.Errorf("unable_to_find_an_order")

	mockRepo.On("GetOrderByID", ctx, uid).Return(expResp, false, repoErr).Times(1)

	order, found, err := usecase.GetOrder(ctx, uid)

	require.Error(t, err)
	assert.Equal(t, expResp, order)
	assert.Equal(t, false, found)
}

func TestProcessOrder(t *testing.T) {
	mockRepo := repository.NewMockOrderRepository(t)

	usecase := service.NewOrderService(mockRepo)

	ctx := context.Background()
	order := testdata.Good

	mockRepo.On("SaveOrder", ctx, order).Return(nil).Times(1)

	err := usecase.ProcessOrder(ctx, order)

	require.NoError(t, err)
}

func TestProcessOrderErr(t *testing.T) {
	mockRepo := repository.NewMockOrderRepository(t)

	usecase := service.NewOrderService(mockRepo)

	ctx := context.Background()
	order := testdata.Good

	mockRepo.On("SaveOrder", ctx, order).Return(fmt.Errorf("error_from_repo")).Times(1)

	err := usecase.ProcessOrder(ctx, order)

	require.Error(t, err)
}
