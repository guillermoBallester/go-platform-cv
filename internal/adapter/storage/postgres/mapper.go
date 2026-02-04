package postgres

import "github.com/guillermoBallester/go-platform-cv/internal/core/domain"

// toDomainSkill converts a Skill object to a domain.Skill object with modified field names.
func toDomainSkill(s Skill) domain.Skill {
	return domain.Skill{
		ID:          s.ID,
		Name:        s.Name,
		Category:    s.Category,
		Proficiency: s.Proficiency.Int32,
		LogoPath:    s.LogoUrl.String,
	}
}

// toDomainSkills converts a slice of Skill objects to a slice of domain.Skill objects using toDomainSkill for each element.
func toDomainSkills(dbSkills []Skill) []domain.Skill {
	skills := make([]domain.Skill, len(dbSkills))
	for i, s := range dbSkills {
		skills[i] = toDomainSkill(s)
	}
	return skills
}
