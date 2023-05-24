package repository

import (
	"order/pkg/db"
	"order/pkg/repository/models"
)

func CreateOrder(h *db.Database, order *models.Order) error {
	return h.DB.Create(order).Error
}

func DeleteOrder(h *db.Database, orderId int64) error {
	return h.DB.Delete(&models.Order{}, orderId).Error
}
