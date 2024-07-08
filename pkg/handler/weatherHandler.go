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

func (h *Handler) getFullInfoAboutCityAndDate(c *gin.Context) {
	var input inputCityAndDate
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json format"})
		return
	}
	t, err := parseDate(input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}
	res, err := h.service.WeatherService.GetForecastByCityNameAndDate(input.City, t)
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
	t, err := time.Parse(time.DateTime, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
