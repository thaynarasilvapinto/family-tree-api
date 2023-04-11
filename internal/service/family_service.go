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

func (s *FamilyService) FindFamilyById(id int64) ([]entity.Family, error) {
	return s.repo.FindFamilyById(id)
}

func (s *FamilyService) FindById(id int64) (entity.Family, error) {
	return s.repo.FindById(id)
}
