package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SECRET_KEY")
}
