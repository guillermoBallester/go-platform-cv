package postgres

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// SkillRepo represents a repository for managing skills, using the provided Queries struct.
type SkillRepo struct {
	queries *Queries
}

// NewSkillRepository creates a new instance of SkillRepo initialized with the provided Queries struct.
func NewSkillRepository(q *Queries) *SkillRepo {
	return &SkillRepo{queries: q}
}

// GetSkills retrieves a list of domain.Skill objects by querying the database and converting them
// to the appropriate model structure.
func (r *SkillRepo) GetSkills(ctx context.Context) ([]domain.Skill, error) {
	dbSkills, err := r.queries.ListSkills(ctx)
	if err != nil {
		return nil, err
	}

	var skills []domain.Skill
	for _, s := range dbSkills {
		skills = append(skills, domain.Skill{
			ID:          s.ID,
			Name:        s.Name,
			Category:    s.Category,
			Proficiency: s.Proficiency.Int32,
		})
	}
	return skills, nil
}
