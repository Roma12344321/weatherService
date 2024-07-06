package mapper

import (
	"sort"
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
	for _, forecast := range list {
		tempSum += forecast.Temp
		res.Dates = append(res.Dates, forecast.Date)
	}
	res.AvgTemp = tempSum / float64(len(list))
	sort.Slice(res.Dates, func(i, j int) bool {
		return res.Dates[i].Before(res.Dates[j])
	})
	return res
}
