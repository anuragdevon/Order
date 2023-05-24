package db

import (
	"log"

	"order/pkg/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func Init(url string) Database {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Order{})

	return Database{db}
}

func (h *Database) Close() {
	sqlDB, err := h.DB.DB()
	if err != nil {
		log.Println("Failed to get underlying *sql.DB:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Failed to close database connection:", err)
	}
}
