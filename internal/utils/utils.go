package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln(err)
	}
	return os.Getenv(key)
}
