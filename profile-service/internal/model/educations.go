package model

import "time"

type Education struct {
	ID          int64      `json:"id"          db:"id"`
	ProfileCode int64      `json:"profileCode" db:"profile_code"`
	School      string     `json:"school"      db:"school"`
	Degree      string     `json:"degree"      db:"degree"`
	StartDate   string     `json:"startDate"   db:"start_date"`
	EndDate     string     `json:"endDate"     db:"end_date"`
	City        string     `json:"city"        db:"city"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"-"           db:"created_at"`
	UpdatedAt   time.Time  `json:"-"           db:"updated_at"`
	DeletedAt   *time.Time `json:"-"           db:"deleted_at"`
}
