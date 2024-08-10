package main

import (
	"auth-service/internal/app"
	"auth-service/pkg/configs"
	"log"
)

func main() {
	conf, err := configs.LoadConfigs()
	if err != nil {
		log.Println("failed to load configs: ", err)
		return
	}

	err = app.Start(conf)
	if err != nil {
		log.Println("failed to start server: ", err)
		return
	}
}
