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
