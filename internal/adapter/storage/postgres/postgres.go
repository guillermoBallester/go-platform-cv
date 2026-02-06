package postgres

import "github.com/jackc/pgx/v5/pgxpool"

// Repositories represents a collection of different repositories to manage skills, experiences, achievements, and projects.
type Repositories struct {
	Skills       *SkillRepo
	Experiences  *ExperienceRepo
	Achievements *AchievementRepo
	Projects     *ProjectRepo
}

// NewRepositories creates a new instance of Repositories with repositories for managing skills, experiences, achievements, and projects.
func NewRepositories(db *pgxpool.Pool) *Repositories {
	queries := New(db)
	return &Repositories{
		Skills:       NewSkillRepository(queries),
		Experiences:  NewExperienceRepository(queries),
		Achievements: NewAchievementRepository(queries),
		Projects:     NewProjectRepository(queries),
	}
}
