package test

import (
	"encoding/json"
	"net/http"
	"testing"

	_ "github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	api "github.com/thaynarasilvapinto/family-tree-api/api"
	entity "github.com/thaynarasilvapinto/family-tree-api/internal/entity"
)

func TestFamilyTree(t *testing.T) {

	t.Run("FamilyTree", func(t *testing.T) {

		resp := setupGetFamilyTree("10", t)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var actualResult api.FamilyResponse
		err := json.NewDecoder(resp.Body).Decode(&actualResult)
		if err != nil {
			t.Errorf("Unexpected error in deserialize response body: %v", err)
		}

		if len(actualResult.Members) != 4 {
			t.Errorf("Expected 4 family members, got %d", len(actualResult.Members))
		}

		expectedResult := api.FamilyResponse{
			Id:   10,
			Name: "Sophie",
			Members: []api.MemberResponse{
				{Id: 5, Name: "Lucy", Relationship: "parents"},
				{Id: 6, Name: "Marcos", Relationship: "parents"},
				{Id: 1, Name: "John", Relationship: "grandparents"},
				{Id: 2, Name: "Mary", Relationship: "grandparents"},
			},
		}

		assert.Equal(t, expectedResult.Id, actualResult.Id)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
		assert.ElementsMatch(t, expectedResult.Members, actualResult.Members)
	})
}

func TestInsertMember_SuccessfulInsertion(t *testing.T) {

	t.Run("FamilyTree", func(t *testing.T) {

		fakeBody := []byte(`
			{
				"name":"João",
				"relationship":{
					"parent":{
						"parent1":1,
						"parent2":2
					},
					"children":[]
				}
			}`)

		resp := setupCreateFamilyTree(fakeBody, t)

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response entity.Family
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Errorf("Unexpected error in deserialize response body: %v", err)
		}
		assert.Equal(t, "João", response.Name)
		assert.Equal(t, int64(1), *response.ParentId1)
		assert.Equal(t, int64(2), *response.ParentId2)
	})
}

func TestInsertMember_DuplicateChildID(t *testing.T) {

	t.Run("FamilyTree", func(t *testing.T) {

		fakeBody := []byte(`
			{"name": "João","relationship": {"parent": {"parent1": 1,"parent2": 1},"children": []}}	
		`)

		resp := setupCreateFamilyTree(fakeBody, t)

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response api.ErrorMessage
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Errorf("Unexpected error in deserialize response body: %v", err)
		}
		assert.Equal(t, "these parents cannot be parents since they are related", response.Message)
	})
}

func TestInsertMember_ChildNotFound(t *testing.T) {

	t.Run("FamilyTree", func(t *testing.T) {

		fakeBody := []byte(`{"name":"Test Member","relationship":{"parent":{"parent1":1,"parent2":2},"children":[999]}}`)

		resp := setupCreateFamilyTree(fakeBody, t)

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response api.ErrorMessage
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Errorf("Unexpected error in deserialize response body: %v", err)
		}
		assert.Equal(t, "parent not found in the database, please register a parent before trying to insert a child", response.Message)
	})
}

func TestInsertMember_RelatedParents(t *testing.T) {

	t.Run("FamilyTree", func(t *testing.T) {

		fakeBody := []byte(`{"name":"Test Member","relationship":{"parent":{"parent1":1,"parent2":1},"children":[3,4]}}`)

		resp := setupCreateFamilyTree(fakeBody, t)

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response api.ErrorMessage
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Errorf("Unexpected error in deserialize response body: %v", err)
		}
		assert.Equal(t, "these parents cannot be parents since they are related", response.Message)
	})
}

func TestInsertMember_ParentNotFound(t *testing.T) {

	t.Run("FamilyTree", func(t *testing.T) {

		fakeBody := []byte(`{"name":"Test Member","relationship":{"parent":{"parent1":999,"parent2":888},"children":[3,4]}}`)

		resp := setupCreateFamilyTree(fakeBody, t)

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var response api.ErrorMessage
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Errorf("Unexpected error in deserialize response body: %v", err)
		}
		assert.Equal(t, "parent not found in the database, please register a parent before trying to insert a child", response.Message)
	})
}
