package model

import "time"

type WeatherForecast struct {
	Id     int
	Date   time.Time
	Temp   float64
	Data   string
	CityID int
	City   *City
}
