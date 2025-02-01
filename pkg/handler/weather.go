package handler

import (
	"github.com/Glack134/ScyCast/pkg/service"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	weatherService *service.WeatherService
}

func NewWeatherHandler(weatherService *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{weatherService: weatherService}
}

func (h *WeatherHandler) Main(c *gin.Context) {
	cities := []string{"London", "Paris", "New York", "Tokyo"}

	weatherData := make(map[string]interface{})
	for _, city := range cities {
		weather, err := h.weatherService.GetWeather(city)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch weather data"})
			return
		}
		weatherData[city] = weather
	}

	c.JSON(200, gin.H{"weather": weatherData})
}

func (h *WeatherHandler) Search(c *gin.Context) {
	city := c.Query("city") // Получаем город из query параметра
	if city == "" {
		c.JSON(400, gin.H{"error": "City parameter is required"})
		return
	}

	weather, err := h.weatherService.GetWeather(city)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch weather data"})
		return
	}

	c.JSON(200, gin.H{"weather": weather})
}
