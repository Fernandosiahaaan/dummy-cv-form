package model

import (
	"time"
)

// Profile struct untuk tabel `profiles`
type Profile struct {
	ProfileCode    int64      `json:"profileCode"         db:"profile_code"`
	WantedJobTitle string     `json:"wantedJobTitle"      db:"wanted_job_title"`
	FirstName      string     `json:"firstName"           db:"first_name"`
	LastName       string     `json:"lastName"            db:"last_name"`
	Email          string     `json:"email"               db:"email"`
	Phone          string     `json:"phone"               db:"phone"`
	Country        string     `json:"country"             db:"country"`
	City           string     `json:"city"                db:"city"`
	Address        string     `json:"address"             db:"address"`
	PostalCode     int        `json:"postalCode"          db:"postal_code"`
	DrivingLicense string     `json:"drivingLicense"      db:"driving_license"`
	Nationality    string     `json:"nationality"         db:"nationality"`
	PlaceOfBirth   string     `json:"placeOfBirth"        db:"place_of_birth"`
	DateOfBirth    string     `json:"dateOfBirth"         db:"date_of_birth"`
	PhotoURL       string     `json:"photoUrl"            db:"photo_url"`
	CreatedAt      time.Time  `json:"-"                   db:"created_at"`
	UpdatedAt      time.Time  `json:"-"                   db:"updated_at"`
	DeletedAt      *time.Time `json:"-,omitempty"         db:"deleted_at"` // default null
}
