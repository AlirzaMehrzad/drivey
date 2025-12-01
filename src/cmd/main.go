package main

import (
	"github.com/alirzamehrzad/drivey/api"
	"github.com/alirzamehrzad/drivey/config"
	"github.com/alirzamehrzad/drivey/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	api.InitServer(cfg)
}
