package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

func TestCityRepository_SaveCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewCityRepositoryImpl(sqlxDB)
	city := &model.City{
		Id:      1,
		Name:    "Test City",
		Country: "Test Country",
		Lat:     0.0,
		Lon:     0.0,
	}
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO city(name, country, lat, lon) VALUES ($1, $2, $3, $4) ON CONFLICT (name) DO UPDATE 
    SET country = excluded.country, lat = excluded.lat, lon = excluded.lon, id = city.id RETURNING id`)).
		WithArgs(city.Name, city.Country, city.Lat, city.Lon).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	err = repo.SaveCity(city)
	assert.NoError(t, err, "SaveCity should not return an error")
	assert.Equal(t, 1, city.Id)
}

func TestCityRepository_GetAllCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewCityRepositoryImpl(sqlxDB)
	cities := []model.City{
		{Id: 1, Name: "Test City 1", Country: "Test Country 1", Lat: 0.0, Lon: 0.0},
		{Id: 2, Name: "Test City 2", Country: "Test Country 2", Lat: 1.0, Lon: 1.0},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "country", "lat", "lon"})
	for _, city := range cities {
		rows.AddRow(city.Id, city.Name, city.Country, city.Lat, city.Lon)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM city ORDER BY name`)).WillReturnRows(rows)
	result, err := repo.GetAllCity()
	assert.NoError(t, err, "GetAllCity should not return an error")
	assert.Equal(t, cities, result, "Expected cities to match")
}

func TestCityRepository_GetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewCityRepositoryImpl(sqlxDB)
	city := model.City{
		Id:      1,
		Name:    "Test City",
		Country: "Test Country",
		Lat:     0.0,
		Lon:     0.0,
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM city WHERE name=$1`)).
		WithArgs(city.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "country", "lat", "lon"}).
			AddRow(city.Id, city.Name, city.Country, city.Lat, city.Lon))
	savedCity, err := repo.GetByName(city.Name)
	assert.NoError(t, err, "GetByName should not return an error")
	assert.Equal(t, city, savedCity, "Expected cities to match")
}
