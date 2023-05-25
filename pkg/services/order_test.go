package services

import (
	"context"
	"net/http"
	"testing"

	"order/pkg/client"
	"order/pkg/pb"
	"order/pkg/pb/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	mockClient := &mocks.InventoryServiceClient{}
	client := &client.InventoryServiceClient{
		Client: mockClient,
	}
	orderService := &OrderService{
		InventorySvc: *client,
		db:           db.DB,
	}

	t.Run("CreateOrder method to return 200 StatusOK for successfull valid order creation", func(t *testing.T) {
		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data: &pb.GetItemData{
				Id:       123,
				Quantity: 10,
			},
		}, nil)

		request := &pb.CreateOrderRequest{
			ItemId:   123,
			UserId:   456,
			Quantity: 5,
		}

		response, err := orderService.CreateOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusCreated), response.Status, "Unexpected response status")
		assert.Equal(t, int64(1), response.Id, "Unexpected order ID")
	})

	t.Run("CreateOrder method to return http StatusConflict for order quanity more than inventory quanity", func(t *testing.T) {
		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data: &pb.GetItemData{
				Id:       123,
				Quantity: 10,
			},
		}, nil)

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
	})
}
