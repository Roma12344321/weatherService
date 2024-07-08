package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) addToFavourite(c *gin.Context) {
	city, ok := c.GetQuery("city")
	if !ok {
		c.JSON(http.StatusBadRequest, "query params must contain city")
		return
	}
	personId, err := getPersonId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.FavouriteService.AddCityToFavourite(city, personId)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "city was not found"})
		return
	}
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": city + " was added to favourite"})
}

func (h *Handler) getAllFavouriteCity(c *gin.Context) {
	personId, err := getPersonId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.FavouriteService.GetAllFavouriteCity(personId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "server error"})
		return
	}
	c.JSON(http.StatusOK, res)
}
