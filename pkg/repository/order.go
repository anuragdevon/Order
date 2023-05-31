package repository

import (
	"order/pkg/repository/models"
)

func (db *Database) CreateOrder(order *models.Order) error {
	return db.DB.Create(order).Error
}

func (db *Database) DeleteOrder(orderId int64) error {
	return db.DB.Delete(&models.Order{}, orderId).Error
}

func (db *Database) GetOrder(orderId int64) (*models.Order, error) {
	order := &models.Order{}
	result := db.DB.First(order, orderId)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}
