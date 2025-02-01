package main

import (
	"log"

	"github.com/Glack134/ScyCast"
	"github.com/Glack134/ScyCast/pkg/handler"
	"github.com/Glack134/ScyCast/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}

	weatherService := service.NewWeatherService(viper.GetString("weather_api_key"))
	weatherHandler := handler.NewWeatherHandler(weatherService)

	mainHandler := new(handler.Handler)
	mainHandler.WeatherHandler = weatherHandler

	srv := new(ScyCast.Server)
	if err := srv.Run(viper.GetString("port"), mainHandler.InitRoutes()); err != nil {
		log.Fatalf("running server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
