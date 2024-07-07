package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) getFullInfoAboutCityAndDate(c *gin.Context) {
	city, ok := c.GetQuery("city")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query params must contain city"})
		return
	}
	date, ok := c.GetQuery("date")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query params must contain date"})
		return
	}
	t, err := parseDate(date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}
	res, err := h.service.WeatherService.GetForecastByCityNameAndDate(city, t)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
