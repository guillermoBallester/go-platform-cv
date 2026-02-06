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
	GetExperienceByCompanyAndTitle(ctx context.Context, companyName, jobTitle string) (domain.Experience, error)
	CreateExperience(ctx context.Context, exp domain.Experience) (int32, error)
	UpdateExperience(ctx context.Context, exp domain.Experience) error
	AddSkillToExperience(ctx context.Context, experienceID, skillID int32) error
	ClearSkillsFromExperience(ctx context.Context, experienceID int32) error
}
