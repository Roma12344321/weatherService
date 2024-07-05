package repository

import (
	"github.com/jmoiron/sqlx"
	"time"
	"weatherService/pkg/model"
)

type WeatherRepositoryImpl struct {
	db *sqlx.DB
}

func NewWeatherRepositoryImpl(db *sqlx.DB) *WeatherRepositoryImpl {
	return &WeatherRepositoryImpl{db: db}
}

func (r *WeatherRepositoryImpl) DeleteOldDates() error {
	query := "DELETE FROM weather_forecast where date<$1"
	if _, err := r.db.Exec(query, time.Now()); err != nil {
		return err
	}
	return nil
}

func (r *WeatherRepositoryImpl) SaveWeatherForeCast(forecast *model.WeatherForecast) error {
	query := `INSERT INTO weather_forecast(date, temp, data, city_id) VALUES ($1,$2,$3,$4) ON CONFLICT (date,city_id)
    DO UPDATE SET date=excluded.date,temp=excluded.temp,data=excluded.data,id=weather_forecast.id RETURNING id`
	row := r.db.QueryRow(query, forecast.Date, forecast.Temp, forecast.Data, forecast.CityID)
	var id int
	if err := row.Scan(&id); err != nil {
		return err
	}
	forecast.Id = id
	return nil
}
