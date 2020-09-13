package main

import (
	"log"
	"os"

	"github.com/data-scrape/data-scrape-server/app"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		if os.Getenv("ENV") != "Prod" {
			log.Fatalf("Couldn't load env variables in dev: %v", err)
		}
		log.Printf("Error loading .env file")
	}
}

func main() {
	app.SetupApp()
}
