package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary Get all cities
// @Description Get the list of all cities
// @Tags cities
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 502 {object} map[string]interface{}
// @Router /api/city [get]
func (h *Handler) getAllCity(c *gin.Context) {
	cities, err := h.service.CityService.GetAllCity()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, cities)
}

// @Summary Get information about a city
// @Description Get information about a city by name
// @Tags cities
// @Produce json
// @Param name path string true "City Name"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 502 {object} map[string]interface{}
// @Router /api/city/{name} [get]
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
