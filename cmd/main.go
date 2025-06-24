package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jjkobayashi/weather-service/internal/handlers"
	"github.com/jjkobayashi/weather-service/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	if weatherAPIKey == "" {
		log.Fatal("WEATHER_API_KEY environment variable is required")
	}

	weatherService := services.NewWeatherService(weatherAPIKey)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/weather/:zipcode", weatherHandler.GetWeather)

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
