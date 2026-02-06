package port

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// ExperienceRepository specifies methods for accessing and manipulating experiences in a repository.
type ExperienceRepository interface {
	GetExperiences(ctx context.Context) ([]domain.Experience, error)
	GetExperienceWithSkills(ctx context.Context, id int32) (domain.Experience, error)
	GetAllExperiencesWithSkills(ctx context.Context) ([]domain.Experience, error)
	CreateExperience(ctx context.Context, exp domain.Experience) (int32, error)
	AddSkillToExperience(ctx context.Context, experienceID, skillID int32) error
}
