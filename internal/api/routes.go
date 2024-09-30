package api

import (
	"github.com/GoLembrar/goLembrar-emailService/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.HealthChecker)

	return r
}
