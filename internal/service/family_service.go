package service

import (
	"github.com/thaynarasilvapinto/family-tree-api/internal/entity"
	"github.com/thaynarasilvapinto/family-tree-api/internal/repository"
)

type FamilyService struct {
	repo *repository.FamilyRepository
}

func NewFamilyService(repo *repository.FamilyRepository) *FamilyService {
	return &FamilyService{repo: repo}
}

func (s *FamilyService) Create(f *entity.Family) error {
	if err := s.repo.Create(f); err != nil {
		return err
	}
	return nil
}
