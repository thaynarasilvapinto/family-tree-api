package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	entity "github.com/thaynarasilvapinto/family-tree-api/internal/entity"
)

func (h *Handler) InsertMember(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	var memberRequest MemberRequest
	json.Unmarshal(body, &memberRequest)

	err := h.parentaValidation(memberRequest)

	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.childrenValidation(memberRequest.Relationship.Children, memberRequest.Relationship.Parent.Parent1, memberRequest.Relationship.Parent.Parent2)

	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var family entity.Family
	family.Name = memberRequest.Name
	family.ParentId1 = memberRequest.Relationship.Parent.Parent1
	family.ParentId2 = memberRequest.Relationship.Parent.Parent2

	h.FamilyService.Create(&family)

	resp, err := json.Marshal(family)
	if err != nil {
		sendErrorResponse(w, "Unexpected error in serialize object", http.StatusTeapot)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h *Handler) childrenValidation(childrens []int64, parent1, parent2 *int64) error {

	if err := validateUniqueIDs(childrens); err != nil {
		return err
	}

	for _, son := range childrens {
		if err := h.validationIfThePersonExists(&son); err != nil {
			return err
		}
		if err := h.validateParentRelationship(parent1, &son); err != nil {
			return err
		}
		if err := h.validateParentRelationship(parent2, &son); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) parentaValidation(memberRequest MemberRequest) error {

	if (memberRequest.Relationship.Parent.Parent1 != nil && memberRequest.Relationship.Parent.Parent2 != nil) && *memberRequest.Relationship.Parent.Parent1 == *memberRequest.Relationship.Parent.Parent2 {
		return errors.New("these parents cannot be parents since they are related")
	}

	if err := h.validationIfThePersonExists(memberRequest.Relationship.Parent.Parent1); err != nil {
		return err
	}
	if err := h.validationIfThePersonExists(memberRequest.Relationship.Parent.Parent2); err != nil {
		return err
	}
	if err := h.validateParentRelationship(memberRequest.Relationship.Parent.Parent2, memberRequest.Relationship.Parent.Parent1); err != nil {
		return err
	}
	if err := h.validateParentRelationship(memberRequest.Relationship.Parent.Parent1, memberRequest.Relationship.Parent.Parent2); err != nil {
		return err
	}
	return nil
}

func (h *Handler) validationIfThePersonExists(parentID *int64) error {
	if parentID != nil {
		person, _ := h.FamilyService.FindById(*parentID)
		if person.Id == 0 {
			return errors.New("parent not found in the database, please register a parent before trying to insert a child")
		}
	}
	return nil
}

func (h *Handler) validateParentRelationship(parent *int64, compareParent *int64) error {
	if compareParent == nil || parent == nil {
		return nil
	}

	family, err := h.FamilyService.FindFamilyById(*parent)
	if err != nil {
		return errors.New("it was not possible to validate the relationship between them")
	}

	for _, item := range family {
		if item.Id == *compareParent {
			return errors.New("these parents cannot be parents since they are related")

		}
	}
	return nil
}

func validateUniqueIDs(ids []int64) error {
	seen := make(map[int64]bool)
	for _, id := range ids {
		if seen[id] {
			return errors.New(fmt.Sprintf("duplicate ID found: %s", id))
		}
		seen[id] = true
	}
	return nil
}
