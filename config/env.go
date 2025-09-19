package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if _, ok := os.LookupEnv("RAILWAY_ENVIRONMENT"); ok {
		return
	}
	env := os.Getenv("APP_ENV")
	if env == "" || env == "local" || env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading env file")
		}
	}
}
