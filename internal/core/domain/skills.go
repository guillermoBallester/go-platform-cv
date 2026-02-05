package domain

import "strings"

// Skill represents a specific skill with an ID, name, category, and proficiency level.
type Skill struct {
	ID          int32
	Name        string
	Category    string
	Proficiency int32
	LogoPath    string `json:"logo_url"`
}

// NewSkill creates a validated Skill. Returns error if validation fails.
func NewSkill(name, category string, proficiency int32, logoPath string) (Skill, error) {
	s := Skill{
		Name:        strings.TrimSpace(name),
		Category:    strings.TrimSpace(category),
		Proficiency: proficiency,
		LogoPath:    strings.TrimSpace(logoPath),
	}

	if err := s.Validate(); err != nil {
		return Skill{}, err
	}

	return s, nil
}

// Validate checks all business rules for Skill.
func (s Skill) Validate() error {
	if s.Name == "" {
		return ErrEmptyName
	}
	if s.Category == "" {
		return ErrEmptyCategory
	}
	if s.Proficiency < 0 || s.Proficiency > 100 {
		return ErrInvalidProficiency
	}
	return nil
}

// IsExpert returns true if proficiency is 80 or above.
func (s Skill) IsExpert() bool {
	return s.Proficiency >= 80
}

// IsProficient returns true if proficiency is 60 or above.
func (s Skill) IsProficient() bool {
	return s.Proficiency >= 60
}
