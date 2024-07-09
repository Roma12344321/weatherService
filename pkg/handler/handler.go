package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/registration", h.createPerson)
		auth.POST("/login", h.logIn)
	}
	api := router.Group("/api")
	{
		api.GET("/city", h.getAllCity)
		api.GET("/city/:name", h.getInfoAboutCity)
		api.GET("/weather", h.getFullInfoAboutCityAndDate)
		securedApi := api.Group("", h.personIdentity)
		{
			securedApi.POST("/favourite", h.addToFavourite)
			securedApi.GET("/favourite", h.getAllFavouriteCity)
		}
	}
	return router
}
