package main

import (
	"github.com/alirzamehrzad/drivey/api"
	"github.com/alirzamehrzad/drivey/config"
	"github.com/alirzamehrzad/drivey/data/cache"
	"github.com/alirzamehrzad/drivey/data/db"
	"github.com/alirzamehrzad/drivey/data/db/migrations"
	"github.com/alirzamehrzad/drivey/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migrations.Up_1()

	api.InitServer(cfg)
}
