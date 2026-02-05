package port

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// SkillRepository specifies methods for accessing and manipulating skills in a repository.
type SkillRepository interface {
	GetSkills(ctx context.Context) ([]domain.Skill, error)
	GetSkillByName(ctx context.Context, name string) (domain.Skill, error)
	CreateSkill(ctx context.Context, skill domain.Skill) error
}
