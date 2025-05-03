package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	if err := os.MkdirAll("data", 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	log.Println("database connected")

	return nil
}
