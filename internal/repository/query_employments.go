package repository

import (
	"context"
	"dummy-cv-form/internal/model"
	"fmt"
)

func (r *Repository) CreateNewEmployment(emp *model.Employment) (int64, error) {
	var employmentID int64
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	// insert a new employment
	query := `
		INSERT INTO employments (
			profile_code, 
			job_title, 
			employer, 
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
		emp.ProfileCode,
		emp.JobTitle,
		emp.Employer,
		emp.StartDate,
		emp.EndDate,
		emp.City,
		emp.Description,
	).Scan(&employmentID)
	if err != nil {
		return 0, fmt.Errorf("failed to create new employment for profile %d. err: %w", emp.ProfileCode, err)
	}

	return employmentID, nil
}

func (r *Repository) GetEmploymentByID(id int64) (*model.Employment, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	// select employment by ID
	query := `
		SELECT id, profile_code, job_title, employer, start_date, end_date, city, description, created_at, updated_at, deleted_at
		FROM employments
		WHERE id = $1 AND deleted_at IS NULL
		;
	`

	emp := &model.Employment{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&emp.ID,
		&emp.ProfileCode,
		&emp.JobTitle,
		&emp.Employer,
		&emp.StartDate,
		&emp.EndDate,
		&emp.City,
		&emp.Description,
		&emp.CreatedAt,
		&emp.UpdatedAt,
		&emp.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get employment by ID %d. err: %w", id, err)
	}

	return emp, nil
}

func (r *Repository) GetEmploymentsByProfileCode(profileCode int64) ([]*model.Employment, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		SELECT id, profile_code, job_title, employer, start_date, end_date, city, description, created_at, updated_at, deleted_at
		FROM employments
		WHERE profile_code = $1 AND deleted_at IS NULL
		;
	`

	rows, err := r.DB.QueryContext(ctx, query, profileCode)
	if err != nil {
		return nil, fmt.Errorf("failed to get employments for profile_code %d. err: %w", profileCode, err)
	}
	defer rows.Close()

	var employments []*model.Employment

	for rows.Next() {
		employment := &model.Employment{}
		err := rows.Scan(
			&employment.ID,
			&employment.ProfileCode,
			&employment.JobTitle,
			&employment.Employer,
			&employment.StartDate,
			&employment.EndDate,
			&employment.City,
			&employment.Description,
			&employment.CreatedAt,
			&employment.UpdatedAt,
			&employment.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan employment row. err: %w", err)
		}
		employments = append(employments, employment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over employment rows. err: %w", err)
	}

	return employments, nil
}

func (r *Repository) SoftDeleteEmployment(id int64) error {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	// update deleted_at to perform a soft delete
	query := `
		UPDATE employments
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
		;
	`

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to soft delete employment with ID %d. err: %w", id, err)
	}

	return nil
}
