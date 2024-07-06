package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getAllCity(c *gin.Context) {
	cities, err := h.service.CityService.GetAllCity()
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, cities)
}

func (h *Handler) getInfoAboutCity(c *gin.Context) {
	city, ok := c.Params.Get("name")
	if !ok {
		c.JSON(http.StatusBadRequest, "expected url param")
		return
	}
	res, err := h.service.WeatherService.GetForecastByCityName(city)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, res)
}
