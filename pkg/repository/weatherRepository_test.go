package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

func TestWeatherRepository_SaveWeatherForeCast(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewWeatherRepositoryImpl(sqlxDB)

	forecast := &model.WeatherForecast{
		Date:   time.Now(),
		Temp:   25.0,
		Data:   []byte("Sunny"),
		CityID: 1,
	}
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO weather_forecast(date, temp, data, city_id) VALUES ($1,$2,$3,$4) ON CONFLICT (date,city_id)
    DO UPDATE SET temp=excluded.temp,data=excluded.data,id=weather_forecast.id,city_id=weather_forecast.city_id RETURNING id;`)).
		WithArgs(forecast.Date, forecast.Temp, forecast.Data, forecast.CityID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	err = repo.SaveWeatherForeCast(forecast)
	assert.NoError(t, err, "SaveWeatherForeCast should not return an error")
	assert.Equal(t, 1, forecast.Id)
}

func TestWeatherRepository_GetWeatherForeCastByCityName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewWeatherRepositoryImpl(sqlxDB)
	city := "Test City"
	forecast := model.WeatherForecast{
		Id:   1,
		Date: time.Now(),
		Temp: 25.0,
		Data: []byte("Sunny"),
		City: &model.City{
			Id:      1,
			Name:    city,
			Country: "Test Country",
			Lat:     0.0,
			Lon:     0.0,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "date", "temp", "data", "city.id", "city.name", "city.country", "city.lat", "city.lon"}).
		AddRow(forecast.Id, forecast.Date, forecast.Temp, forecast.Data, forecast.City.Id, forecast.City.Name, forecast.City.Country, forecast.City.Lat, forecast.City.Lon)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT weather_forecast.id, weather_forecast.date, weather_forecast.temp, weather_forecast.data,
		city.id AS "city.id", city.name AS "city.name", city.country AS "city.country", city.lat AS "city.lat", city.lon AS "city.lon"
	FROM weather_forecast LEFT JOIN city ON weather_forecast.city_id = city.id WHERE city.name=$1 AND weather_forecast.date>$2`)).
		WithArgs(city, sqlmock.AnyArg()).
		WillReturnRows(rows)
	result, err := repo.GetWeatherForeCastByCityName(city)
	assert.NoError(t, err, "GetWeatherForeCastByCityName should not return an error")
	assert.NotEmpty(t, result, "Expected forecasts list to be not empty")
	assert.Equal(t, forecast, result[0], "Expected forecast to match")
}

func TestWeatherRepository_GetForecastByCityNameAndDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewWeatherRepositoryImpl(sqlxDB)

	city := "Test City"
	date := time.Now()
	forecast := model.WeatherForecast{
		Id:   1,
		Date: date,
		Temp: 25.0,
		Data: []byte("Sunny"),
		City: &model.City{
			Id:      1,
			Name:    city,
			Country: "Test Country",
			Lat:     0.0,
			Lon:     0.0,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "date", "temp", "data", "city.id", "city.name", "city.country", "city.lat", "city.lon"}).
		AddRow(forecast.Id, forecast.Date, forecast.Temp, forecast.Data, forecast.City.Id, forecast.City.Name, forecast.City.Country, forecast.City.Lat, forecast.City.Lon)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT weather_forecast.id, weather_forecast.date, weather_forecast.temp, weather_forecast.data, 
    city.id AS "city.id", city.name AS "city.name", city.country AS "city.country", city.lat AS "city.lat", city.lon AS "city.lon" 
	FROM weather_forecast LEFT JOIN city ON weather_forecast.city_id = city.id WHERE city.name=$1 AND weather_forecast.date=$2`)).
		WithArgs(city, date).
		WillReturnRows(rows)
	result, err := repo.GetForecastByCityNameAndDate(city, date)
	assert.NoError(t, err, "GetForecastByCityNameAndDate should not return an error")
	assert.Equal(t, forecast, result, "Expected forecast to match")
}
