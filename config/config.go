package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string
var DATABASE_HOST string
var DATABASE_PORT string
var DATABASE_USERNAME string
var DATABASE_PASSWORD string
var DATABASE_NAME string

func GetEnvVariable() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
	DATABASE_HOST = os.Getenv("DATABASE_HOST")
	DATABASE_PORT = os.Getenv("DATABASE_PORT")
	DATABASE_USERNAME = os.Getenv("DATABASE_USERNAME")
	DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DATABASE_NAME = os.Getenv("DATABASE_NAME")

}
