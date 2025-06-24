package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jjkobayashi/weather-service/internal/handlers"
	"github.com/jjkobayashi/weather-service/internal/models"
	"github.com/jjkobayashi/weather-service/internal/services"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	weatherService := services.NewWeatherService("test-key")
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	r := gin.Default()
	r.GET("/weather/:zipcode", weatherHandler.GetWeather)

	return r
}

func TestWeatherEndpoint(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name           string
		zipcode        string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "invalid zipcode format",
			zipcode:        "1234567",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": "invalid zipcode",
			},
		},
		{
			name:           "invalid zipcode with letters",
			zipcode:        "1234567a",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": "invalid zipcode",
			},
		},
		{
			name:           "non-existent zipcode",
			zipcode:        "99999999",
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error": "can not find zipcode",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/weather/"+tt.zipcode, nil)
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Errorf("Failed to unmarshal response: %v", err)
			}

			if response["error"] != tt.expectedBody["error"] {
				t.Errorf("Expected error %s, got %s", tt.expectedBody["error"], response["error"])
			}
		})
	}
}

func TestWeatherResponseStructure(t *testing.T) {
	// Teste para verificar se a estrutura da resposta está correta
	response := models.WeatherResponse{
		TempC: 25.0,
		TempF: 77.0,
		TempK: 298.0,
	}

	// Verificar se os campos estão presentes
	if response.TempC == 0 {
		t.Error("TempC should not be zero")
	}
	if response.TempF == 0 {
		t.Error("TempF should not be zero")
	}
	if response.TempK == 0 {
		t.Error("TempK should not be zero")
	}

	// Verificar se as conversões estão corretas
	expectedF := response.TempC*1.8 + 32
	expectedK := response.TempC + 273

	if response.TempF != expectedF {
		t.Errorf("Expected TempF %f, got %f", expectedF, response.TempF)
	}
	if response.TempK != expectedK {
		t.Errorf("Expected TempK %f, got %f", expectedK, response.TempK)
	}
}
