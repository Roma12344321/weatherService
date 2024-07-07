package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
	"weatherService/pkg/mapper"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
	"weatherService/pkg/repository/mocks"
)

func TestWeatherService_SaveWeatherForeCast(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepository)
	mockCityRepo := new(mocks.CityRepository)
	client := &http.Client{}
	weatherService := NewWeatherServiceImpl(context.Background(), &repository.Repository{
		WeatherRepository: mockWeatherRepo,
		CityRepository:    mockCityRepo,
	}, client)
	cities := []model.City{
		{Id: 1, Name: "city1", Lat: 10.0, Lon: 20.0},
		{Id: 2, Name: "city2", Lat: 30.0, Lon: 40.0},
	}
	mockWeatherRepo.On("DeleteOldDates").Return(nil)
	mockWeatherRepo.On("SaveWeatherForeCast", mock.AnythingOfType("*model.WeatherForecast")).Return(nil)
	result, err := weatherService.SaveWeatherForeCast(cities, "02c02ade5939988a2cffd1fb499fa3fc")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockWeatherRepo.AssertExpectations(t)
}

func TestWeatherService_GetForecastByCityName(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepository)
	client := &http.Client{}
	weatherService := NewWeatherServiceImpl(context.Background(), &repository.Repository{
		WeatherRepository: mockWeatherRepo,
	}, client)
	cityName := "city1"
	forecasts := []model.WeatherForecast{
		{CityID: 1, Date: time.Now(), Temp: 25.0, City: &model.City{Name: cityName, Country: "Country1"}},
	}
	mockWeatherRepo.On("GetWeatherForeCastByCityName", cityName).Return(forecasts, nil)
	result, err := weatherService.GetForecastByCityName(cityName)
	expectedDto := mapper.MapWeatherForecastListToWeatherDto(forecasts)
	assert.NoError(t, err)
	assert.Equal(t, expectedDto, result)
	mockWeatherRepo.AssertExpectations(t)
}

func TestWeatherService_GetForecastByCityNameAndDate(t *testing.T) {
	mockWeatherRepo := new(mocks.WeatherRepository)
	client := &http.Client{}
	weatherService := NewWeatherServiceImpl(context.Background(), &repository.Repository{
		WeatherRepository: mockWeatherRepo,
	}, client)
	cityName := "city1"
	date := time.Now()
	forecast := model.WeatherForecast{
		CityID: 1, Date: date, Temp: 25.0,
		City: &model.City{Name: cityName, Country: "Country1"},
	}
	mockWeatherRepo.On("GetForecastByCityNameAndDate", cityName, date).Return(forecast, nil)
	result, err := weatherService.GetForecastByCityNameAndDate(cityName, date)
	assert.NoError(t, err)
	assert.Equal(t, forecast, result)
	mockWeatherRepo.AssertExpectations(t)
}
