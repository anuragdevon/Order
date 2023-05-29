package repository

import (
	"order/pkg/repository/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgreSQL(t *testing.T) {
	t.Run("Connect method should connect to database successfully with valid dbname", func(t *testing.T) {
		db := &Database{}
		err := db.Connect("testdb")
		if err != nil {
			t.Errorf("error connecting to database: %s", err)
		}
		defer db.Close()

		if db.DB == nil {
			t.Error("failed to connect to database")
		}
	})

	t.Run("Close method should close connection to db successfully", func(t *testing.T) {
		db := &Database{}
		err := db.Connect("testdb")
		if err != nil {
			t.Errorf("error connecting to database: %s", err)
		}

		err = db.Close()
		if err != nil {
			t.Errorf("error closing database connection: %s", err)
		}
	})

	t.Run("Connect method should return error for connecting with database with invalid dbName", func(t *testing.T) {
		db := &Database{}
		err := db.Connect("invalid_db")
		if err == nil {
			t.Error("expected an error connecting to invalid database, but got nil")
		}
	})

	t.Run("Connect method should create migration tables after successful db connection", func(t *testing.T) {
		db := &Database{}
		err := db.Connect("testdb")
		assert.NoError(t, err, "Failed to connect to the database")

		assert.True(t, db.DB.Migrator().HasTable(&models.Order{}), "Migration table for Order does not exist")
	})
}
