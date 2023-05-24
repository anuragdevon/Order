package repository_test

import (
	"testing"

	"order/pkg/db"
	"order/pkg/repository"
	"order/pkg/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository(t *testing.T) {
	t.Run("CreateOrder", func(t *testing.T) {
		testDB := setupTestDB()

		handler := db.Database{DB: testDB}

		order := &models.Order{
			ItemId: 123,
			UserId: 456,
		}

		err := repository.CreateOrder(&handler, order)
		if err != nil {
			t.Errorf("Failed to create order: %v", err)
		}
	})

	t.Run("DeleteOrder", func(t *testing.T) {
		testDB := setupTestDB()

		handler := db.Database{DB: testDB}

		orderID := int64(789)
		err := repository.DeleteOrder(&handler, orderID)
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
