package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetFamilyTree(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64) //TRATAR O ERRO

	response, _ := h.FamilyService.FindFamilyById(id) //TRATAR O ERRO
	person, _ := h.FamilyService.FindById(id)         //TRATAR O ERRO

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

	resp, err := json.Marshal(family)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
