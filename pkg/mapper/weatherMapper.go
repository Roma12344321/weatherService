package mapper

import (
	"sort"
	"time"
	"weatherService/pkg/dto"
	"weatherService/pkg/model"
)

func MapWeatherForecastListToWeatherDto(list []model.WeatherForecast) dto.WeatherDto {
	var res dto.WeatherDto
	if len(list) < 1 {
		return res
	}
	res.Name = list[0].City.Name
	res.Country = list[0].City.Country
	var tempSum float64
	sort.Slice(list, func(i, j int) bool {
		return list[i].Date.Before(list[j].Date)
	})
	for _, forecast := range list {
		tempSum += forecast.Temp
		res.Dates = append(res.Dates, forecast.Date.Format(time.DateTime))
	}
	res.AvgTemp = tempSum / float64(len(list))
	return res
}
