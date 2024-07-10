package repository

import (
	"github.com/jmoiron/sqlx"
	"time"
	"weatherService/pkg/model"
)

const differenceBetweenUtcAndMoscowTime = -3

type CityRepository interface {
	SaveCity(city *model.City) error
	GetAllCity() ([]model.City, error)
	GetByName(name string) (model.City, error)
}

type WeatherRepository interface {
	DeleteOldDates() error
	SaveWeatherForeCast(forecast *model.WeatherForecast) error
	GetWeatherForeCastByCityName(city string) ([]model.WeatherForecast, error)
	GetForecastByCityNameAndDate(city string, date time.Time) (model.WeatherForecast, error)
}

type PersonRepository interface {
	CreatePerson(person model.Person) (int, error)
	GetPersonByUsernameAndPassword(username, password string) (model.Person, error)
}

type FavouriteRepository interface {
	AddCityToFavourite(personId, cityId int) error
	GetAllFavouriteCity(personId int) ([]model.City, error)
}

type Repository struct {
	CityRepository
	WeatherRepository
	PersonRepository
	FavouriteRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CityRepository:      NewCityRepositoryImpl(db),
		WeatherRepository:   NewWeatherRepositoryImpl(db),
		PersonRepository:    NewPersonRepositoryImpl(db),
		FavouriteRepository: NewFavouriteRepositoryImpl(db),
	}
}
