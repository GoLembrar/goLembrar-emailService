package main

import (
	"fmt"
	"log"

	"github.com/GoLembrar/goLembrar-emailService/internal/api"
)

func main() {
	const port = 9797

	r, err := api.SetupRoutes()
	if err != nil {
		log.Fatalf("error on setup routes: %v\n", err)
	}

	log.Printf("goLembrar EmailService at port :%d\n", port)
	err = r.Run(":" + fmt.Sprint(port))
	if err != nil {
		log.Fatalf("error starting server: %v\n", err)
	}
}
