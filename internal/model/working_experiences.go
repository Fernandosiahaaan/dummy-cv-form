package model

import "time"

type WorkingExperience struct {
	ID          int64      `json:"id"                db:"id"`
	ProfileCode int64      `json:"profileCode"       db:"profile_code"`
	Experience  string     `json:"workingExperience" db:"experience"`
	CreatedAt   time.Time  `json:"-"                 db:"created_at"`
	UpdatedAt   time.Time  `json:"-"                 db:"updated_at"`
	DeletedAt   *time.Time `json:"-"                 db:"deleted_at"`
}
