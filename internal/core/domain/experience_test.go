package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExperience(t *testing.T) {
	now := time.Now()
	past := now.AddDate(-1, 0, 0)
	future := now.AddDate(1, 0, 0)

	tests := []struct {
		name        string
		companyName string
		jobTitle    string
		location    string
		startDate   time.Time
		endDate     *time.Time
		description string
		highlights  string
		wantErr     error
	}{
		{
			name:        "valid experience with end date",
			companyName: "Acme Corp",
			jobTitle:    "Software Engineer",
			location:    "Remote",
			startDate:   past,
			endDate:     &now,
			description: "Building awesome software",
			highlights:  "Led team of 5",
			wantErr:     nil,
		},
		{
			name:        "valid current position (no end date)",
			companyName: "Tech Inc",
			jobTitle:    "Senior Developer",
			location:    "New York",
			startDate:   past,
			endDate:     nil,
			description: "Developing new features",
			highlights:  "",
			wantErr:     nil,
		},
		{
			name:        "empty company name",
			companyName: "",
			jobTitle:    "Developer",
			location:    "Remote",
			startDate:   past,
			endDate:     nil,
			description: "Work",
			highlights:  "",
			wantErr:     ErrEmptyCompanyName,
		},
		{
			name:        "whitespace company name",
			companyName: "   ",
			jobTitle:    "Developer",
			location:    "Remote",
			startDate:   past,
			endDate:     nil,
			description: "Work",
			highlights:  "",
			wantErr:     ErrEmptyCompanyName,
		},
		{
			name:        "empty job title",
			companyName: "Company",
			jobTitle:    "",
			location:    "Remote",
			startDate:   past,
			endDate:     nil,
			description: "Work",
			highlights:  "",
			wantErr:     ErrEmptyJobTitle,
		},
		{
			name:        "empty description",
			companyName: "Company",
			jobTitle:    "Developer",
			location:    "Remote",
			startDate:   past,
			endDate:     nil,
			description: "",
			highlights:  "",
			wantErr:     ErrEmptyDescription,
		},
		{
			name:        "zero start date",
			companyName: "Company",
			jobTitle:    "Developer",
			location:    "Remote",
			startDate:   time.Time{},
			endDate:     nil,
			description: "Work",
			highlights:  "",
			wantErr:     ErrEmptyStartDate,
		},
		{
			name:        "end date before start date",
			companyName: "Company",
			jobTitle:    "Developer",
			location:    "Remote",
			startDate:   now,
			endDate:     &past,
			description: "Work",
			highlights:  "",
			wantErr:     ErrEndDateBeforeStart,
		},
		{
			name:        "future end date is valid",
			companyName: "Company",
			jobTitle:    "Developer",
			location:    "Remote",
			startDate:   now,
			endDate:     &future,
			description: "Work",
			highlights:  "",
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp, err := NewExperience(tt.companyName, tt.jobTitle, tt.location, tt.startDate, tt.endDate, tt.description, tt.highlights)

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Equal(t, Experience{}, exp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.companyName, exp.CompanyName)
				assert.Equal(t, tt.jobTitle, exp.JobTitle)
			}
		})
	}
}

func TestExperience_IsCurrent(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		endDate  *time.Time
		expected bool
	}{
		{"nil end date is current", nil, true},
		{"with end date is not current", &now, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp := Experience{EndDate: tt.endDate}
			assert.Equal(t, tt.expected, exp.IsCurrent())
		})
	}
}

func TestExperience_Duration(t *testing.T) {
	now := time.Now()
	oneYearAgo := now.AddDate(-1, 0, 0)
	sixMonthsAgo := now.AddDate(0, -6, 0)

	t.Run("current position duration", func(t *testing.T) {
		exp := Experience{StartDate: oneYearAgo, EndDate: nil}
		duration := exp.Duration()
		assert.True(t, duration >= 365*24*time.Hour)
	})

	t.Run("completed position duration", func(t *testing.T) {
		exp := Experience{StartDate: oneYearAgo, EndDate: &sixMonthsAgo}
		duration := exp.Duration()
		assert.True(t, duration >= 180*24*time.Hour && duration <= 185*24*time.Hour)
	})
}
