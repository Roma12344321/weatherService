package dto

import "time"

type WeatherDto struct {
	Name    string      `json:"name"`
	Country string      `json:"country"`
	AvgTemp float64     `json:"avg_temp"`
	Dates   []time.Time `json:"dates"`
}
