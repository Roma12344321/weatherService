package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sync"
	"time"
	"weatherService/pkg/dto"
	"weatherService/pkg/mapper"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

type WeatherServiceImpl struct {
	ctx    context.Context
	repo   *repository.Repository
	client *http.Client
}

func NewWeatherServiceImpl(ctx context.Context, repo *repository.Repository, client *http.Client) *WeatherServiceImpl {
	return &WeatherServiceImpl{ctx: ctx, repo: repo, client: client}
}

type weatherResponse struct {
	List []struct {
		Dt   string `json:"dt_txt" db:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		Sys        struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		Rain struct {
			H float64 `json:"3h"`
		} `json:"rain,omitempty"`
	} `json:"list"`
}

func (s *WeatherServiceImpl) SaveWeatherForeCast(cities []model.City) ([]model.WeatherForecast, error) {
	if err := s.repo.WeatherRepository.DeleteOldDates(); err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	ch := make(chan []model.WeatherForecast)
	ct, cancel := context.WithTimeout(s.ctx, 15*time.Second)
	defer cancel()

	var e error
	wg.Add(len(cities))
	for _, city := range cities {
		go func(city model.City) {
			defer wg.Done()
			w, err := s.saveForecastForCity(ct, city)
			if err != nil {
				e = err
				cancel()
				return
			}
			ch <- w
		}(city)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	res := make([]model.WeatherForecast, 0, len(cities)*40)
	for {
		select {
		case <-ct.Done():
			if e != nil {
				return nil, e
			}
			return nil, ct.Err()
		case w, ok := <-ch:
			if !ok {
				return res, nil
			}
			res = append(res, w...)
		}
	}
}

func (s *WeatherServiceImpl) saveForecastForCity(ctx context.Context, city model.City) ([]model.WeatherForecast, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric",
		city.Lat, city.Lon, viper.GetString("apikey"))
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var weatherResp weatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return nil, err
	}
	result := make([]model.WeatherForecast, 0, 40)
	for _, i := range weatherResp.List {
		if isDateGreaterThanToday(i.Dt) {
			date, err := time.Parse("2006-01-02 15:04:05", i.Dt)
			if err != nil {
				return nil, err
			}
			jsonData, _ := json.Marshal(i)
			w := model.WeatherForecast{
				Date:   date,
				Temp:   i.Main.Temp,
				CityID: city.Id,
				City:   &city,
				Data:   string(jsonData),
			}
			err = s.repo.WeatherRepository.SaveWeatherForeCast(&w)
			if err != nil {
				return nil, err
			}
			result = append(result, w)
		}
	}
	return result, nil
}

func isDateGreaterThanToday(dateStr string) bool {
	date, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		return false
	}
	currentDate := time.Now().Add(-3 * time.Hour).Format("2006-01-02 15:04:05")
	currentParsedDate, err := time.Parse("2006-01-02 15:04:05", currentDate)
	if err != nil {
		log.Printf("Error parsing current date: %v", err)
		return false
	}
	return date.After(currentParsedDate)
}

func (s *WeatherServiceImpl) GetForecastByCityName(city string) (dto.WeatherDto, error) {
	forecasts, err := s.repo.WeatherRepository.GetWeatherForeCastByCityName(city)
	if err != nil {
		return dto.WeatherDto{}, err
	}
	return mapper.MapWeatherForecastListToWeatherDto(forecasts), nil
}
