package db

import (
	"testing"

	"order/pkg/models"
)

func TestInit(t *testing.T) {
	dbURL := "postgres://postgres:postgres@localhost:5432/testdb2?sslmode=disable"

	handler := Init(dbURL)

	if handler.DB == nil {
		t.Error("Failed to initialize database connection")
	}

	order := models.Order{
		Price:  100,
		ItemId: 1,
		UserId: 1,
	}
	result := handler.DB.Create(&order)
	if result.Error != nil {
		t.Errorf("Failed to create order: %v", result.Error)
	}

	var retrievedOrder models.Order
	result = handler.DB.First(&retrievedOrder, order.Id)
	if result.Error != nil {
		t.Errorf("Failed to retrieve order: %v", result.Error)
	}

	if retrievedOrder.Price != order.Price || retrievedOrder.ItemId != order.ItemId || retrievedOrder.UserId != order.UserId {
		t.Error("Retrieved order does not match the created order")
	}
	handler.Close()
}
