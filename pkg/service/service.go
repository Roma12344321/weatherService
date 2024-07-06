package service

import (
	"context"
	"net/http"
	"weatherService/pkg/dto"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

type CityService interface {
	SaveCities(names []string) ([]model.City, error)
	GetAllCity() ([]model.City, error)
}

type WeatherService interface {
	SaveWeatherForeCast([]model.City) ([]model.WeatherForecast, error)
	GetForecastByCityName(city string) (dto.WeatherDto, error)
}

type Service struct {
	CityService
	WeatherService
}

func NewService(ctx context.Context, repo *repository.Repository, client *http.Client) *Service {
	return &Service{
		CityService:    NewCityServiceImpl(ctx, repo, client),
		WeatherService: NewWeatherServiceImpl(ctx, repo, client),
	}
}
