package handler

import (
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
	mockDateStr := "2024-07-08T14:00:00Z"
	mockDate, _ := time.Parse(time.RFC3339, mockDateStr)
	mockWeather := model.WeatherForecast{
		Id:     1,
		Date:   mockDate,
		Temp:   22.5,
		Data:   json.RawMessage(`{"weather":"sunny"}`),
		CityID: 1,
		City: &model.City{
			Id:      1,
			Name:    "moscow",
			Country: "Ru",
		},
	}
	mockWeatherService.On("GetForecastByCityNameAndDate", "moscow", mockDate).Return(mockWeather, nil)
	handler := &Handler{
		service: &service.Service{
			WeatherService: mockWeatherService,
		},
	}
	r := gin.Default()
	r.GET("/weather", handler.getFullInfoAboutCityAndDate)
	req, err := http.NewRequest("GET", "/weather?city=moscow&date=2024-07-08T14:00:00Z", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedJSON := `{"weather":"sunny"}`
	assert.JSONEq(t, expectedJSON, w.Body.String())
	mockWeatherService.AssertExpectations(t)
}
