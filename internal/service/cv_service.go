package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/guillermoBallester/go-platform-cv/internal/core/port"
)

// CVService represents a service that provides methods for handling CV operations.
type CVService struct {
	skillRepo port.SkillRepository
	expRepo   port.ExperienceRepository
}

// NewCVService creates a new CVService instance with the provided repositories.
func NewCVService(skillRepo port.SkillRepository, expRepo port.ExperienceRepository) *CVService {
	return &CVService{
		skillRepo: skillRepo,
		expRepo:   expRepo,
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

// SeedSkills initializes the database with skills data if no skills currently exist.
func (s *CVService) SeedSkills(ctx context.Context, data []byte) error {
	current, err := s.skillRepo.GetSkills(ctx)
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
		if err := s.skillRepo.CreateSkill(ctx, sk); err != nil {
			return err
		}
	}
	return nil
}

// experienceSeed represents the JSON structure for seeding experiences.
type experienceSeed struct {
	CompanyName string   `json:"company_name"`
	JobTitle    string   `json:"job_title"`
	Location    string   `json:"location"`
	StartDate   string   `json:"start_date"`
	EndDate     *string  `json:"end_date"`
	Description string   `json:"description"`
	Highlights  string   `json:"highlights"`
	Skills      []string `json:"skills"`
}

// SeedExperiences initializes the database with experience data if no experiences exist.
func (s *CVService) SeedExperiences(ctx context.Context, data []byte) error {
	current, err := s.expRepo.GetExperiences(ctx)
	if err != nil {
		return err
	}
	if len(current) > 0 {
		return nil
	}

	var seeds []experienceSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		exp, err := s.parseExperienceSeed(seed)
		if err != nil {
			return err
		}

		expID, err := s.expRepo.CreateExperience(ctx, exp)
		if err != nil {
			return err
		}

		// Link skills to this experience
		for _, skillName := range seed.Skills {
			skill, err := s.skillRepo.GetSkillByName(ctx, skillName)
			if err != nil {
				continue // Skip if skill not found
			}
			if err := s.expRepo.AddSkillToExperience(ctx, expID, skill.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *CVService) parseExperienceSeed(seed experienceSeed) (domain.Experience, error) {
	startDate, err := parseDate(seed.StartDate)
	if err != nil {
		return domain.Experience{}, err
	}

	var endDate *time.Time
	if seed.EndDate != nil {
		parsed, err := parseDate(*seed.EndDate)
		if err != nil {
			return domain.Experience{}, err
		}
		endDate = &parsed
	}

	return domain.Experience{
		CompanyName: seed.CompanyName,
		JobTitle:    seed.JobTitle,
		Location:    seed.Location,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: seed.Description,
		Highlights:  seed.Highlights,
	}, nil
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}
