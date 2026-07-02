package router

import (
	"github.com/afiffaizun/mining-asset-management/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(health handler.HealthHandler, ready handler.ReadyHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", health.Health)
	r.GET("/ready", ready.Ready)

	return r
}
