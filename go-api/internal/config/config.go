package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Cargar .env solo si el archivo realmente existe
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  Error loading .env file:", err)
		} else {
			log.Println("✅  .env file loaded successfully")
		}
	} else {
		log.Println("ℹ️  No .env file found (using environment variables from Docker or system)")
	}
}
