package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Configuration() {

	err := godotenv.Load("app.env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
