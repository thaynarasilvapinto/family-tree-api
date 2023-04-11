package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	api "github.com/thaynarasilvapinto/family-tree-api/api"
	familyService "github.com/thaynarasilvapinto/family-tree-api/internal/service"
)

func Router(familyService familyService.FamilyService) {

	router := mux.NewRouter()
	handler := &api.Handler{FamilyService: familyService}
	router.HandleFunc("/family/history/{id}", handler.GetFamilyTree).Methods("GET")

	fmt.Println("Server is up and running on localhost:8080")
	http.ListenAndServe(":8080", router)
}
