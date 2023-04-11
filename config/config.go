package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Configuration() {

	err := godotenv.Load("family-tree.env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
