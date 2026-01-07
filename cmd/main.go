package main

import (
	"log"
	"os"

	"github.com/andreybrumatti/crud-init/internal/database/mysql"
	"github.com/andreybrumatti/crud-init/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	mysql.Connect()
	db := mysql.DB

	r := gin.Default()
	routes.SetupRoutes(r, db)

    if err := r.Run(":" + port); err != nil {
        log.Fatal("Failed to run server: ", err)
    }

	log.Println("Server running on port " + port)
}