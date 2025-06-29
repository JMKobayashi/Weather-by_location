package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

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
		log.Printf("Invalid zipcode: %s", zipcode)
		return nil, fmt.Errorf("invalid zipcode")
	}

	// Buscar localização pelo CEP
	location, err := s.getLocationByZipcode(zipcode)
	if err != nil {
		log.Printf("Erro ao buscar localização para o CEP %s: %v", zipcode, err)
		return nil, fmt.Errorf("can not find zipcode")
	}

	// Verificar se a localização foi encontrada
	if location.Localidade == "" {
		log.Printf("Localidade não encontrada para o CEP: %s", zipcode)
		return nil, fmt.Errorf("can not find zipcode")
	}

	// Buscar temperatura pela localização
	tempC, err := s.getTemperatureByLocation(location.Localidade)
	if err != nil {
		log.Printf("Erro ao buscar temperatura para a localidade %s (CEP %s): %v", location.Localidade, zipcode, err)
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
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	log.Printf("Consultando ViaCEP: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro de requisição HTTP para ViaCEP: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("ViaCEP retornou status %d para o CEP %s", resp.StatusCode, zipcode)
		// Log do corpo da resposta para debug
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Resposta do ViaCEP: %s", string(bodyBytes))

		if resp.StatusCode == http.StatusBadGateway || resp.StatusCode == http.StatusServiceUnavailable {
			return nil, fmt.Errorf("viacep service temporarily unavailable")
		}
		return nil, fmt.Errorf("zipcode not found")
	}

	var location models.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		log.Printf("Erro ao decodificar resposta do ViaCEP: %v", err)
		return nil, err
	}

	// Verificar se o CEP foi encontrado (ViaCEP retorna erro quando não encontra)
	if location.Cep == "" {
		log.Printf("ViaCEP não encontrou o CEP: %s", zipcode)
		return nil, fmt.Errorf("zipcode not found")
	}

	return &location, nil
}

func (s *WeatherService) getTemperatureByLocation(location string) (float64, error) {
	// Usar encoding URL correto para caracteres especiais
	encodedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", s.weatherAPIKey, encodedLocation)
	log.Printf("Consultando WeatherAPI: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Erro de requisição HTTP para WeatherAPI: %v", err)
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("WeatherAPI retornou status %d para localidade %s", resp.StatusCode, location)
		// Log do corpo da resposta para debug
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Resposta da WeatherAPI: %s", string(bodyBytes))
		return 0, fmt.Errorf("weather API error: %d", resp.StatusCode)
	}

	var weatherResp models.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		log.Printf("Erro ao decodificar resposta da WeatherAPI: %v", err)
		return 0, err
	}

	return weatherResp.Current.TempC, nil
}

func isValidZipcode(zipcode string) bool {
	// Remove hífens e espaços
	cleanZipcode := strings.ReplaceAll(strings.ReplaceAll(zipcode, "-", ""), " ", "")

	// Verifica se tem exatamente 8 dígitos
	pattern := `^\d{8}$`
	match, _ := regexp.MatchString(pattern, cleanZipcode)
	return match
}
