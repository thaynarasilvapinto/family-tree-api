package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	entity "github.com/thaynarasilvapinto/family-tree-api/internal/entity"
)

func (h *Handler) GetFamilyTree(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Unexpected error trying to identify which person you are looking for.", http.StatusBadRequest)
		return
	}

	person, err := h.FamilyService.FindById(id)

	if err != nil || person.Id == 0 {
		sendErrorResponse(w, "Unexpected error fetching this person's family.", http.StatusBadRequest)
		return
	}

	response, err := h.FamilyService.FindFamilyById(id)

	if err != nil {
		sendErrorResponse(w, "Error finding the person you are looking for. Please try another id.", http.StatusBadRequest)
		return
	}

	family := dtoFamily(response, person)

	resp, err := json.Marshal(family)
	if err != nil {
		sendErrorResponse(w, "Unexpected error in serialize object", http.StatusTeapot)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func dtoFamily(response []entity.Family, person entity.Family) FamilyResponse {
	relatives := make(map[int]string)
	relatives[1] = "parents"
	relatives[2] = "grandparents"
	relatives[3] = "great grandparents"

	var family FamilyResponse
	var members []MemberResponse

	for _, item := range response {
		var member MemberResponse
		member.Id = item.Id
		member.Name = item.Name
		if *item.Generation >= 3 {
			member.Relationship = relatives[3]
		} else {
			member.Relationship = relatives[*item.Generation]
		}

		members = append(members, member)
	}

	family.Id = person.Id
	family.Name = person.Name
	family.Members = members
	return family
}
