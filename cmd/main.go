package main

import (
	"log"
	"github.com/andreybrumatti/crud-init/internal/database/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.Connect()
}
