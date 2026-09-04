package main

import (
	"log"

	"contextualgamegenerator/handlers"
	"contextualgamegenerator/weather"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	weather.CheckConnection()

	r := gin.Default()

	// Setup routes
	handlers.SetupRoutes(r)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
