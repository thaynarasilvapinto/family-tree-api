package repository

import (
	"fmt"

	"github.com/thaynarasilvapinto/family-tree-api/internal/adapter/postgres"
	"github.com/thaynarasilvapinto/family-tree-api/internal/entity"
)

type FamilyRepository struct {
	db *postgres.PostgresDatabase
}

func NewFamilyRepository(db *postgres.PostgresDatabase) *FamilyRepository {
	return &FamilyRepository{db: db}
}

func (r *FamilyRepository) Create(f *entity.Family) error {
	query := `INSERT INTO family (name, lft, rgt, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`

	rows, err := r.db.Query(query, f.Name, f.Lft, f.Rgt, f.ParentId)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&f.ID); err != nil {
			return fmt.Errorf("failed to create family: %w", err)
		}
	}

	return nil
}
