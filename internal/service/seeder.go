package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/guillermoBallester/go-platform-cv/internal/core/port"
	"github.com/guillermoBallester/go-platform-cv/sql/data"
	"log"
	"time"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

// SeedService manages seeding data by providing methods to interact with different repository types.
type SeedService struct {
	skillRepo       port.SkillRepository
	expRepo         port.ExperienceRepository
	achievementRepo port.AchievementRepository
	projectRepo     port.ProjectRepository
}

func NewSeedService(
	skillRepo port.SkillRepository,
	expRepo port.ExperienceRepository,
	achievementRepo port.AchievementRepository,
	projectRepo port.ProjectRepository,
) *SeedService {
	return &SeedService{
		skillRepo:       skillRepo,
		expRepo:         expRepo,
		achievementRepo: achievementRepo,
		projectRepo:     projectRepo,
	}
}

func (s *SeedService) Run(ctx context.Context) error {
	tasks := []struct {
		name string
		fn   func(context.Context, []byte) error
		data []byte
	}{
		{"skills", s.SeedSkills, data.SkillsJSON},
		{"experiences", s.SeedExperiences, data.ExperiencesJSON},
		{"achievements", s.SeedAchievements, data.AchievementsJSON},
		{"projects", s.SeedProjects, data.ProjectsJSON},
	}

	for _, task := range tasks {
		if err := task.fn(ctx, task.data); err != nil {
			// We log the error but continue to the next task
			log.Printf("Warning: could not seed %s: %v", task.name, err)
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

// SeedExperiences upserts experience data - creates new experiences or updates existing ones.
func (s *SeedService) SeedExperiences(ctx context.Context, data []byte) error {
	var seeds []experienceSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		exp, err := s.parseExperienceSeed(seed)
		if err != nil {
			return err
		}

		existing, err := s.expRepo.GetExperienceByCompanyAndTitle(ctx, seed.CompanyName, seed.JobTitle)
		if errors.Is(err, pgx.ErrNoRows) {
			// Create new experience
			expID, err := s.expRepo.CreateExperience(ctx, exp)
			if err != nil {
				return err
			}
			if err := s.linkSkillsToExperience(ctx, expID, seed.Skills); err != nil {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}

		// Update existing experience
		exp.ID = existing.ID
		if err := s.expRepo.UpdateExperience(ctx, exp); err != nil {
			return err
		}

		// Re-link skills (clear + re-add)
		if err := s.expRepo.ClearSkillsFromExperience(ctx, existing.ID); err != nil {
			return err
		}
		if err := s.linkSkillsToExperience(ctx, existing.ID, seed.Skills); err != nil {
			return err
		}
	}
	return nil
}

// linkSkillsToExperience links skills by name to an experience.
func (s *SeedService) linkSkillsToExperience(ctx context.Context, expID int32, skillNames []string) error {
	for _, skillName := range skillNames {
		skill, err := s.skillRepo.GetSkillByName(ctx, skillName)
		if err != nil {
			continue // Skip if skill not found
		}
		if err := s.expRepo.AddSkillToExperience(ctx, expID, skill.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s *SeedService) parseExperienceSeed(seed experienceSeed) (domain.Experience, error) {
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

// SeedAchievements upserts achievement data - creates new achievements or updates existing ones.
func (s *SeedService) SeedAchievements(ctx context.Context, data []byte) error {
	var seeds []achievementSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		ach, err := s.parseAchievementSeed(seed)
		if err != nil {
			return err
		}

		existing, err := s.achievementRepo.GetAchievementByTitle(ctx, seed.Title)
		if errors.Is(err, pgx.ErrNoRows) {
			// Create new achievement
			achID, err := s.achievementRepo.CreateAchievement(ctx, ach)
			if err != nil {
				return err
			}
			if err := s.linkSkillsToAchievement(ctx, achID, seed.Skills); err != nil {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}

		// Update existing achievement
		ach.ID = existing.ID
		if err := s.achievementRepo.UpdateAchievement(ctx, ach); err != nil {
			return err
		}

		// Re-link skills (clear + re-add)
		if err := s.achievementRepo.ClearSkillsFromAchievement(ctx, existing.ID); err != nil {
			return err
		}
		if err := s.linkSkillsToAchievement(ctx, existing.ID, seed.Skills); err != nil {
			return err
		}
	}
	return nil
}

// linkSkillsToAchievement links skills by name to an achievement.
func (s *SeedService) linkSkillsToAchievement(ctx context.Context, achID int32, skillNames []string) error {
	for _, skillName := range skillNames {
		skill, err := s.skillRepo.GetSkillByName(ctx, skillName)
		if err != nil {
			continue // Skip if skill not found
		}
		if err := s.achievementRepo.AddSkillToAchievement(ctx, achID, skill.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s *SeedService) parseAchievementSeed(seed achievementSeed) (domain.Achievement, error) {
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

// SeedProjects upserts project data - creates new projects or updates existing ones.
func (s *SeedService) SeedProjects(ctx context.Context, data []byte) error {
	var seeds []projectSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		proj, err := s.parseProjectSeed(seed)
		if err != nil {
			return err
		}

		existing, err := s.projectRepo.GetProjectByName(ctx, seed.Name)
		if errors.Is(err, pgx.ErrNoRows) {
			// Create new project
			projID, err := s.projectRepo.CreateProject(ctx, proj)
			if err != nil {
				return err
			}
			if err := s.linkSkillsToProject(ctx, projID, seed.Skills); err != nil {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}

		// Update existing project
		proj.ID = existing.ID
		if err := s.projectRepo.UpdateProject(ctx, proj); err != nil {
			return err
		}

		// Re-link skills (clear + re-add)
		if err := s.projectRepo.ClearSkillsFromProject(ctx, existing.ID); err != nil {
			return err
		}
		if err := s.linkSkillsToProject(ctx, existing.ID, seed.Skills); err != nil {
			return err
		}
	}
	return nil
}

// linkSkillsToProject links skills by name to a project.
func (s *SeedService) linkSkillsToProject(ctx context.Context, projID int32, skillNames []string) error {
	for _, skillName := range skillNames {
		skill, err := s.skillRepo.GetSkillByName(ctx, skillName)
		if err != nil {
			continue // Skip if skill not found
		}
		if err := s.projectRepo.AddSkillToProject(ctx, projID, skill.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s *SeedService) parseProjectSeed(seed projectSeed) (domain.Project, error) {
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

// SeedSkills upserts skills data - creates new skills or updates existing ones.
func (s *SeedService) SeedSkills(ctx context.Context, data []byte) error {
	var seeds []skillSeed
	if err := json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	for _, seed := range seeds {
		skill, err := domain.NewSkill(seed.Name, seed.Category, seed.Proficiency, seed.LogoPath)
		if err != nil {
			return err
		}

		existing, err := s.skillRepo.GetSkillByName(ctx, seed.Name)
		if errors.Is(err, pgx.ErrNoRows) {
			// Create new skill
			if err := s.skillRepo.CreateSkill(ctx, skill); err != nil {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}

		// Update existing skill
		skill.ID = existing.ID
		if err := s.skillRepo.UpdateSkill(ctx, skill); err != nil {
			return err
		}
	}
	return nil
}
