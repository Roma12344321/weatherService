package repository

import (
	"github.com/jmoiron/sqlx"
	"weatherService/pkg/model"
)

type CityRepository interface {
	SaveCity(city *model.City) error
	GetAllCity() ([]model.City, error)
}

type WeatherRepository interface {
	DeleteOldDates() error
	SaveWeatherForeCast(forecast *model.WeatherForecast) error
	GetWeatherForeCastByCityName(city string) ([]model.WeatherForecast, error)
}

type Repository struct {
	CityRepository
	WeatherRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{CityRepository: NewCityRepositoryImpl(db), WeatherRepository: NewWeatherRepositoryImpl(db)}
}
