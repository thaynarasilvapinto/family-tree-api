package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"

	_ "github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	api "github.com/thaynarasilvapinto/family-tree-api/api"
	storage "github.com/thaynarasilvapinto/family-tree-api/internal/adapter/postgres"
	familyRepository "github.com/thaynarasilvapinto/family-tree-api/internal/repository"
	familyService "github.com/thaynarasilvapinto/family-tree-api/internal/service"
)

func setupTestEnv(t *testing.T) {
	t.Cleanup(teardownTestEnv)
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5433")
	os.Setenv("DATABASE_USER", "test")
	os.Setenv("DATABASE_PASSWORD", "test1234")
	os.Setenv("DATABASE_DB_NAME", "family_tree_test")
}

func teardownTestEnv() {
	os.Unsetenv("DATABASE_HOST")
	os.Unsetenv("DATABASE_PORT")
	os.Unsetenv("DATABASE_USER")
	os.Unsetenv("DATABASE_PASSWORD")
	os.Unsetenv("DATABASE_DB_NAME")
}

func setupDbTest(t *testing.T) (*familyService.FamilyService, *storage.PostgresDatabase) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_DB_NAME")

	connect, err := storage.NewPostgresDatabase(host, port, user, password, dbname)
	if err != nil {
		t.Fatalf("Ocorreu um erro ao tentar conectar-se ao banco de dados. Por favor, verifique suas configurações de conexão: %s", err)
	}
	familyRepository := familyRepository.NewFamilyRepository(connect)
	familyService := familyService.NewFamilyService(familyRepository)

	return familyService, connect
}

func afterDbTest(db *storage.PostgresDatabase) {
	db.Close()
}

func setupGetFamilyTree(id string, t *testing.T) *http.Response {
	setupTestEnv(t)
	familyService, connect := setupDbTest(t)
	handler := &api.Handler{FamilyService: *familyService}

	router := mux.NewRouter()
	router.HandleFunc("/family/history/{id}", handler.GetFamilyTree).Methods("GET")

	req := httptest.NewRequest(http.MethodGet, "/family/history/"+id, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	afterDbTest(connect)

	resp := w.Result()
	defer resp.Body.Close()
	return resp
}

func setupCreateFamilyTree(body []byte, t *testing.T) *http.Response {
	setupTestEnv(t)
	familyService, connect := setupDbTest(t)
	handler := &api.Handler{FamilyService: *familyService}

	router := mux.NewRouter()
	router.HandleFunc("/family/member", handler.InsertMember).Methods("POST")

	req := httptest.NewRequest(http.MethodPost, "/family/member", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	afterDbTest(connect)

	resp := w.Result()
	defer resp.Body.Close()
	return resp
}
