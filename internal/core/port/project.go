package port

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// ProjectRepository specifies methods for accessing and manipulating projects in a repository.
type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]domain.Project, error)
	GetAllProjectsWithSkills(ctx context.Context) ([]domain.Project, error)
	GetProjectByName(ctx context.Context, name string) (domain.Project, error)
	CreateProject(ctx context.Context, proj domain.Project) (int32, error)
	UpdateProject(ctx context.Context, proj domain.Project) error
	AddSkillToProject(ctx context.Context, projectID, skillID int32) error
	ClearSkillsFromProject(ctx context.Context, projectID int32) error
}
