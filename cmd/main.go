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
	// Configurar Gin para modo release em produção
	gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting application on port %s", port)

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	if weatherAPIKey == "" {
		log.Fatal("WEATHER_API_KEY environment variable is required")
	}

	log.Println("Weather API Key configured successfully")

	weatherService := services.NewWeatherService(weatherAPIKey)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/weather/:zipcode", weatherHandler.GetWeather)

	log.Printf("Server starting on port %s", port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
