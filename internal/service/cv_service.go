package service

import (
	"context"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/guillermoBallester/go-platform-cv/internal/core/port"
)

// CVService represents a service that provides methods for handling skills related operations.
type CVService struct {
	repo port.SkillRepository
}

// NewCVService creates a new CVService instance with the provided SkillRepository.
func NewCVService(repo port.SkillRepository) *CVService {
	return &CVService{repo: repo}
}

// GetSkills retrieves a list of skills from the repository using the provided context.
func (s *CVService) GetSkills(ctx context.Context) ([]domain.Skill, error) {
	return s.repo.GetSkills(ctx)
}
