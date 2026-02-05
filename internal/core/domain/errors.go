package domain

import "errors"

var (

	// ErrEmptyName represents an error indicating that a name cannot be empty.
	ErrEmptyName = errors.New("name cannot be empty")
	// ErrEmptyTitle represents an error indicating that a title cannot be empty.
	ErrEmptyTitle = errors.New("title cannot be empty")
	// ErrEmptyDescription represents an error indicating that a description cannot be empty.
	ErrEmptyDescription = errors.New("description cannot be empty")
	// ErrEmptyCategory represents an error indicating that a category cannot be empty.
	ErrEmptyCategory = errors.New("category cannot be empty")
	// ErrEmptyCompanyName represents an error indicating that a company name cannot be empty.
	ErrEmptyCompanyName = errors.New("company name cannot be empty")
	// ErrEmptyJobTitle represents an error indicating that a job title cannot be empty.
	ErrEmptyJobTitle = errors.New("job title cannot be empty")
	// ErrEmptyStartDate represents an error indicating that a start date is required.
	ErrEmptyStartDate = errors.New("start date is required")
	// ErrEndDateBeforeStart represents an error indicating that an end date cannot be before the start date.
	ErrEndDateBeforeStart = errors.New("end date cannot be before start date")
	// ErrInvalidProficiency represents an error indicating that proficiency must be between 0 and 100.
	ErrInvalidProficiency = errors.New("proficiency must be between 0 and 100")
)
