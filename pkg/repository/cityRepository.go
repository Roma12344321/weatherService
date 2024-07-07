package repository

import (
	"github.com/jmoiron/sqlx"
	"weatherService/pkg/model"
)

type CityRepositoryImpl struct {
	db *sqlx.DB
}

func NewCityRepositoryImpl(db *sqlx.DB) *CityRepositoryImpl {
	return &CityRepositoryImpl{db: db}
}

func (r *CityRepositoryImpl) SaveCity(city *model.City) error {
	query := `INSERT INTO city(name, country, lat, lon) VALUES ($1, $2, $3, $4) ON CONFLICT (name) DO UPDATE 
    SET country = excluded.country, lat = excluded.lat, lon = excluded.lon, id = city.id RETURNING id`
	row := r.db.QueryRow(query, city.Name, city.Country, city.Lat, city.Lon)
	var id int
	if err := row.Scan(&id); err != nil {
		return err
	}
	city.Id = id
	return nil
}

func (r *CityRepositoryImpl) GetAllCity() ([]model.City, error) {
	query := "SELECT * FROM city ORDER BY name"
	var res []model.City
	if err := r.db.Select(&res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CityRepositoryImpl) GetByName(name string) (model.City, error) {
	query := "SELECT * FROM city where name=$1"
	var res model.City
	err := r.db.Get(&res, query, name)
	return res, err
}
