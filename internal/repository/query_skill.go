package repository

import (
	"context"
	"dummy-cv-form/internal/model"
	"fmt"
)

func (r *Repository) CreateNewSkill(skill *model.Skill) (int64, error) {
	var skillID int64
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		INSERT INTO skills (
			profile_code, 
			skill, 
			level, 
			created_at, 
			updated_at, 
			deleted_at
		) 
		VALUES (
			$1, $2, $3, $4, $5, NULL
		)
		RETURNING id
		;
	`

	err := r.DB.QueryRowContext(ctx, query,
		skill.ProfileCode,
		skill.Skill,
		skill.Level,
		skill.CreatedAt,
		skill.UpdatedAt,
	).Scan(&skillID)
	if err != nil {
		return 0, fmt.Errorf("failed to create new skill for profile %d. err: %w", skill.ProfileCode, err)
	}

	return skillID, nil
}

func (r *Repository) GetSkillByID(id int64) (*model.Skill, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		SELECT id, profile_code, skill, level, created_at, updated_at, deleted_at
		FROM skills
		WHERE id = $1 AND deleted_at IS NULL
		;
	`

	skill := &model.Skill{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&skill.ID,
		&skill.ProfileCode,
		&skill.Skill,
		&skill.Level,
		&skill.CreatedAt,
		&skill.UpdatedAt,
		&skill.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get skill by ID %d. err: %w", id, err)
	}

	return skill, nil
}

func (r *Repository) GetSkillsByProfileCode(profileCode int64) ([]*model.Skill, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		SELECT id, profile_code, skill, level, created_at, updated_at, deleted_at
		FROM skills
		WHERE profile_code = $1 AND deleted_at IS NULL
		;
	`

	rows, err := r.DB.QueryContext(ctx, query, profileCode)
	if err != nil {
		return nil, fmt.Errorf("failed to get skills for profile_code %d. err: %w", profileCode, err)
	}
	defer rows.Close()

	var skills []*model.Skill

	for rows.Next() {
		skill := &model.Skill{}
		err := rows.Scan(
			&skill.ID,
			&skill.ProfileCode,
			&skill.Skill,
			&skill.Level,
			&skill.CreatedAt,
			&skill.UpdatedAt,
			&skill.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan skill row. err: %w", err)
		}
		skills = append(skills, skill)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over skill rows. err: %w", err)
	}

	return skills, nil
}

func (r *Repository) SoftDeleteSkill(id int64) error {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		UPDATE skills
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND deleted_at IS NULL
		;
	`

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to soft delete skill with id %d. err: %w", id, err)
	}

	return nil
}
