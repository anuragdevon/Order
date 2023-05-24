package repository

import (
	"testing"

	"order/pkg/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository(t *testing.T) {
	t.Run("CreateOrder method to successfully create database entry for new order", func(t *testing.T) {
		testDB := setupTestDB()

		handler := Database{DB: testDB}

		order := &models.Order{
			ItemId: 123,
			UserId: 456,
		}

		err := handler.CreateOrder(order)
		if err != nil {
			t.Errorf("Failed to create order: %v", err)
		}
	})

	t.Run("DeleteOrder to successfully delete an order for failure occured from InventorySvc", func(t *testing.T) {
		testDB := setupTestDB()

		handler := Database{DB: testDB}

		orderID := int64(789)
		err := handler.DeleteOrder(orderID)
		if err != nil {
			t.Errorf("Failed to delete order: %v", err)
		}
	})
}

func setupTestDB() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/testdb2?sslmode=disable"

	testDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}

	testDB.AutoMigrate(&models.Order{})

	return testDB
}
