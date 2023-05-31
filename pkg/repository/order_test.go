package repository

import (
	"testing"

	"order/pkg/repository/models"

	"github.com/stretchr/testify/assert"
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

	t.Run("GetOrder to successfully retrieve an existing order from the database", func(t *testing.T) {
		order := &models.Order{
			ItemId:   456,
			UserId:   789,
			Quantity: 3,
		}
		db.CreateOrder(order)

		retrievedOrder, err := db.GetOrder(order.Id, order.UserId)
		assert.Nil(t, err)
		assert.NotNil(t, retrievedOrder)
		assert.Equal(t, order.ItemId, retrievedOrder.ItemId)
	})

	t.Run("GetOrder to return nil when order does not exist", func(t *testing.T) {
		nonExistentOrderID := int64(999999)
		userId := int64(55)
		retrievedOrder, err := db.GetOrder(nonExistentOrderID, userId)
		assert.NotNil(t, err)
		assert.Nil(t, retrievedOrder)
	})

	t.Run("GetOrder to return nil when requested user does not match with order userId", func(t *testing.T) {
		order := &models.Order{
			ItemId:   456,
			UserId:   789,
			Quantity: 3,
		}
		db.CreateOrder(order)

		retrievedOrder, err := db.GetOrder(order.Id, 777)
		assert.NotNil(t, err)
		assert.Nil(t, retrievedOrder)
	})

	t.Run("GetOrdersByUserID to successfully retrieve existing orders for a user from the database", func(t *testing.T) {
		userID := int64(123)

		orders := []*models.Order{
			{ItemId: 1, UserId: userID},
			{ItemId: 2, UserId: userID},
			{ItemId: 3, UserId: userID},
		}

		for _, order := range orders {
			db.CreateOrder(order)
		}

		retrievedOrders, err := db.GetOrdersByUserID(userID)
		assert.Nil(t, err)
		assert.NotNil(t, retrievedOrders)
		assert.Len(t, retrievedOrders, len(orders))

		for _, order := range retrievedOrders {
			assert.Equal(t, userID, order.UserId)
		}
	})

	t.Run("GetOrdersByUserID to return empty when no orders exist for the user", func(t *testing.T) {
		userID := int64(999)

		retrievedOrders, err := db.GetOrdersByUserID(userID)
		assert.Nil(t, err)
		assert.NotNil(t, retrievedOrders)
		assert.Empty(t, retrievedOrders)
	})
}
