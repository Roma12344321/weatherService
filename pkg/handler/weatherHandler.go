package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type inputCityAndDate struct {
	City string `json:"city"`
	Date string `json:"date"`
}

// @Summary Get full weather info
// @Description Get full weather information for a city and date
// @Tags weather
// @Accept json
// @Produce json
// @Param city query string true "City name"
// @Param date query string true "Date (formatted as 2006-05-06T23:59:55Z)"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 502 {object} map[string]interface{} "Bad Gateway"
// @Router /api/weather [get]
func (h *Handler) getFullInfoAboutCityAndDate(c *gin.Context) {
	city := c.Query("city")
	date := c.Query("date")
	if city == "" || date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city and date are required query parameters"})
		return
	}
	t, err := parseDate(date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}
	res, err := h.service.WeatherService.GetForecastByCityNameAndDate(city, t)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
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
