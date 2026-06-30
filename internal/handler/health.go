package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	AppName string
	Env     string
}

func (h HealthHandler) Health(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":      "UP",
		"application": h.AppName,
		"environment": h.Env,
		"version":     "1.0.0",
	})

}