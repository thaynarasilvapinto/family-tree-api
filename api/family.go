package api

import (
	"encoding/json"
	"net/http"

	familyService "github.com/thaynarasilvapinto/family-tree-api/internal/service"
	logger "github.com/thaynarasilvapinto/family-tree-api/logger"
)

type Handler struct {
	FamilyService familyService.FamilyService
}

type ErrorMessage struct {
	Message string `json:"nessage"`
}

type FamilyResponse struct {
	Id      int64            `json:"id"`
	Name    string           `json:"name"`
	Members []MemberResponse `json:"members"`
}

type MemberResponse struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
}

type MemberRequest struct {
	Name         string              `json:"name"`
	Relationship RelationshipRequest `json:"relationship"`
}

type RelationshipRequest struct {
	Parent   ParentRequest `json:"parent"`
	Children []int64       `json:"children"`
}

type ParentRequest struct {
	Parent1 *int64 `json:"parent1"`
	Parent2 *int64 `json:"parent2"`
}

func sendErrorResponse(w http.ResponseWriter, errorMessage string, status int) {
	logger.Error(errorMessage)
	resp, _ := json.Marshal(ErrorMessage{
		Message: errorMessage,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}
