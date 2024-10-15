package main

import (
	"fmt"
	"log"

	_ "github.com/GoLembrar/goLembrar-emailService/docs"

	"github.com/GoLembrar/goLembrar-emailService/internal/api"
	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/fatih/color"
)

// @title			goLembrar email microservice
// @version		0.1
// @description	Dedicated microservice for send the emails for https://golembrar.com
// @BasePath		/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	port := utils.GetEnvVar("PORT")

	r, err := api.SetupRoutes()
	if err != nil {
		log.Fatalf("error on setup routes: %v\n", err)
	}

	color.Cyan("goLembrar EmailService at port :%s\n", port)
	err = r.Run(":" + fmt.Sprint(port))
	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
