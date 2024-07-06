package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	"weatherService/pkg/model"
	"weatherService/pkg/repository"
)

type CityServiceImpl struct {
	ctx    context.Context
	repo   *repository.Repository
	client *http.Client
}

func NewCityServiceImpl(ctx context.Context, repo *repository.Repository, client *http.Client) *CityServiceImpl {
	return &CityServiceImpl{ctx: ctx, repo: repo, client: client}
}

func (s *CityServiceImpl) SaveCities(names []string) ([]model.City, error) {
	url := "http://api.openweathermap.org/geo/1.0/direct?limit=1&appid=" + viper.GetString("apikey") + "&q="
	cities := make([]model.City, 0, len(names))
	var errs error
	var wg sync.WaitGroup
	ch := make(chan model.City)
	ct, cancel := context.WithTimeout(s.ctx, 10*time.Second)
	defer cancel()
	wg.Add(len(names))
	for _, city := range names {
		go func(city string) {
			defer wg.Done()
			res, err := s.saveOneCity(s.ctx, city, url)
			if err != nil {
				log.Println(err.Error() + " for " + city)
				errs = err
				cancel()
				return
			}
			ch <- res
		}(city)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for {
		select {
		case <-ct.Done():
			if errs != nil {
				return nil, errs
			}
			return nil, ct.Err()
		case c, ok := <-ch:
			if !ok {
				return cities, nil
			}
			cities = append(cities, c)
		}
	}
}

func (s *CityServiceImpl) saveOneCity(ctx context.Context, city string, url string) (model.City, error) {
	var res []model.City
	req, err := http.NewRequestWithContext(ctx, "GET", url+city, nil)
	if err != nil {
		return model.City{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return model.City{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return model.City{}, err
	}
	if len(res) == 0 {
		return model.City{}, errors.New("empty response")
	}
	res[0].Name = strings.ToLower(res[0].Name)
	err = s.repo.SaveCity(&res[0])
	if err != nil {
		return model.City{}, err
	}
	return res[0], nil
}

func (s *CityServiceImpl) GetAllCity() ([]model.City, error) {
	return s.repo.CityRepository.GetAllCity()
}
