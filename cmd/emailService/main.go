package main

import (
	"log"

	"github.com/GoLembrar/goLembrar-emailService/internal/api"
)

func main() {
	const port = "9797"

	r, _ := api.SetupRoutes()

	log.Printf("goLembrar EmailService at port :%s\n", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
