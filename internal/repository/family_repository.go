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

func (r *FamilyRepository) Create(family *entity.Family) error {
	query := `INSERT INTO family (name, parent_id, name, parent1_id, parent2_id) VALUES ($1, $2, $3, $4) RETURNING id`

	rows, err := r.db.Query(query, family.Name, family.ParentId1, family.ParentId2)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&family.Id); err != nil {
			return fmt.Errorf("failed to create family: %w", err)
		}
	}

	return nil
}

func (r *FamilyRepository) FindFamilyById(id int64) ([]entity.Family, error) {
	var familyList []entity.Family

	query := `
		WITH RECURSIVE family_tree_recursive(id, name, parent1_id, parent2_id, generation) AS (
			SELECT id, name, parent1_id, parent2_id, 0 FROM family_tree WHERE id = $1
			UNION ALL
			SELECT ft.id, ft.name, ft.parent1_id, ft.parent2_id, generation + 1
			FROM family_tree ft
			JOIN family_tree_recursive ftr ON ft.id = ftr.parent1_id OR ft.id = ftr.parent2_id
		)
		SELECT id, name, parent1_id, parent2_id, generation
		FROM family_tree_recursive
		WHERE id != $1;
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return familyList, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var family entity.Family
		if err := rows.Scan(&family.Id, &family.Name, &family.ParentId1, &family.ParentId2, &family.Generation); err != nil {
			return familyList, fmt.Errorf("failed to scan family: %w", err)
		}
		familyList = append(familyList, family)
	}

	return familyList, nil
}

func (r *FamilyRepository) FindById(id int64) (entity.Family, error) {
	var family entity.Family

	query := `
		SELECT id, name, parent1_id, parent2_id FROM family_tree WHERE id = $1;
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return family, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&family.Id, &family.Name, &family.ParentId1, &family.ParentId2); err != nil {
			return family, fmt.Errorf("failed to scan family: %w", err)
		}
	}

	return family, nil
}
