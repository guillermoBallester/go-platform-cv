package domain

import "time"

// Achievement represents an accomplishment, optionally tied to an experience or project.
type Achievement struct {
	ID           int32
	Title        string
	Description  string
	Date         *time.Time
	ExperienceID *int32  // Optional link to experience
	ProjectID    *int32  // Optional link to project
	Skills       []Skill // Related skills
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
