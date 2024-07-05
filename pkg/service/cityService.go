package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
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

func (s *CityServiceImpl) SaveCities(names []string, url string) ([]model.City, error) {
	cities := make([]model.City, 0, len(names))
	var errs error
	var wg sync.WaitGroup
	ch := make(chan model.City)
	wg.Add(len(names))
	for _, city := range names {
		go func(city string) {
			defer wg.Done()
			res, err := s.saveOneCity(city, url)
			if err != nil {
				log.Println(err.Error() + " for " + city)
				errs = err
				return
			}
			ch <- res
		}(city)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for c := range ch {
		cities = append(cities, c)
	}
	if errs != nil {
		return nil, errs
	}
	return cities, nil

}

func (s *CityServiceImpl) saveOneCity(city string, url string) (model.City, error) {
	ct, cancel := context.WithTimeout(s.ctx, 5*time.Second)
	defer cancel()
	var res []model.City
	req, err := http.NewRequestWithContext(ct, "GET", url+city, nil)
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
	err = s.repo.SaveCity(&res[0])
	if err != nil {
		return model.City{}, err
	}
	return res[0], nil
}
