package handler

import (
	"github.com/gin-gonic/gin"
	"weatherService/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/city", h.getAllCity)
		api.GET("/city/:name", h.getInfoAboutCity)
	}
	return router
}
