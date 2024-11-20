package repository

import (
	"context"
	"dummy-cv-form/internal/model"
	"fmt"
)

func (r *Repository) CreateNewEducation(edu *model.Education) (int64, error) {
	var educationID int64
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		INSERT INTO educations (
			profile_code, 
			school, 
			degree, 
			start_date, 
			end_date, 
			city, 
			description, 
			created_at, 
			updated_at, 
			deleted_at
		) 
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL
		)
		RETURNING id
		;
	`

	err := r.DB.QueryRowContext(ctx, query,
		edu.ProfileCode,
		edu.School,
		edu.Degree,
		edu.StartDate,
		edu.EndDate,
		edu.City,
		edu.Description,
	).Scan(&educationID)
	if err != nil {
		return 0, fmt.Errorf("failed to create new education for profile %d. err: %w", edu.ProfileCode, err)
	}

	return educationID, nil
}

func (r *Repository) GetEducationByID(id int64) (*model.Education, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		SELECT id, profile_code, school, degree, start_date, end_date, city, description, created_at, updated_at, deleted_at
		FROM educations
		WHERE id = $1 AND deleted_at IS NULL
		;
	`

	edu := &model.Education{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&edu.ID,
		&edu.ProfileCode,
		&edu.School,
		&edu.Degree,
		&edu.StartDate,
		&edu.EndDate,
		&edu.City,
		&edu.Description,
		&edu.CreatedAt,
		&edu.UpdatedAt,
		&edu.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get education by ID %d. err: %w", id, err)
	}

	return edu, nil
}

func (r *Repository) GetEducationsByProfileCode(profileCode int64) ([]*model.Education, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	// Query untuk mengambil data pendidikan berdasarkan profile_code
	query := `
		SELECT id, profile_code, school, degree, start_date, end_date, city, description, created_at, updated_at, deleted_at
		FROM educations
		WHERE profile_code = $1 AND deleted_at IS NULL
		;
	`

	rows, err := r.DB.QueryContext(ctx, query, profileCode)
	if err != nil {
		return nil, fmt.Errorf("failed to get educations for profile_code %d. err: %w", profileCode, err)
	}
	defer rows.Close()

	var educations []*model.Education

	for rows.Next() {
		education := &model.Education{}
		err := rows.Scan(
			&education.ID,
			&education.ProfileCode,
			&education.School,
			&education.Degree,
			&education.StartDate,
			&education.EndDate,
			&education.City,
			&education.Description,
			&education.CreatedAt,
			&education.UpdatedAt,
			&education.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan education row. err: %w", err)
		}
		educations = append(educations, education)
	}

	// Cek apakah ada error saat iterasi
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over education rows. err: %w", err)
	}

	return educations, nil
}

func (r *Repository) SoftDeleteEducation(id int64) error {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		UPDATE educations
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND deleted_at IS NULL
		;
	`

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to soft delete education with id %d. err: %w", id, err)
	}

	return nil
}
