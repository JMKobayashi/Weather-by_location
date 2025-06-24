package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/jjkobayashi/weather-service/internal/models"
)

type WeatherService struct {
	weatherAPIKey string
}

func NewWeatherService(weatherAPIKey string) *WeatherService {
	return &WeatherService{
		weatherAPIKey: weatherAPIKey,
	}
}

func (s *WeatherService) GetWeatherByZipcode(zipcode string) (*models.WeatherResponse, error) {
	// Validar formato do CEP
	if !isValidZipcode(zipcode) {
		return nil, fmt.Errorf("invalid zipcode")
	}

	// Buscar localização pelo CEP
	location, err := s.getLocationByZipcode(zipcode)
	if err != nil {
		return nil, fmt.Errorf("can not find zipcode")
	}

	// Buscar temperatura pela localização
	tempC, err := s.getTemperatureByLocation(location.Localidade)
	if err != nil {
		return nil, err
	}

	// Converter temperaturas
	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	return &models.WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}

func (s *WeatherService) getLocationByZipcode(zipcode string) (*models.ViaCEPResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("zipcode not found")
	}

	var location models.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return &location, nil
}

func (s *WeatherService) getTemperatureByLocation(location string) (float64, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", s.weatherAPIKey, location)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var weatherResp models.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return 0, err
	}

	return weatherResp.Current.TempC, nil
}

func isValidZipcode(zipcode string) bool {
	pattern := `^\d{8}$`
	match, _ := regexp.MatchString(pattern, zipcode)
	return match
}
