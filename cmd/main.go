package main

import (
	config "family-tree-api/config"
	router "family-tree-api/router"
)

func main() {
	config.Configuration()
	router.ApiRouter()
}
