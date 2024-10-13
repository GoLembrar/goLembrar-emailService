package api

import (
	"time"

	"github.com/GoLembrar/goLembrar-emailService/internal/api/handler"
	"github.com/GoLembrar/goLembrar-emailService/internal/email"
	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	goEnv := utils.GetEnvVar("GO_ENV")

	if goEnv == "development" {
		r.Use(cors.Default())
	} else {
		r.Use(cors.New(cors.Config{
			AllowOrigins:  []string{"https://api.golembrar.com"},
			AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders: []string{"Content-Length"},
			MaxAge:        12 * time.Hour,
		}))
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
