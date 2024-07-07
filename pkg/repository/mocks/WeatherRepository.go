// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "weatherService/pkg/model"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// WeatherRepository is an autogenerated mock type for the WeatherRepository type
type WeatherRepository struct {
	mock.Mock
}

// DeleteOldDates provides a mock function with given fields:
func (_m *WeatherRepository) DeleteOldDates() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeleteOldDates")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetForecastByCityNameAndDate provides a mock function with given fields: city, date
func (_m *WeatherRepository) GetForecastByCityNameAndDate(city string, date time.Time) (model.WeatherForecast, error) {
	ret := _m.Called(city, date)

	if len(ret) == 0 {
		panic("no return value specified for GetForecastByCityNameAndDate")
	}

	var r0 model.WeatherForecast
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Time) (model.WeatherForecast, error)); ok {
		return rf(city, date)
	}
	if rf, ok := ret.Get(0).(func(string, time.Time) model.WeatherForecast); ok {
		r0 = rf(city, date)
	} else {
		r0 = ret.Get(0).(model.WeatherForecast)
	}

	if rf, ok := ret.Get(1).(func(string, time.Time) error); ok {
		r1 = rf(city, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWeatherForeCastByCityName provides a mock function with given fields: city
func (_m *WeatherRepository) GetWeatherForeCastByCityName(city string) ([]model.WeatherForecast, error) {
	ret := _m.Called(city)

	if len(ret) == 0 {
		panic("no return value specified for GetWeatherForeCastByCityName")
	}

	var r0 []model.WeatherForecast
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]model.WeatherForecast, error)); ok {
		return rf(city)
	}
	if rf, ok := ret.Get(0).(func(string) []model.WeatherForecast); ok {
		r0 = rf(city)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.WeatherForecast)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveWeatherForeCast provides a mock function with given fields: forecast
func (_m *WeatherRepository) SaveWeatherForeCast(forecast *model.WeatherForecast) error {
	ret := _m.Called(forecast)

	if len(ret) == 0 {
		panic("no return value specified for SaveWeatherForeCast")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.WeatherForecast) error); ok {
		r0 = rf(forecast)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWeatherRepository creates a new instance of WeatherRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWeatherRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WeatherRepository {
	mock := &WeatherRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
