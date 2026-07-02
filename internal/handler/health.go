package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	AppName string
	Env     string
}

type ReadyHandler struct {
	DB *gorm.DB
}

func (h HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "UP",
		"application": h.AppName,
		"environment": h.Env,
		"version":     "1.0.0",
	})
}

func (r ReadyHandler) Ready(c *gin.Context) {
	sqlDB, err := r.DB.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unavailable",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unavailable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
