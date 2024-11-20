package model

import (
	"time"
)

// Profile struct untuk tabel `employments`
type Employment struct {
	ID          int64      `json:"id"          db:"id"`
	ProfileCode int64      `json:"profileCode" db:"profile_code"`
	JobTitle    string     `json:"jobTitle"    db:"job_title"`
	Employer    string     `json:"employer"    db:"employer"`
	StartDate   string     `json:"startDate"   db:"start_date"`
	EndDate     string     `json:"endDate"     db:"end_date"`
	City        string     `json:"city"        db:"city"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"-"           db:"created_at"`
	UpdatedAt   time.Time  `json:"-"           db:"updated_at"`
	DeletedAt   *time.Time `json:"-"           db:"deleted_at"`
}
