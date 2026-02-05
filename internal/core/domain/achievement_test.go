package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAchievement(t *testing.T) {
	now := time.Now()
	expID := int32(1)
	projID := int32(2)

	tests := []struct {
		name         string
		title        string
		description  string
		date         *time.Time
		experienceID *int32
		projectID    *int32
		wantErr      error
	}{
		{
			name:         "valid achievement with all fields",
			title:        "Employee of the Month",
			description:  "Recognized for outstanding performance",
			date:         &now,
			experienceID: &expID,
			projectID:    nil,
			wantErr:      nil,
		},
		{
			name:         "valid achievement linked to project",
			title:        "Project Launch",
			description:  "Successfully launched the product",
			date:         &now,
			experienceID: nil,
			projectID:    &projID,
			wantErr:      nil,
		},
		{
			name:         "valid standalone achievement",
			title:        "Certification",
			description:  "Obtained AWS certification",
			date:         nil,
			experienceID: nil,
			projectID:    nil,
			wantErr:      nil,
		},
		{
			name:         "empty title",
			title:        "",
			description:  "Some description",
			date:         nil,
			experienceID: nil,
			projectID:    nil,
			wantErr:      ErrEmptyTitle,
		},
		{
			name:         "whitespace title",
			title:        "   ",
			description:  "Some description",
			date:         nil,
			experienceID: nil,
			projectID:    nil,
			wantErr:      ErrEmptyTitle,
		},
		{
			name:         "empty description",
			title:        "Achievement",
			description:  "",
			date:         nil,
			experienceID: nil,
			projectID:    nil,
			wantErr:      ErrEmptyDescription,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ach, err := NewAchievement(tt.title, tt.description, tt.date, tt.experienceID, tt.projectID)

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Equal(t, Achievement{}, ach)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.title, ach.Title)
				assert.Equal(t, tt.description, ach.Description)
			}
		})
	}
}

func TestAchievement_HasContext(t *testing.T) {
	expID := int32(1)
	projID := int32(2)

	tests := []struct {
		name         string
		experienceID *int32
		projectID    *int32
		expected     bool
	}{
		{"no context", nil, nil, false},
		{"linked to experience", &expID, nil, true},
		{"linked to project", nil, &projID, true},
		{"linked to both", &expID, &projID, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ach := Achievement{ExperienceID: tt.experienceID, ProjectID: tt.projectID}
			assert.Equal(t, tt.expected, ach.HasContext())
		})
	}
}

func TestAchievement_IsLinkedToExperience(t *testing.T) {
	expID := int32(1)

	tests := []struct {
		name         string
		experienceID *int32
		expected     bool
	}{
		{"not linked", nil, false},
		{"linked", &expID, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ach := Achievement{ExperienceID: tt.experienceID}
			assert.Equal(t, tt.expected, ach.IsLinkedToExperience())
		})
	}
}

func TestAchievement_IsLinkedToProject(t *testing.T) {
	projID := int32(2)

	tests := []struct {
		name      string
		projectID *int32
		expected  bool
	}{
		{"not linked", nil, false},
		{"linked", &projID, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ach := Achievement{ProjectID: tt.projectID}
			assert.Equal(t, tt.expected, ach.IsLinkedToProject())
		})
	}
}
