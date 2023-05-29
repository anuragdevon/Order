package repository

import (
	"log"
	"os"
	"testing"
)

var db *Database

func TestMain(m *testing.M) {
	db = &Database{}
	err := db.Connect("testdb")
	if err != nil {
		log.Panic("failed to connect to database:", err)
	}

	exitCode := m.Run()

	db.Close()

	os.Exit(exitCode)
}
