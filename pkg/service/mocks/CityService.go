// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "weatherService/pkg/model"

	mock "github.com/stretchr/testify/mock"
)

// CityService is an autogenerated mock type for the CityService type
type CityService struct {
	mock.Mock
}

// GetAllCity provides a mock function with given fields:
func (_m *CityService) GetAllCity() ([]model.City, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllCity")
	}

	var r0 []model.City
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.City, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.City); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.City)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCities provides a mock function with given fields: names, apikey
func (_m *CityService) SaveCities(names []string, apikey string) ([]model.City, error) {
	ret := _m.Called(names, apikey)

	if len(ret) == 0 {
		panic("no return value specified for SaveCities")
	}

	var r0 []model.City
	var r1 error
	if rf, ok := ret.Get(0).(func([]string, string) ([]model.City, error)); ok {
		return rf(names, apikey)
	}
	if rf, ok := ret.Get(0).(func([]string, string) []model.City); ok {
		r0 = rf(names, apikey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.City)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, string) error); ok {
		r1 = rf(names, apikey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCityService creates a new instance of CityService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCityService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CityService {
	mock := &CityService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
