package main

import (
	"log"

	"github.com/GoLembrar/goLembrar-emailService/internal/api"
)

func main() {

	r := api.SetupRoutes()

	log.Println("goLembrar EmailService at port :8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
