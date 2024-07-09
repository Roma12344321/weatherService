package main

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"weatherService"
	_ "weatherService/docs"
	"weatherService/pkg/handler"
	"weatherService/pkg/repository"
	"weatherService/pkg/scheduler"
	"weatherService/pkg/service"
)

var cityArr = []string{"moscow", "london", "warsaw", "berlin", "madrid", "barcelona", "sidney", "canberra", "toronto", "minsk", "kiev", "amsterdam", "paris", "leon", "rome", "milan", "ankara", "bangkok", "bern", "bratislava"}

// @title Weather Service API
// @version 1.0
// @description This is a weather service API documentation.
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chanOs := make(chan os.Signal)
	signal.Notify(chanOs, syscall.SIGINT, syscall.SIGTERM)

	initConfig()
	db := initDb()
	client := &http.Client{}
	repositories := repository.NewRepository(db)
	services := service.NewService(ctx, repositories, client)
	schedulers := scheduler.NewScheduler(ctx, services)

	cities, err := services.CityService.SaveCities(cityArr, viper.GetString("apikey"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = services.WeatherService.SaveWeatherForeCast(cities, viper.GetString("apikey"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	schedulers.Schedule()

	handlers := handler.NewHandler(services)
	server := new(weatherService.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to run server: %s", err)
		}
	}()

	<-chanOs

	err = server.ShutDown(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	cancel()
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
