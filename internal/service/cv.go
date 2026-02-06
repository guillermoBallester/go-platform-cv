package service

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/guillermoBallester/go-platform-cv/internal/core/port"
)

// CVService represents a service that provides methods for handling CV operations.
type CVService struct {
	skillRepo       port.SkillRepository
	expRepo         port.ExperienceRepository
	achievementRepo port.AchievementRepository
	projectRepo     port.ProjectRepository
}

// NewCVService creates a new CVService instance with the provided repositories.
func NewCVService(
	skillRepo port.SkillRepository,
	expRepo port.ExperienceRepository,
	achievementRepo port.AchievementRepository,
	projectRepo port.ProjectRepository,
) *CVService {
	return &CVService{
		skillRepo:       skillRepo,
		expRepo:         expRepo,
		achievementRepo: achievementRepo,
		projectRepo:     projectRepo,
	}
}

// GetSkills retrieves a list of skills from the repository using the provided context.
func (s *CVService) GetSkills(ctx context.Context) ([]domain.Skill, error) {
	return s.skillRepo.GetSkills(ctx)
}

// GetExperiences retrieves all experiences with their associated skills.
func (s *CVService) GetExperiences(ctx context.Context) ([]domain.Experience, error) {
	return s.expRepo.GetAllExperiencesWithSkills(ctx)
}

// GetAchievements retrieves all achievements with their associated skills.
func (s *CVService) GetAchievements(ctx context.Context) ([]domain.Achievement, error) {
	return s.achievementRepo.GetAllAchievementsWithSkills(ctx)
}

// GetProjects retrieves all projects with their associated skills.
func (s *CVService) GetProjects(ctx context.Context) ([]domain.Project, error) {
	return s.projectRepo.GetAllProjectsWithSkills(ctx)
}
