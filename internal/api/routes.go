package api

import (
	"github.com/GoLembrar/goLembrar-emailService/internal/api/handler"
	"github.com/GoLembrar/goLembrar-emailService/internal/email"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() (*gin.Engine, error) {
	r := gin.Default()

	emailService, err := email.NewEmailService()
	if err != nil {
		return nil, err
	}

	emailHandler := handler.NewEmailHandler(emailService)

	r.POST("/send-email", emailHandler.SendEmail)
	r.GET("/check", handler.HealthChecker)

	return r, nil
}
