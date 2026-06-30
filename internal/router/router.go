package router

import (
	"net/http"

	"github.com/afiffaizun/mining-asset-management/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(health handler.HealthHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app_name": health.AppName,
			"env":      health.Env,
			"status":   "ok",
		})
	})

	return r
}
