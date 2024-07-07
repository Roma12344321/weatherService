package service

import (
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

type FavouriteServiceImpl struct {
	repo *repository.Repository
}

func NewFavouriteServiceImpl(repo *repository.Repository) *FavouriteServiceImpl {
	return &FavouriteServiceImpl{repo: repo}
}

func (s *FavouriteServiceImpl) AddCityToFavourite(name string, personId int) error {
	city, err := s.repo.CityRepository.GetByName(name)
	if err != nil {
		return err
	}
	err = s.repo.FavouriteRepository.AddCityToFavourite(personId, city.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavouriteServiceImpl) GetAllFavouriteCity(personId int) ([]model.City, error) {
	return s.repo.GetAllFavouriteCity(personId)
}
