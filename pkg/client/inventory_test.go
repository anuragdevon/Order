package client

import (
	"order/pkg/client/mocks"
	"testing"
)

func TestInventoryServiceClient(t *testing.T) {
	t.Run("GetItem method to return status 200 OK for successful grpc call to inventory service for getItem details", func(t *testing.T) {
		mockClient := &mocks.MockInventoryServiceClient{}

		client := &InventoryServiceClient{
			Client: mockClient,
		}

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

	t.Run("UpdateItem method to return status 200 OK for successful grpc call to inventory service for update item quanity", func(t *testing.T) {
		mockClient := &mocks.MockInventoryServiceClient{}

		client := &InventoryServiceClient{
			Client: mockClient,
		}

		itemID := int64(123)
		quantity := int64(20)
		price := int64(200)
		name := "New Item Name"
		response, err := client.UpdateItem(itemID, quantity, price, name)

		if err != nil {
			t.Errorf("Failed to update item: %v", err)
		}

		if response.Status != 200 {
			t.Errorf("Unexpected status code: %d", response.Status)
		}

		if response.Data == nil || response.Data.Id != itemID || response.Data.Name != name || response.Data.Quantity != quantity || response.Data.Price != price {
			t.Error("Invalid updated item data")
		}
	})
}
