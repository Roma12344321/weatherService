package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"weatherService/pkg/dto"
	"weatherService/pkg/model"
	"weatherService/pkg/service"
	"weatherService/pkg/service/mocks"
)

func TestHandler_getAllCity(t *testing.T) {
	mockCityService := new(mocks.CityService)
	mockCities := []model.City{
		{Id: 1, Name: "New York", Country: "USA", Lat: 0, Lon: 0},
		{Id: 2, Name: "Los Angeles", Country: "USA", Lat: 0, Lon: 0},
	}
	mockCityService.On("GetAllCity").Return(mockCities, nil)
	h := &Handler{
		service: &service.Service{
			CityService: mockCityService,
		},
	}
	r := gin.Default()
	r.GET("/city", h.getAllCity)
	req, _ := http.NewRequest("GET", "/city", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[{"id":1,"name":"New York","country":"USA","lat":0,"lon":0},{"id":2,"name":"Los Angeles","country":"USA","lat":0,"lon":0}]`, w.Body.String())
	mockCityService.AssertExpectations(t)
}

func TestHandler_getInfoAboutCity(t *testing.T) {
	mockWeatherService := new(mocks.WeatherService)
	mockCityWeather := dto.WeatherDto{
		Name:    "New York",
		Country: "USA",
		AvgTemp: 22.5,
		Dates:   []string{"2024-07-08 14:00:00"},
	}
	mockWeatherService.On("GetForecastByCityName", "New York").Return(mockCityWeather, nil)
	h := &Handler{
		service: &service.Service{
			WeatherService: mockWeatherService,
		},
	}
	r := gin.Default()
	r.GET("/city/:name", h.getInfoAboutCity)
	req, _ := http.NewRequest("GET", "/city/New York", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"name":"New York","country":"USA","avg_temp":22.5,"dates":["2024-07-08 14:00:00"]}`, w.Body.String())
	mockWeatherService.AssertExpectations(t)
}
