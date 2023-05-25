package service

import (
	"context"
	"net/http"
	"testing"

	"order/pkg/client"
	"order/pkg/client/mocks"
	"order/pkg/pb"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func TestCreateOrder_Success(t *testing.T) {

	mockClient := &mocks.MockInventoryServiceClient{}

	client := &client.InventoryServiceClient{
		Client: mockClient,
	}

	mockDB := &gorm.DB{}

	orderService := &OrderService{
		InventorySvc: *client,
		db:           mockDB,
	}

	request := &pb.CreateOrderRequest{
		ItemId:   123,
		UserId:   456,
		Quantity: 5,
	}

	response, err := orderService.CreateOrder(context.Background(), request)

	// Assert the expected behavior
	assert.NoError(t, err, "Unexpected error")
	assert.NotNil(t, response, "Response is nil")
	assert.Equal(t, http.StatusCreated, response.Status, "Unexpected response status")
	assert.Equal(t, int64(1), response.Id, "Unexpected order ID")
}

func TestCreateOrder_InsufficientQuantity(t *testing.T) {
	mockClient := &mocks.MockInventoryServiceClient{}

	client := &client.InventoryServiceClient{
		Client: mockClient,
	}

	mockDB := &gorm.DB{}

	orderService := &OrderService{
		InventorySvc: *client,
		db:           mockDB,
	}

	request := &pb.CreateOrderRequest{
		ItemId:   123,
		UserId:   456,
		Quantity: 15,
	}

	response, err := orderService.CreateOrder(context.Background(), request)

	assert.NoError(t, err, "Unexpected error")
	assert.NotNil(t, response, "Response is nil")
	assert.Equal(t, int64(http.StatusConflict), response.Status, "Unexpected response status")
	assert.Contains(t, response.Error, "Quantity is insufficient", "Unexpected error message")
}
