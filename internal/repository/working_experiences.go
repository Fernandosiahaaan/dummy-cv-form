package repository

import (
	"context"
	"database/sql"
	"dummy-cv-form/internal/model"
	"fmt"
)

func (r *Repository) CreateNewWorkingExperience(workingExperience *model.WorkingExperience) (int64, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		INSERT INTO working_experiences (
			profile_code, 
			experience, 
			created_at, 
			updated_at,
			deleted_at
		)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL)
		RETURNING id
		;
	`

	var id int64
	err := r.DB.QueryRowContext(ctx, query, workingExperience.ProfileCode, workingExperience.Experience).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create new working experience. err: %w", err)
	}

	return id, nil
}

func (r *Repository) GetWorkingExperienceByProfileCode(profileCode int64) (*model.WorkingExperience, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		SELECT 
			id, 
			profile_code, 
			experience, 
			created_at, 
			updated_at, 
			deleted_at
		FROM working_experiences
		WHERE profile_code = $1 AND deleted_at IS NULL
		LIMIT 1
		;
	`

	var workingExperience model.WorkingExperience
	err := r.DB.QueryRowContext(ctx, query, profileCode).Scan(
		&workingExperience.ID,
		&workingExperience.ProfileCode,
		&workingExperience.Experience,
		&workingExperience.CreatedAt,
		&workingExperience.UpdatedAt,
		&workingExperience.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get working experience for profile_code %d. err: %w", profileCode, err)
	}

	return &workingExperience, nil
}

func (r *Repository) UpdateWorkingExperienceByProfileCode(workingExperience *model.WorkingExperience) error {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		UPDATE working_experiences
		SET 
			experience = $1, 
			updated_at = CURRENT_TIMESTAMP
		WHERE profile_code = $2 AND deleted_at IS NULL
		RETURNING id, experience, updated_at;
	`

	err := r.DB.QueryRowContext(ctx, query, workingExperience.Experience, workingExperience.ProfileCode).
		Scan(&workingExperience.ID, &workingExperience.Experience, &workingExperience.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to update working experience with id %d. err: %w", workingExperience.ID, err)
	}

	return nil
}
