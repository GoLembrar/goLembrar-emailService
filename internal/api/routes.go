package api

import (
	"github.com/GoLembrar/goLembrar-emailService/internal/api/handler"
	"github.com/GoLembrar/goLembrar-emailService/internal/api/middlewares"
	"github.com/GoLembrar/goLembrar-emailService/internal/email"
	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	goEnv := utils.GetEnvVar("GO_ENV")

	if goEnv != "development" {
		r.Use(middlewares.CorsMiddleware())
	}

	emailService, err := email.NewEmailService()
	if err != nil {
		return nil, err
	}

	emailHandler := handler.NewEmailHandler(emailService)

	r.POST("/send-email", emailHandler.SendEmail)
	r.POST("/schedule-email", emailHandler.ScheduleEmail)
	r.GET("/check", handler.HealthChecker)
	return r, nil
}
