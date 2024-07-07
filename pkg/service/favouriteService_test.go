package service_test

import (
	"testing"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
	"weatherService/pkg/repository/mocks"
	"weatherService/pkg/service"

	"github.com/stretchr/testify/assert"
)

func TestFavouriteService_AddCityToFavourite(t *testing.T) {
	mockCityRepo := new(mocks.CityRepository)
	mockFavRepo := new(mocks.FavouriteRepository)
	favouriteService := service.NewFavouriteServiceImpl(&repository.Repository{
		CityRepository:      mockCityRepo,
		FavouriteRepository: mockFavRepo,
	})
	personId := 1
	cityName := "city1"
	city := model.City{Id: 1, Name: cityName}
	mockCityRepo.On("GetByName", cityName).Return(city, nil)
	mockFavRepo.On("AddCityToFavourite", personId, city.Id).Return(nil)
	err := favouriteService.AddCityToFavourite(cityName, personId)
	assert.NoError(t, err)
	mockCityRepo.AssertExpectations(t)
	mockFavRepo.AssertExpectations(t)
}

func TestFavouriteService_GetAllFavouriteCity(t *testing.T) {
	mockFavRepo := new(mocks.FavouriteRepository)
	favouriteService := service.NewFavouriteServiceImpl(&repository.Repository{
		FavouriteRepository: mockFavRepo,
	})
	personId := 1
	cities := []model.City{
		{Id: 1, Name: "city1"},
		{Id: 2, Name: "city2"},
	}
	mockFavRepo.On("GetAllFavouriteCity", personId).Return(cities, nil)
	result, err := favouriteService.GetAllFavouriteCity(personId)
	assert.NoError(t, err)
	assert.Equal(t, cities, result)
	mockFavRepo.AssertExpectations(t)
}
