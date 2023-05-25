package client

import (
	"net/http"
	"order/pkg/pb"
	"order/pkg/pb/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestInventoryServiceClient(t *testing.T) {
	mockClient := &mocks.InventoryServiceClient{}

	client := &InventoryServiceClient{
		Client: mockClient,
	}

	t.Run("GetItem method to return status 200 OK for successful grpc call to inventory service for getItem details", func(t *testing.T) {
		mockClient.On("GetItem", mock.Anything, mock.Anything).Return(&pb.GetItemResponse{
			Status: http.StatusOK,
			Error:  "",
			Data: &pb.GetItemData{
				Id:       123,
				Quantity: 10,
			},
		}, nil)

		itemID := int64(123)
		response, err := client.GetItem(itemID)

		if err != nil {
			t.Errorf("Failed to get item: %v", err)
		}

		if response.Status != 200 {
			t.Errorf("Unexpected status code: %d", response.Status)
		}

		if response.Data == nil || response.Data.Id != itemID {
			t.Error("Invalid item data")
		}
	})

	t.Run("DecreaseItemQuantity method to return status 200 OK for successful grpc call to inventory service for update item quanity", func(t *testing.T) {
		mockClient.On("DecreaseItemQuantity", mock.Anything, mock.Anything).Return(&pb.DecreaseItemQuantityResponse{
			Status: http.StatusOK,
			Error:  "",
		}, nil)

		itemID := int64(123)
		quantity := int64(5)
		response, err := client.DecreaseItemQuantity(itemID, quantity)

		if err != nil {
			t.Errorf("Failed to update item: %v", err)
		}

		if response.Status != 200 {
			t.Errorf("Unexpected status code: %d", response.Status)
		}
	})
}
