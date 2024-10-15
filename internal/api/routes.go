package api

import (
	"time"

	"github.com/GoLembrar/goLembrar-emailService/docs"
	"github.com/GoLembrar/goLembrar-emailService/internal/api/handler"
	"github.com/GoLembrar/goLembrar-emailService/internal/email"
	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() (*gin.Engine, error) {

	r := gin.Default()
	rV1 := r.Group("/v1")
	goEnv := utils.GetEnvVar("GO_ENV")

	if goEnv != "development" {
		docs.SwaggerInfo.Host = "sendemail.golembrar.com"
		r.Use(cors.New(cors.Config{
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
			AllowOrigins: []string{
				"https://api.golembrar.com",
			},
			AllowMethods: []string{
				"GET",
				"POST",
				"OPTIONS",
			},
			AllowHeaders: []string{
				"Authorization",
			},
		}))
	} else {
		docs.SwaggerInfo.Host = "localhost:8080"
		r.Use(cors.Default())
	}

	emailService, err := email.NewEmailService()
	if err != nil {
		return nil, err
	}

	emailHandler := handler.NewEmailHandler(emailService)

	// rV1.POST("/schedule-email", emailHandler.ScheduleEmail)
	rV1.POST("/send-email", emailHandler.SendEmail)
	rV1.GET("/check", handler.HealthChecker)

	r.GET("/", handler.RedirectToDocs)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r, nil
}
