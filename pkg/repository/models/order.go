package models

type Order struct {
	Id     int64 `gorm:"primaryKey"`
	ItemId int64 `gorm:"column:item_id"`
	UserId int64 `gorm:"column:user_id"`
}
