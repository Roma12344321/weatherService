package model

import (
	"encoding/json"
	"time"
)

type WeatherForecast struct {
	Id     int
	Date   time.Time
	Temp   float64
	Data   json.RawMessage
	CityID int
	City   *City
}
