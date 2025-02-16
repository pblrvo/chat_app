package main

import (
	"chat_app/internal/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	db.ConnectDB(uri)
	log.Println("Server is running...")
}
