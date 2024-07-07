package service

import (
	"context"
	"net/http"
	"time"
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
	GetForecastByCityNameAndDate(city string, date time.Time) (model.WeatherForecast, error)
}

type AuthService interface {
	Registration(person model.Person) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type FavouriteService interface {
	AddCityToFavourite(name string, personId int) error
	GetAllFavouriteCity(personId int) ([]model.City, error)
}

type Service struct {
	CityService
	WeatherService
	AuthService
	FavouriteService
}

func NewService(ctx context.Context, repo *repository.Repository, client *http.Client) *Service {
	return &Service{
		CityService:      NewCityServiceImpl(ctx, repo, client),
		WeatherService:   NewWeatherServiceImpl(ctx, repo, client),
		AuthService:      NewAuthServiceImpl(repo),
		FavouriteService: NewFavouriteServiceImpl(repo),
	}
}
