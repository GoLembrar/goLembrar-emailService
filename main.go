package main

import (
	"fmt"
	"log"

	"github.com/GoLembrar/goLembrar-emailService/internal/api"
	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/fatih/color"
)

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
