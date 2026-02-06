package service

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/adapter/storage/postgres"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// CVService represents a service that provides methods for handling CvService operations.
type CVService struct {
	dbRepositories postgres.Repositories
}

// NewCVService creates a new CVService instance with the provided repositories.
func NewCVService(
	dbRepositories postgres.Repositories,
) *CVService {
	return &CVService{
		dbRepositories: dbRepositories,
	}
}

// GetSkills retrieves a list of skills from the repository using the provided context.
func (s *CVService) GetSkills(ctx context.Context) ([]domain.Skill, error) {
	return s.dbRepositories.Skills.GetSkills(ctx)
}

// GetExperiences retrieves all experiences with their associated skills.
func (s *CVService) GetExperiences(ctx context.Context) ([]domain.Experience, error) {
	return s.dbRepositories.Experiences.GetAllExperiencesWithSkills(ctx)
}

// GetAchievements retrieves all achievements with their associated skills.
func (s *CVService) GetAchievements(ctx context.Context) ([]domain.Achievement, error) {
	return s.dbRepositories.Achievements.GetAllAchievementsWithSkills(ctx)
}

// GetProjects retrieves all projects with their associated skills.
func (s *CVService) GetProjects(ctx context.Context) ([]domain.Project, error) {
	return s.dbRepositories.Projects.GetAllProjectsWithSkills(ctx)
}
