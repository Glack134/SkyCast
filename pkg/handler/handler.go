package handler

import (
	"github.com/Glack134/SkyCast/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter(c *gin.Context) {
	router := gin.New()

	main := router.Group("/")
	{
		search := main.Group("/search")
		{
			search.GET("/Sky", h.SkyBlue)
			search.GET("/search", h.Search)
		}
	}
}
