package model

type City struct {
	Id      int
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}
