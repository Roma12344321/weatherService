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

func TestFavouriteRepository_AddCityToFavourite(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewFavouriteRepositoryImpl(sqlxDB)
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO favourite(person_id, city_id) VALUES ($1,$2) ON CONFLICT (person_id,city_id) DO NOTHING`)).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddCityToFavourite(1, 1)
	assert.NoError(t, err, "AddCityToFavourite should not return an error")
}

func TestFavouriteRepository_GetAllFavouriteCity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewFavouriteRepositoryImpl(sqlxDB)
	cities := []model.City{
		{Id: 1, Name: "Test City 1", Country: "Test Country 1", Lat: 0.0, Lon: 0.0},
		{Id: 2, Name: "Test City 2", Country: "Test Country 2", Lat: 1.0, Lon: 1.0},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "country", "lat", "lon"})
	for _, city := range cities {
		rows.AddRow(city.Id, city.Name, city.Country, city.Lat, city.Lon)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT city.id, name, country, lat, lon FROM city JOIN favourite ON city.id = favourite.city_id WHERE person_id=$1 ORDER BY name`)).
		WithArgs(1).
		WillReturnRows(rows)
	result, err := repo.GetAllFavouriteCity(1)
	assert.NoError(t, err, "GetAllFavouriteCity should not return an error")
	assert.Equal(t, cities, result, "Expected cities to match")
}
