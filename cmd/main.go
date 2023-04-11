package main

import (
	"fmt"
	"os"

	config "github.com/thaynarasilvapinto/family-tree-api/config"
	storage "github.com/thaynarasilvapinto/family-tree-api/internal/adapter/postgres"
	familyRepository "github.com/thaynarasilvapinto/family-tree-api/internal/repository"
	familyService "github.com/thaynarasilvapinto/family-tree-api/internal/service"
	router "github.com/thaynarasilvapinto/family-tree-api/router"
)

func main() {
	config.Configuration()

	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_DB_NAME")

	connect, err := storage.NewPostgresDatabase(host, port, user, password, dbname)
	if err != nil {
		fmt.Println("Ocorreu um erro ao tentar conectar-se ao banco de dados. Por favor, verifique suas configurações de conexão.")
		return
	}
	familyRepository := familyRepository.NewFamilyRepository(connect)
	familyService := familyService.NewFamilyService(familyRepository)

	router.Router(*familyService)
}
