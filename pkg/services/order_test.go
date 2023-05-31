package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"order/pkg/client"
	"order/pkg/pb"
	"order/pkg/pb/mocks"
	"order/pkg/repository/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	t.Run("CreateOrder method to return 201 StatusCreated for successful valid order creation", func(t *testing.T) {
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data: &pb.GetItemData{
				Id:       123,
				Quantity: 10,
			},
		}, nil)

		mockClient.On("DecreaseItemQuantity", mock.Anything, mock.Anything).Return(nil, nil)

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

	t.Run("CreateOrder method to return 502 StatusBadGateway when GetItem returns an error", func(t *testing.T) {
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(nil, errors.New("GetItem error"))

		request := &pb.CreateOrderRequest{
			ItemId:   123,
			UserId:   456,
			Quantity: 5,
		}

		response, err := orderService.CreateOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusBadGateway), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "GetItem error", "Error message mismatch")
	})

	t.Run("CreateOrder method to return status 404 NotFound when the requested item is not found in the inventory", func(t *testing.T) {
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusNotFound,
			Error:  "Item not found",
			Data:   nil,
		}, nil)

		request := &pb.CreateOrderRequest{
			ItemId:   123,
			UserId:   456,
			Quantity: 5,
		}

		response, err := orderService.CreateOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusNotFound), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "Item not found", "Unexpected error message")
	})

	t.Run("CreateOrder method to return 409 StatusConflict when quantity is insufficient", func(t *testing.T) {
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data: &pb.GetItemData{
				Id:       123,
				Quantity: 3,
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
		assert.Equal(t, int64(http.StatusConflict), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "Quantity is insufficient", "Error message mismatch")
	})

	t.Run("CreateOrder method to return 500 InternalServerError when DecreaseItemQuantity Service call fails", func(t *testing.T) {
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data: &pb.GetItemData{
				Id:       123,
				Quantity: 10,
			},
		}, nil)

		mockClient.On("DecreaseItemQuantity", mock.Anything, mock.Anything).Return(nil, errors.New("DecreaseItemQuantity error"))

		request := &pb.CreateOrderRequest{
			ItemId:   123,
			UserId:   456,
			Quantity: 5,
		}

		response, err := orderService.CreateOrder(context.Background(), request)

		assert.NoError(t, err)
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusInternalServerError), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "Internal server error", "Error message mismatch")

	})

	t.Run("GetOrder method to return status 200 OK and order data for a valid order ID", func(t *testing.T) {
		ItemId := int64(123)
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}
		newOrder := models.Order{
			ItemId:   ItemId,
			UserId:   12,
			Quantity: 2,
		}
		db.CreateOrder(&newOrder)

		expectedGetItemData := &pb.GetItemData{
			Id:    123,
			Name:  "Item-1",
			Price: 30,
		}
		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data:   expectedGetItemData,
		}, nil)

		request := &pb.GetOrderRequest{
			Id: newOrder.Id,
		}

		response, err := orderService.GetOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusOK), response.Status, "Unexpected response status")

		assert.Equal(t, newOrder.Id, response.Data.Id, "Unexpected order ID")
		assert.Equal(t, expectedGetItemData.Id, response.Data.ItemId, "Unexpected item ID")
		assert.Equal(t, expectedGetItemData.Name, response.Data.Name, "Unexpected item name")
		assert.Equal(t, newOrder.Quantity, response.Data.Quantity, "Unexpected order quantity")
		assert.Equal(t, expectedGetItemData.Price, response.Data.Price, "Unexpected item price")
	})

	t.Run("GetOrder method to return 404 NotFound when requested user does not match with order userId", func(t *testing.T) {
		ItemId := int64(123)
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}
		newOrder := models.Order{
			ItemId:   ItemId,
			UserId:   12,
			Quantity: 2,
		}
		db.CreateOrder(&newOrder)

		expectedGetItemData := &pb.GetItemData{
			Id:    123,
			Name:  "Item-1",
			Price: 30,
		}
		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data:   expectedGetItemData,
		}, nil)

		request := &pb.GetOrderRequest{
			Id: newOrder.Id,
		}

		response, err := orderService.GetOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusNotFound), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "Order not found", "Error message mismatch")
	})

	t.Run("GetOrder method to return 404 NotFound for a non-existent order ID", func(t *testing.T) {
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		request := &pb.GetOrderRequest{
			Id: 999,
		}

		response, err := orderService.GetOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusNotFound), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "Order not found", "Error message mismatch")
	})

	t.Run("GetOrder method to return 502 BadGateway when GetItem returns an error", func(t *testing.T) {
		itemId := int64(44)
		mockClient := &mocks.InventoryServiceClient{}
		client := &client.InventoryServiceClient{
			Client: mockClient,
		}
		orderService := &OrderService{
			InventorySvc: *client,
			db:           db.DB,
		}

		newOrder := models.Order{
			ItemId:   itemId,
			UserId:   12,
			Quantity: 2,
		}
		db.CreateOrder(&newOrder)

		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(nil, errors.New("Inventory service error"))

		request := &pb.GetOrderRequest{
			Id: newOrder.Id,
		}

		response, err := orderService.GetOrder(context.Background(), request)

		assert.NoError(t, err, "Unexpected error")
		assert.NotNil(t, response, "Response is nil")
		assert.Equal(t, int64(http.StatusBadGateway), response.Status, "Unexpected response status")
		assert.Contains(t, response.Error, "Inventory service error", "Error message mismatch")
	})
}
