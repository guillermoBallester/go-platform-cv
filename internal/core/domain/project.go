package domain

import (
	"strings"
	"time"
)

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

// NewProject creates a validated Project. Returns error if validation fails.
func NewProject(name, description string, startDate, endDate *time.Time) (Project, error) {
	p := Project{
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := p.Validate(); err != nil {
		return Project{}, err
	}

	return p, nil
}

// Validate checks all business rules for Project.
func (p Project) Validate() error {
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Description == "" {
		return ErrEmptyDescription
	}
	if p.StartDate != nil && p.EndDate != nil && p.EndDate.Before(*p.StartDate) {
		return ErrEndDateBeforeStart
	}
	return nil
}

// IsOngoing returns true if the project has no end date.
func (p Project) IsOngoing() bool {
	return p.EndDate == nil
}

// Duration returns the duration of the project, or zero if no dates set.
func (p Project) Duration() time.Duration {
	if p.StartDate == nil {
		return 0
	}
	end := time.Now()
	if p.EndDate != nil {
		end = *p.EndDate
	}
	return end.Sub(*p.StartDate)
}
