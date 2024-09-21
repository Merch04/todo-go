package main

import (
	"github.com/spf13/viper"
	"log"
	"todo/config"
	"todo/server"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err)
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
