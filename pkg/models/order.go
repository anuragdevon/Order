package models

type Order struct {
	Id     int64 `json:"id" gorm:"primaryKey"`
	Price  int64 `json:"price"`
	ItemId int64 `json:"item_id"`
	UserId int64 `json:"user_id"`
}
