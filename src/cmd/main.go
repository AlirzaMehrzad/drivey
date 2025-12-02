package main

import (
	"log"

	"github.com/alirzamehrzad/drivey/api"
	"github.com/alirzamehrzad/drivey/config"
	"github.com/alirzamehrzad/drivey/data/cache"
	"github.com/alirzamehrzad/drivey/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
