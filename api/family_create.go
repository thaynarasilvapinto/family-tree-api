package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	entity "github.com/thaynarasilvapinto/family-tree-api/internal/entity"
)

/// SE O FILHO JÁ TEM DOIS PAIS (BLOQUEIA)
/// VOCE NÃO PODE SER PAI DOS SEUS ANTEPASSADOS

func (h *Handler) InsertMember(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	var memberRequest MemberRequest
	json.Unmarshal(body, &memberRequest)

	h.validateParent(memberRequest)

	var family entity.Family
	family.Name = memberRequest.Name
	family.ParentId1 = memberRequest.Relationship.Parent.Parent1
	family.ParentId2 = memberRequest.Relationship.Parent.Parent2

	h.FamilyService.Create(&family)

	resp, err := json.Marshal(family)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *Handler) validateParent(memberRequest MemberRequest) error {
	var err error
	validateParentID := func(parentID *int64) {
		if parentID != nil {
			person, _ := h.FamilyService.FindById(*parentID)
			if person.Id == 0 {
				err = errors.New("parent not found in the database, please register a parent before trying to insert a child")
			}
		}
	}
	//certificar que os pais não sao irmaos ou tem uma relacao de incexto

	validateParentID(memberRequest.Relationship.Parent.Parent1)
	validateParentID(memberRequest.Relationship.Parent.Parent2)
	return err
}
