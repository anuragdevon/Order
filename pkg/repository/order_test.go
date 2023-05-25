package repository

import (
	"testing"

	"order/pkg/repository/models"
)

func TestRepository(t *testing.T) {
	t.Run("CreateOrder method to successfully create database entry for new order", func(t *testing.T) {
		order := &models.Order{
			ItemId: 123,
			UserId: 456,
		}
		err := db.CreateOrder(order)
		if err != nil {
			t.Errorf("Failed to create order: %v", err)
		}
	})

	t.Run("DeleteOrder to successfully delete an order from database", func(t *testing.T) {

		order := &models.Order{
			ItemId: 789,
			UserId: 456,
		}
		db.CreateOrder(order)

		err := db.DeleteOrder(order.Id)
		if err != nil {
			t.Errorf("Failed to delete order: %v", err)
		}
	})
}
