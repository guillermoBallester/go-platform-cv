package postgres

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
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

	return toDomainSkills(dbSkills), nil
}

// CreateSkill adds a new skill in the database using the provided context and domain.Skill object.
func (r *SkillRepo) CreateSkill(ctx context.Context, s domain.Skill) error {
	_, err := r.queries.CreateSkill(ctx, CreateSkillParams{
		Name:        s.Name,
		Category:    s.Category,
		Proficiency: pgtype.Int4{Int32: s.Proficiency, Valid: true},
		LogoUrl:     pgtype.Text{String: s.LogoPath, Valid: s.LogoPath != ""},
	})
	return err
}
