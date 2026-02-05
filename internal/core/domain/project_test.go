package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProject(t *testing.T) {
	now := time.Now()
	past := now.AddDate(-1, 0, 0)
	future := now.AddDate(1, 0, 0)

	tests := []struct {
		name        string
		projName    string
		description string
		startDate   *time.Time
		endDate     *time.Time
		wantErr     error
	}{
		{
			name:        "valid project with dates",
			projName:    "My Project",
			description: "A cool project",
			startDate:   &past,
			endDate:     &now,
			wantErr:     nil,
		},
		{
			name:        "valid project without dates",
			projName:    "Side Project",
			description: "Just for fun",
			startDate:   nil,
			endDate:     nil,
			wantErr:     nil,
		},
		{
			name:        "valid ongoing project",
			projName:    "Active Project",
			description: "Still working on it",
			startDate:   &past,
			endDate:     nil,
			wantErr:     nil,
		},
		{
			name:        "empty name",
			projName:    "",
			description: "Some description",
			startDate:   nil,
			endDate:     nil,
			wantErr:     ErrEmptyName,
		},
		{
			name:        "whitespace name",
			projName:    "   ",
			description: "Some description",
			startDate:   nil,
			endDate:     nil,
			wantErr:     ErrEmptyName,
		},
		{
			name:        "empty description",
			projName:    "Project",
			description: "",
			startDate:   nil,
			endDate:     nil,
			wantErr:     ErrEmptyDescription,
		},
		{
			name:        "end date before start date",
			projName:    "Project",
			description: "Description",
			startDate:   &now,
			endDate:     &past,
			wantErr:     ErrEndDateBeforeStart,
		},
		{
			name:        "future dates are valid",
			projName:    "Future Project",
			description: "Planning ahead",
			startDate:   &now,
			endDate:     &future,
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proj, err := NewProject(tt.projName, tt.description, tt.startDate, tt.endDate)

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Equal(t, Project{}, proj)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.projName, proj.Name)
				assert.Equal(t, tt.description, proj.Description)
			}
		})
	}
}

func TestProject_IsOngoing(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		endDate  *time.Time
		expected bool
	}{
		{"nil end date is ongoing", nil, true},
		{"with end date is not ongoing", &now, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proj := Project{EndDate: tt.endDate}
			assert.Equal(t, tt.expected, proj.IsOngoing())
		})
	}
}

func TestProject_Duration(t *testing.T) {
	now := time.Now()
	oneYearAgo := now.AddDate(-1, 0, 0)

	t.Run("no start date returns zero", func(t *testing.T) {
		proj := Project{StartDate: nil}
		assert.Equal(t, time.Duration(0), proj.Duration())
	})

	t.Run("ongoing project duration", func(t *testing.T) {
		proj := Project{StartDate: &oneYearAgo, EndDate: nil}
		duration := proj.Duration()
		assert.True(t, duration >= 365*24*time.Hour)
	})

	t.Run("completed project duration", func(t *testing.T) {
		sixMonthsAgo := now.AddDate(0, -6, 0)
		proj := Project{StartDate: &oneYearAgo, EndDate: &sixMonthsAgo}
		duration := proj.Duration()
		assert.True(t, duration >= 180*24*time.Hour && duration <= 185*24*time.Hour)
	})
}
