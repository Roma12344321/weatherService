package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"weatherService/pkg/model"
	"weatherService/pkg/service"
	"weatherService/pkg/service/mocks"
)

func TestHandler_addToFavourite(t *testing.T) {
	mockFavouriteService := new(mocks.FavouriteService)
	mockFavouriteService.On("AddCityToFavourite", "New York", 1).Return(nil)
	h := &Handler{
		service: &service.Service{
			FavouriteService: mockFavouriteService,
		},
	}
	personIdentity := func(c *gin.Context) {
		c.Set("person_id", 1)
		c.Next()
	}
	r := gin.Default()
	r.POST("/favourite", personIdentity, h.addToFavourite)
	req, _ := http.NewRequest("POST", "/favourite?city=New York", nil)
	req.Header.Set("Authorization", "Bearer valid_token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"success":"New York was added to favourite"}`, w.Body.String())
	mockFavouriteService.AssertExpectations(t)
}

func TestHandler_getAllFavouriteCity(t *testing.T) {
	mockFavouriteService := new(mocks.FavouriteService)
	mockCities := []model.City{
		{Id: 1, Name: "New York", Country: "USA"},
		{Id: 2, Name: "Los Angeles", Country: "USA"},
	}
	mockFavouriteService.On("GetAllFavouriteCity", 1).Return(mockCities, nil)
	h := &Handler{
		service: &service.Service{
			FavouriteService: mockFavouriteService,
		},
	}
	personIdentity := func(c *gin.Context) {
		c.Set("person_id", 1)
		c.Next()
	}
	r := gin.Default()
	r.GET("/favourite", personIdentity, h.getAllFavouriteCity)
	req, _ := http.NewRequest("GET", "/favourite", nil)
	req.Header.Set("Authorization", "Bearer valid_token")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[{"id":1,"name":"New York","country":"USA","lat":0,"lon":0},{"id":2,"name":"Los Angeles","country":"USA","lat":0,"lon":0}]`, w.Body.String())
	mockFavouriteService.AssertExpectations(t)
}
