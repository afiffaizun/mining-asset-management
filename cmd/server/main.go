package main

import (
	"log"

	"github.com/afiffaizun/mining-asset-management/internal/config"
	"github.com/afiffaizun/mining-asset-management/internal/handler"
	"github.com/afiffaizun/mining-asset-management/internal/logger"
	"github.com/afiffaizun/mining-asset-management/internal/router"
)

func main() {

	cfg := config.Load()

	logg := logger.New()

	health := handler.HealthHandler{
		AppName: cfg.AppName,
		Env:     cfg.AppEnv,
	}

	r := router.SetupRouter(health)

	logg.Info("Starting server",
		"port", cfg.AppPort,
	)

	err := r.Run(":" + cfg.AppPort)

	if err != nil {
		log.Fatal(err)
	}

}