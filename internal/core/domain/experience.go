package domain

import (
	"strings"
	"time"
)

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

// NewExperience creates a validated Experience. Returns error if validation fails.
func NewExperience(companyName, jobTitle, location string, startDate time.Time, endDate *time.Time, description, highlights string) (Experience, error) {
	e := Experience{
		CompanyName: strings.TrimSpace(companyName),
		JobTitle:    strings.TrimSpace(jobTitle),
		Location:    strings.TrimSpace(location),
		StartDate:   startDate,
		EndDate:     endDate,
		Description: strings.TrimSpace(description),
		Highlights:  strings.TrimSpace(highlights),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := e.Validate(); err != nil {
		return Experience{}, err
	}

	return e, nil
}

// Validate checks all business rules for Experience.
func (e Experience) Validate() error {
	if e.CompanyName == "" {
		return ErrEmptyCompanyName
	}
	if e.JobTitle == "" {
		return ErrEmptyJobTitle
	}
	if e.StartDate.IsZero() {
		return ErrEmptyStartDate
	}
	if e.Description == "" {
		return ErrEmptyDescription
	}
	if e.EndDate != nil && e.EndDate.Before(e.StartDate) {
		return ErrEndDateBeforeStart
	}
	return nil
}

// IsCurrent returns true if this is the current position (no end date).
func (e Experience) IsCurrent() bool {
	return e.EndDate == nil
}

// Duration returns the duration of the experience.
func (e Experience) Duration() time.Duration {
	end := time.Now()
	if e.EndDate != nil {
		end = *e.EndDate
	}
	return end.Sub(e.StartDate)
}
