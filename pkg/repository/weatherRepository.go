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
	query := "DELETE FROM weather_forecast WHERE date<$1"
	if _, err := r.db.Exec(query, time.Now().Add(differenceBetweenUtcAndMoscowTime*time.Hour)); err != nil {
		return err
	}
	return nil
}

func (r *WeatherRepositoryImpl) SaveWeatherForeCast(forecast *model.WeatherForecast) error {
	query := `INSERT INTO weather_forecast(date, temp, data, city_id) VALUES ($1,$2,$3,$4) ON CONFLICT (date,city_id)
    DO UPDATE SET temp=excluded.temp,data=excluded.data,id=weather_forecast.id,city_id=weather_forecast.city_id RETURNING id;`
	row := r.db.QueryRow(query, forecast.Date, forecast.Temp, forecast.Data, forecast.CityID)
	var id int
	if err := row.Scan(&id); err != nil {
		return err
	}
	forecast.Id = id
	return nil
}

func (r *WeatherRepositoryImpl) GetWeatherForeCastByCityName(city string) ([]model.WeatherForecast, error) {
	query := `SELECT weather_forecast.id, weather_forecast.date, weather_forecast.temp, weather_forecast.data, 
    city.id AS "city.id", city.name AS "city.name", city.country AS "city.country", city.lat AS "city.lat", city.lon AS "city.lon" 
	FROM weather_forecast LEFT JOIN city ON weather_forecast.city_id = city.id WHERE city.name=$1 AND weather_forecast.date>$2`
	var res []model.WeatherForecast
	if err := r.db.Select(&res, query, city, time.Now().Add(differenceBetweenUtcAndMoscowTime*time.Hour)); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *WeatherRepositoryImpl) GetForecastByCityNameAndDate(city string, date time.Time) (model.WeatherForecast, error) {
	query := `SELECT weather_forecast.id, weather_forecast.date, weather_forecast.temp, weather_forecast.data, 
    city.id AS "city.id", city.name AS "city.name", city.country AS "city.country", city.lat AS "city.lat", city.lon AS "city.lon" 
	FROM weather_forecast LEFT JOIN city ON weather_forecast.city_id = city.id WHERE city.name=$1 AND weather_forecast.date=$2`
	var res model.WeatherForecast
	if err := r.db.Get(&res, query, city, date); err != nil {
		return model.WeatherForecast{}, err
	}
	return res, nil
}
