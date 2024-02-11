package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetToken(env string) string {
	err := godotenv.Load()
	if err != nil {
		log.Printf("There was no .env file found to load the token from")
	}	

	token := os.Getenv(env)

	if token == "" {
		log.Panic("You must provide a valid token")
	}

	return token
}