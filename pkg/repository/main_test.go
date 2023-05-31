package repository

import (
	"log"
	"order/pkg/config"
	"os"
	"testing"
)

var db *Database

func TestMain(m *testing.M) {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db = &Database{}
	err = db.Connect(&c)
	if err != nil {
		log.Panic("failed to connect to database:", err)
	}

	exitCode := m.Run()

	db.Close()

	os.Exit(exitCode)
}
