package main

import (
	config "github.com/thaynarasilvapinto/family-tree-api/config"
	router "github.com/thaynarasilvapinto/family-tree-api/router"
)

func main() {

	// host := os.Getenv("DATABASE_HOST")
	// port := os.Getenv("DATABASE_PORT")
	// user := os.Getenv("DATABASE_USER")
	// password := os.Getenv("DATABASE_PASSWORD")
	// dbname := os.Getenv("DATABASE_DB_NAME")

	config.Configuration()
	router.Router()
}
