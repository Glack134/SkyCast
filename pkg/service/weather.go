package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Glack134/ScyCast"
)

type WeatherService struct {
	APIKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{APIKey: apiKey}
}

func (ws *WeatherService) GetWeather(city string) (*ScyCast.WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, ws.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResponse ScyCast.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}
	return &weatherResponse, nil
}
