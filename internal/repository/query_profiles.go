package repository

import (
	"context"
	"database/sql"
	"dummy-cv-form/internal/model"
	"fmt"
)

func (r *Repository) CreateNewProfile(profile *model.Profile) (int64, error) {
	var profileCode int64
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		INSERT INTO profiles (
			wanted_job_title, 
			first_name, 
			last_name, 
			email, 
			phone, 
			country, 
			city, 
			address, 
			postal_code, 
			driving_license, 
			nationality, 
			place_of_birth, 
			date_of_birth, 
			photo_url, 
			created_at, 
			updated_at, 
			deleted_at
		) 
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL
		)
		RETURNING profile_code
		;
	`

	err := r.DB.QueryRowContext(ctx, query,
		profile.WantedJobTitle,
		profile.FirstName,
		profile.LastName,
		profile.Email,
		profile.Phone,
		profile.Country,
		profile.City,
		profile.Address,
		profile.PostalCode,
		profile.DrivingLicense,
		profile.Nationality,
		profile.PlaceOfBirth,
		profile.DateOfBirth,
		profile.PhotoURL,
	).Scan(&profileCode)
	if err != nil {
		return 0, fmt.Errorf("failed to create new profile with name '%s'. err : %w", profile.FirstName, err)
	}

	return profileCode, nil
}

func (r *Repository) GetProfileByCode(profileCode int64) (*model.Profile, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	var profile model.Profile
	query := `
		SELECT 
			profile_code,
			wanted_job_title, 
			first_name, 
			last_name, 
			email, 
			phone, 
			country, 
			city, 
			address, 
			postal_code, 
			driving_license,  
			nationality, 
			place_of_birth, 
			date_of_birth, 
			photo_url, 
			created_at, 
			updated_at, 
			deleted_at
		FROM profiles
		WHERE profile_code = $1 AND deleted_at IS NULL
		;
	`
	err := r.DB.QueryRowContext(ctx, query, profileCode).Scan(
		&profile.ProfileCode,
		&profile.WantedJobTitle,
		&profile.FirstName,
		&profile.LastName,
		&profile.Email,
		&profile.Phone,
		&profile.Country,
		&profile.City,
		&profile.Address,
		&profile.PostalCode,
		&profile.DrivingLicense,
		&profile.Nationality,
		&profile.PlaceOfBirth,
		&profile.DateOfBirth,
		&profile.PhotoURL,
		&profile.CreatedAt,
		&profile.UpdatedAt,
		&profile.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read profile with profile_code: %d. err = %w", profileCode, err)
	}

	return &profile, nil
}

func (r *Repository) GetProfileByEmail(email string) (*model.Profile, error) {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	var profile model.Profile
	query := `
		SELECT 
			profile_code,
			wanted_job_title, 
			first_name, 
			last_name, 
			email, 
			phone, 
			country, 
			city, 
			address, 
			postal_code, 
			driving_license,  
			nationality, 
			place_of_birth, 
			date_of_birth, 
			photo_url, 
			created_at, 
			updated_at, 
			deleted_at
		FROM profiles
		WHERE email = $1 AND deleted_at IS NULL
		;
	`
	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&profile.ProfileCode,
		&profile.WantedJobTitle,
		&profile.FirstName,
		&profile.LastName,
		&profile.Email,
		&profile.Phone,
		&profile.Country,
		&profile.City,
		&profile.Address,
		&profile.PostalCode,
		&profile.DrivingLicense,
		&profile.Nationality,
		&profile.PlaceOfBirth,
		&profile.DateOfBirth,
		&profile.PhotoURL,
		&profile.CreatedAt,
		&profile.UpdatedAt,
		&profile.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &profile, nil
}

func (r *Repository) UpdateProfileByCode(profileCode int64, profile *model.Profile) error {
	ctx, cancel := context.WithTimeout(r.Ctx, defaultTimeoutQuery)
	defer cancel()

	query := `
		UPDATE profiles
		SET 
			wanted_job_title = $1,
			first_name = $2,
			last_name = $3,
			email = $4,
			phone = $5,
			country = $6,
			city = $7,
			address = $8,
			postal_code = $9,
			driving_license = $10,
			nationality = $11,
			place_of_birth = $12,
			date_of_birth = $13,
			photo_url = $14,
			updated_at = $15
		WHERE profile_code = $16 AND deleted_at IS NULL
		;
	`
	_, err := r.DB.ExecContext(ctx, query,
		profile.WantedJobTitle,
		profile.FirstName,
		profile.LastName,
		profile.Email,
		profile.Phone,
		profile.Country,
		profile.City,
		profile.Address,
		profile.PostalCode,
		profile.DrivingLicense,
		profile.Nationality,
		profile.PlaceOfBirth,
		profile.DateOfBirth,
		profile.PhotoURL,
		profile.UpdatedAt,
		profileCode,
	)
	if err != nil {
		return fmt.Errorf("failed to update profile with profile_code '%d'. err : %w", profileCode, err)
	}

	return nil
}
