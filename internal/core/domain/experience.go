package domain

import "time"

// Experience represents a job experience entry.
type Experience struct {
	ID          int32
	CompanyName string
	JobTitle    string
	Location    string
	StartDate   time.Time
	EndDate     *time.Time // nil = current position
	Description string
	Highlights  string
	Skills      []Skill // Related skills
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
