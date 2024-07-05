package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"weatherService/pkg/repository"
	"weatherService/pkg/scheduler"
	"weatherService/pkg/service"
)

var cityArr = []string{"moscow", "london", "warsaw", "berlin", "madrid", "barcelona", "sidney", "canberra"}

func main() {
	ctx := context.Background()
	initConfig()
	db := initDb()
	client := &http.Client{}
	repositories := repository.NewRepository(db)
	services := service.NewService(ctx, repositories, client)
	schedulers := scheduler.NewScheduler(ctx, services)
	cityUrl := "http://api.openweathermap.org/geo/1.0/direct?limit=1&appid=" + viper.GetString("apikey") + "&q="
	cities, err := services.CityService.SaveCities(cityArr, cityUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(cities)
	_, err = services.WeatherService.SaveWeatherForeCast(cities)
	if err != nil {
		log.Fatalln(err.Error())
	}
	schedulers.Schedule(cities)
}

func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

func initDb() *sqlx.DB {
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err)
	}
	return db
}
