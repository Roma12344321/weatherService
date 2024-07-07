package repository

import (
	"github.com/jmoiron/sqlx"
	"weatherService/pkg/model"
)

type FavouriteRepositoryImpl struct {
	db *sqlx.DB
}

func NewFavouriteRepositoryImpl(db *sqlx.DB) *FavouriteRepositoryImpl {
	return &FavouriteRepositoryImpl{db: db}
}

func (r *FavouriteRepositoryImpl) AddCityToFavourite(personId, cityId int) error {
	query := "INSERT INTO favourite(person_id, city_id) VALUES ($1,$2) ON CONFLICT (person_id,city_id) DO NOTHING"
	_, err := r.db.Exec(query, personId, cityId)
	return err
}

func (r *FavouriteRepositoryImpl) GetAllFavouriteCity(personId int) ([]model.City, error) {
	query := "SELECT city.id, name, country, lat, lon FROM city JOIN favourite ON city.id = favourite.city_id WHERE person_id=$1 ORDER BY name"
	var res []model.City
	err := r.db.Select(&res, query, personId)
	if err != nil {
		return nil, err
	}
	return res, err
}
