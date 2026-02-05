package domain

import (
	"strings"
	"time"
)

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

// NewAchievement creates a validated Achievement. Returns error if validation fails.
func NewAchievement(title, description string, date *time.Time, experienceID, projectID *int32) (Achievement, error) {
	a := Achievement{
		Title:        strings.TrimSpace(title),
		Description:  strings.TrimSpace(description),
		Date:         date,
		ExperienceID: experienceID,
		ProjectID:    projectID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := a.Validate(); err != nil {
		return Achievement{}, err
	}

	return a, nil
}

// Validate checks all business rules for Achievement.
func (a Achievement) Validate() error {
	if a.Title == "" {
		return ErrEmptyTitle
	}
	if a.Description == "" {
		return ErrEmptyDescription
	}
	return nil
}

// HasContext returns true if the achievement is linked to an experience or project.
func (a Achievement) HasContext() bool {
	return a.ExperienceID != nil || a.ProjectID != nil
}

// IsLinkedToExperience returns true if linked to an experience.
func (a Achievement) IsLinkedToExperience() bool {
	return a.ExperienceID != nil
}

// IsLinkedToProject returns true if linked to a project.
func (a Achievement) IsLinkedToProject() bool {
	return a.ProjectID != nil
}
