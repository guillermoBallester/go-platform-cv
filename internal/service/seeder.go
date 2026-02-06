package service

import (
	"context"
	"encoding/json"
	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"time"
)

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

	return domain.NewExperience(
		seed.CompanyName,
		seed.JobTitle,
		seed.Location,
		startDate,
		endDate,
		seed.Description,
		seed.Highlights,
	)
}

// achievementSeed represents the JSON structure for seeding achievements.
type achievementSeed struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Date         *string  `json:"date"`
	ExperienceID *int32   `json:"experience_id"`
	ProjectID    *int32   `json:"project_id"`
	Skills       []string `json:"skills"`
}

// SeedAchievements initializes the database with achievement data if no achievements exist.
func (s *CVService) SeedAchievements(ctx context.Context, data []byte) error {
	current, err := s.achievementRepo.GetAchievements(ctx)
	if err != nil {
		return err
	}
	if len(current) > 0 {
		return nil
	}

	var seeds []achievementSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		ach, err := s.parseAchievementSeed(seed)
		if err != nil {
			return err
		}

		achID, err := s.achievementRepo.CreateAchievement(ctx, ach)
		if err != nil {
			return err
		}

		// Link skills to this achievement
		for _, skillName := range seed.Skills {
			skill, err := s.skillRepo.GetSkillByName(ctx, skillName)
			if err != nil {
				continue // Skip if skill not found
			}
			if err := s.achievementRepo.AddSkillToAchievement(ctx, achID, skill.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *CVService) parseAchievementSeed(seed achievementSeed) (domain.Achievement, error) {
	var date *time.Time
	if seed.Date != nil {
		parsed, err := parseDate(*seed.Date)
		if err != nil {
			return domain.Achievement{}, err
		}
		date = &parsed
	}

	return domain.NewAchievement(
		seed.Title,
		seed.Description,
		date,
		seed.ExperienceID,
		seed.ProjectID,
	)
}

// projectSeed represents the JSON structure for seeding projects.
type projectSeed struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	StartDate   *string  `json:"start_date"`
	EndDate     *string  `json:"end_date"`
	Skills      []string `json:"skills"`
}

// SeedProjects initializes the database with project data if no projects exist.
func (s *CVService) SeedProjects(ctx context.Context, data []byte) error {
	current, err := s.projectRepo.GetProjects(ctx)
	if err != nil {
		return err
	}
	if len(current) > 0 {
		return nil
	}

	var seeds []projectSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		proj, err := s.parseProjectSeed(seed)
		if err != nil {
			return err
		}

		projID, err := s.projectRepo.CreateProject(ctx, proj)
		if err != nil {
			return err
		}

		// Link skills to this project
		for _, skillName := range seed.Skills {
			skill, err := s.skillRepo.GetSkillByName(ctx, skillName)
			if err != nil {
				continue // Skip if skill not found
			}
			if err := s.projectRepo.AddSkillToProject(ctx, projID, skill.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *CVService) parseProjectSeed(seed projectSeed) (domain.Project, error) {
	var startDate *time.Time
	if seed.StartDate != nil {
		parsed, err := parseDate(*seed.StartDate)
		if err != nil {
			return domain.Project{}, err
		}
		startDate = &parsed
	}

	var endDate *time.Time
	if seed.EndDate != nil {
		parsed, err := parseDate(*seed.EndDate)
		if err != nil {
			return domain.Project{}, err
		}
		endDate = &parsed
	}

	return domain.NewProject(
		seed.Name,
		seed.Description,
		startDate,
		endDate,
	)
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// skillSeed represents the JSON structure for seeding skills.
type skillSeed struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Proficiency int32  `json:"proficiency"`
	LogoPath    string `json:"logo_url"`
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

	var seeds []skillSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		skill, err := domain.NewSkill(seed.Name, seed.Category, seed.Proficiency, seed.LogoPath)
		if err != nil {
			return err
		}
		if err := s.skillRepo.CreateSkill(ctx, skill); err != nil {
			return err
		}
	}
	return nil
}
