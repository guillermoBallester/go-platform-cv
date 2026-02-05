package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSkill(t *testing.T) {
	tests := []struct {
		name        string
		skillName   string
		category    string
		proficiency int32
		logoPath    string
		wantErr     error
	}{
		{
			name:        "valid skill",
			skillName:   "Go",
			category:    "Backend",
			proficiency: 90,
			logoPath:    "/assets/logos/go.svg",
			wantErr:     nil,
		},
		{
			name:        "valid skill with zero proficiency",
			skillName:   "Rust",
			category:    "Backend",
			proficiency: 0,
			logoPath:    "",
			wantErr:     nil,
		},
		{
			name:        "valid skill with max proficiency",
			skillName:   "Python",
			category:    "Backend",
			proficiency: 100,
			logoPath:    "",
			wantErr:     nil,
		},
		{
			name:        "empty name",
			skillName:   "",
			category:    "Backend",
			proficiency: 50,
			logoPath:    "",
			wantErr:     ErrEmptyName,
		},
		{
			name:        "whitespace name",
			skillName:   "   ",
			category:    "Backend",
			proficiency: 50,
			logoPath:    "",
			wantErr:     ErrEmptyName,
		},
		{
			name:        "empty category",
			skillName:   "Go",
			category:    "",
			proficiency: 50,
			logoPath:    "",
			wantErr:     ErrEmptyCategory,
		},
		{
			name:        "negative proficiency",
			skillName:   "Go",
			category:    "Backend",
			proficiency: -1,
			logoPath:    "",
			wantErr:     ErrInvalidProficiency,
		},
		{
			name:        "proficiency over 100",
			skillName:   "Go",
			category:    "Backend",
			proficiency: 101,
			logoPath:    "",
			wantErr:     ErrInvalidProficiency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill, err := NewSkill(tt.skillName, tt.category, tt.proficiency, tt.logoPath)

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Equal(t, Skill{}, skill)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.skillName, skill.Name)
				assert.Equal(t, tt.category, skill.Category)
				assert.Equal(t, tt.proficiency, skill.Proficiency)
			}
		})
	}
}

func TestSkill_IsExpert(t *testing.T) {
	tests := []struct {
		name        string
		proficiency int32
		expected    bool
	}{
		{"79 is not expert", 79, false},
		{"80 is expert", 80, true},
		{"100 is expert", 100, true},
		{"0 is not expert", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := Skill{Proficiency: tt.proficiency}
			assert.Equal(t, tt.expected, skill.IsExpert())
		})
	}
}

func TestSkill_IsProficient(t *testing.T) {
	tests := []struct {
		name        string
		proficiency int32
		expected    bool
	}{
		{"59 is not proficient", 59, false},
		{"60 is proficient", 60, true},
		{"80 is proficient", 80, true},
		{"0 is not proficient", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skill := Skill{Proficiency: tt.proficiency}
			assert.Equal(t, tt.expected, skill.IsProficient())
		})
	}
}
