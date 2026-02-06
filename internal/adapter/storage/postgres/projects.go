package postgres

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

// ProjectRepo represents a repository for managing projects.
type ProjectRepo struct {
	queries *Queries
}

// NewProjectRepository creates a new instance of ProjectRepo.
func NewProjectRepository(q *Queries) *ProjectRepo {
	return &ProjectRepo{queries: q}
}

// GetProjects retrieves all projects ordered by start date.
func (r *ProjectRepo) GetProjects(ctx context.Context) ([]domain.Project, error) {
	dbProjs, err := r.queries.ListProjects(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainProjects(dbProjs), nil
}

// GetAllProjectsWithSkills retrieves all projects, each with their associated skills.
func (r *ProjectRepo) GetAllProjectsWithSkills(ctx context.Context) ([]domain.Project, error) {
	dbProjs, err := r.queries.ListProjects(ctx)
	if err != nil {
		return nil, err
	}

	projects := make([]domain.Project, 0, len(dbProjs))
	for _, dbProj := range dbProjs {
		skills, err := r.queries.ListSkillsForProject(ctx, dbProj.ID)
		if err != nil {
			return nil, err
		}

		proj := toDomainProject(dbProj)
		proj.Skills = toDomainSkills(skills)
		projects = append(projects, proj)
	}

	return projects, nil
}

// CreateProject adds a new project to the database and returns its ID.
func (r *ProjectRepo) CreateProject(ctx context.Context, p domain.Project) (int32, error) {
	params := CreateProjectParams{
		Name:        p.Name,
		Description: p.Description,
	}
	if p.StartDate != nil {
		params.StartDate = pgtype.Date{Time: *p.StartDate, Valid: true}
	}
	if p.EndDate != nil {
		params.EndDate = pgtype.Date{Time: *p.EndDate, Valid: true}
	}

	proj, err := r.queries.CreateProject(ctx, params)
	if err != nil {
		return 0, err
	}
	return proj.ID, nil
}

// AddSkillToProject links a skill to a project.
func (r *ProjectRepo) AddSkillToProject(ctx context.Context, projectID, skillID int32) error {
	return r.queries.AddSkillToProject(ctx, AddSkillToProjectParams{
		ProjectID: projectID,
		SkillID:   skillID,
	})
}
