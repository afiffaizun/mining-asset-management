package app

import (
	"fmt"
	"log/slog"

	"github.com/afiffaizun/mining-asset-management/internal/config"
	"github.com/afiffaizun/mining-asset-management/internal/database"
	"github.com/afiffaizun/mining-asset-management/internal/handler"
	"github.com/afiffaizun/mining-asset-management/internal/logger"
	"github.com/afiffaizun/mining-asset-management/internal/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
	Router *gin.Engine
	DB     *gorm.DB
}

func New() (*App, error) {
	cfg := config.Load()
	logg := logger.New()

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	health := handler.HealthHandler{
		AppName: cfg.AppName,
		Env:     cfg.AppEnv,
	}

	ready := handler.ReadyHandler{
		DB: db,
	}

	r := router.SetupRouter(health, ready)

	return &App{
		Config: cfg,
		Logger: logg,
		Router: r,
		DB:     db,
	}, nil
}

func (a *App) Run() error {
	a.Logger.Info("Starting server",
		"port", a.Config.AppPort,
	)
	return a.Router.Run(":" + a.Config.AppPort)
}
