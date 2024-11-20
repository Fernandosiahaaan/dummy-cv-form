package model

import "time"

type Skill struct {
	ID          int64      `json:"id"          db:"id"`
	ProfileCode int64      `json:"profileCode" db:"profile_code"`
	Skill       string     `json:"skill"       db:"skill"`
	Level       string     `json:"level"       db:"level"`
	CreatedAt   time.Time  `json:"-"           db:"created_at"`
	UpdatedAt   time.Time  `json:"-"           db:"updated_at"`
	DeletedAt   *time.Time `json:"-"           db:"deleted_at"`
}
