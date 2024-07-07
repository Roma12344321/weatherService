package scheduler

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"time"
	"weatherService/pkg/service"
)

type Scheduler struct {
	ctx     context.Context
	service *service.Service
}

func NewScheduler(ctx context.Context, service *service.Service) *Scheduler {
	return &Scheduler{ctx: ctx, service: service}
}

func (s *Scheduler) Schedule() {
	go func() {
		apikey := viper.GetString("apikey")
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-ticker.C:
				cities, err := s.service.CityService.GetAllCity()
				if err != nil {
					log.Println(err.Error())
					continue
				}
				_, err = s.service.WeatherService.SaveWeatherForeCast(cities, apikey)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}()
}
