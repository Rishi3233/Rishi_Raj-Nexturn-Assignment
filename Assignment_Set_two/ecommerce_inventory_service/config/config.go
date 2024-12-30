package config

import (
	"fmt"
	"os"
	"log"
)

// LoadConfig loads environment variables or config files.
func LoadConfig() {
	// Load the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// For now, print the database URL. You can also add other configurations.
	fmt.Printf("Database URL: %s\n", dbURL)
}
