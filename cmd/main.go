package main

import (
	"dashboard/internal/config"
	"dashboard/internal/dashboard"
	"dashboard/internal/router"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

func main() {
	cfg := config.Config{}

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatal("cannot read config: " + err.Error())
	}

	dashboard, err := dashboard.Init(cfg.DB)
	if err != nil {
		log.Fatal("Couldn't init dashboard service", err)
	}

	router := router.SetupRouter(dashboard)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
