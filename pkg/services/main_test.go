package services

import (
	"log"
	"order/pkg/repository"
	"os"
	"testing"
)

var db *repository.Database

func TestMain(m *testing.M) {
	db = &repository.Database{}
	err := db.Connect("testdb2")
	if err != nil {
		log.Panic("failed to connect to database:", err)
	}

	exitCode := m.Run()

	db.Close()

	os.Exit(exitCode)
}
