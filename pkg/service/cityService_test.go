package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"sort"
	"testing"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
	"weatherService/pkg/repository/mocks"
	"weatherService/pkg/service"
)

func TestCityService_SaveCities(t *testing.T) {
	mockCityRepo := new(mocks.CityRepository)
	client := &http.Client{}
	cityService := service.NewCityServiceImpl(context.Background(), &repository.Repository{
		CityRepository: mockCityRepo,
	}, client)
	names := []string{"moscow", "barcelona"}
	mockCityRepo.On("SaveCity", mock.AnythingOfType("*model.City")).Return(nil).Times(len(names))
	result, err := cityService.SaveCities(names, "02c02ade5939988a2cffd1fb499fa3fc")
	assert.NoError(t, err)
	assert.Equal(t, len(names), len(result))
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	for i := 0; i < len(names); i++ {
		assert.Equal(t, names[i], result[i].Name)
	}
	mockCityRepo.AssertExpectations(t)
}

func TestCityService_GetAllCity(t *testing.T) {
	mockCityRepo := new(mocks.CityRepository)
	client := &http.Client{}
	cityService := service.NewCityServiceImpl(context.Background(), &repository.Repository{
		CityRepository: mockCityRepo,
	}, client)
	cities := []model.City{
		{Id: 1, Name: "city1", Lat: 10.0, Lon: 20.0},
		{Id: 2, Name: "city2", Lat: 30.0, Lon: 40.0},
	}
	mockCityRepo.On("GetAllCity").Return(cities, nil)
	result, err := cityService.GetAllCity()
	assert.NoError(t, err)
	assert.Equal(t, cities, result)
	mockCityRepo.AssertExpectations(t)
}
