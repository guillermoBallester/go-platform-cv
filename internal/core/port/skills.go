package port

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// SkillRepository specifies methods for accessing and manipulating skills in a repository.
type SkillRepository interface {
	GetSkills(ctx context.Context) ([]domain.Skill, error)
}
