package main

import (
	"github.com/spf13/viper"
	"github.com/khuchuz/go-clean-architecture/config"
	"github.com/khuchuz/go-clean-architecture/server"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
