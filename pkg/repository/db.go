package repository

import (
	"fmt"
	"log"

	"order/pkg/config"
	"order/pkg/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (p *Database) Connect(c *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", c.DbHost, c.DbUser, c.DbPassword, c.DbName)
	var err error
	p.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := p.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying SQL DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	p.DB.AutoMigrate(&models.Order{})

	return nil
}

func (p *Database) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		log.Println("failed to get underlying SQL DB: %w", err)
		return err
	}
	err = p.DB.Migrator().DropTable(
		&models.Order{},
	)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Failed to close database connection:", err)
	}
	return err
}
