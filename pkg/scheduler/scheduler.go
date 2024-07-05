package scheduler

import (
	"context"
	"log"
	"time"
	"weatherService/pkg/model"
	"weatherService/pkg/service"
)

type Scheduler struct {
	ctx     context.Context
	service *service.Service
}

func NewScheduler(ctx context.Context, service *service.Service) *Scheduler {
	return &Scheduler{ctx: ctx, service: service}
}

func (s *Scheduler) Schedule(cities []model.City) {
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-time.After(5 * time.Minute):
				_, err := s.service.WeatherService.SaveWeatherForeCast(cities)
				if err != nil {
					log.Fatalln(err.Error())
				}
			}
		}
	}()
}
