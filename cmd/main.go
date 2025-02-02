package main

import (
	"log"

	skycast "github.com/Glack134/SkyCast"
	"github.com/spf13/viper"
)

func main() {
	srv := new(skycast.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
