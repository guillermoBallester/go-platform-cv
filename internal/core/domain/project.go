package domain

import "time"

// Project represents a portfolio project.
type Project struct {
	ID          int32
	Name        string
	Description string
	StartDate   *time.Time
	EndDate     *time.Time
	Skills      []Skill // Related skills
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
