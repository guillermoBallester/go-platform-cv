package service

import (
	"context"
	"encoding/json"
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

// SeedSkills initializes the database with skills data if no skills currently exist.
func (s *CVService) SeedSkills(ctx context.Context, data []byte) error {
	current, err := s.repo.GetSkills(ctx)
	if err != nil {
		return err
	}
	if len(current) > 0 {
		return nil
	}

	var skills []domain.Skill
	if err := json.Unmarshal(data, &skills); err != nil {
		return err
	}

	for _, sk := range skills {
		if err := s.repo.CreateSkill(ctx, sk); err != nil {
			return err
		}
	}
	return nil
}
