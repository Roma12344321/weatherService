package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weatherService/pkg/model"
	"weatherService/pkg/service"
	"weatherService/pkg/service/mocks"
)

func TestHandler_getFullInfoAboutCityAndDate(t *testing.T) {
	mockWeatherService := new(mocks.WeatherService)
	mockDate, _ := time.Parse(time.DateTime, "2024-07-08 14:00:00")
	mockWeather := model.WeatherForecast{
		Id:     1,
		Date:   mockDate,
		Temp:   22.5,
		Data:   json.RawMessage(`{"weather":"sunny"}`),
		CityID: 1,
		City: &model.City{
			Id:      1,
			Name:    "New York",
			Country: "USA",
		},
	}
	mockWeatherService.On("GetForecastByCityNameAndDate", "New York", mockDate).Return(mockWeather, nil)
	h := &Handler{
		service: &service.Service{
			WeatherService: mockWeatherService,
		},
	}
	r := gin.Default()
	r.GET("/weather", h.getFullInfoAboutCityAndDate)
	body, err := json.Marshal(inputCityAndDate{City: "New York", Date: "2024-07-08 14:00:00"})
	assert.NoError(t, err)
	req, err := http.NewRequest("GET", "/weather", bytes.NewReader(body))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"weather":"sunny"}`, w.Body.String())
	mockWeatherService.AssertExpectations(t)
}
