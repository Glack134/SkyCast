package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	WeatherHandler *WeatherHandler
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	weather := router.Group("/main")
	{
		weather.GET("/", h.WeatherHandler.Main)
		weather.GET("/search", h.WeatherHandler.Search)
	}
	return router
}
