package postgres

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

// AchievementRepo represents a repository for managing achievements.
type AchievementRepo struct {
	queries *Queries
}

// NewAchievementRepository creates a new instance of AchievementRepo.
func NewAchievementRepository(q *Queries) *AchievementRepo {
	return &AchievementRepo{queries: q}
}

// GetAchievements retrieves all achievements ordered by date.
func (r *AchievementRepo) GetAchievements(ctx context.Context) ([]domain.Achievement, error) {
	dbAchs, err := r.queries.ListAchievements(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainAchievements(dbAchs), nil
}

// GetAllAchievementsWithSkills retrieves all achievements, each with their associated skills.
func (r *AchievementRepo) GetAllAchievementsWithSkills(ctx context.Context) ([]domain.Achievement, error) {
	dbAchs, err := r.queries.ListAchievements(ctx)
	if err != nil {
		return nil, err
	}

	achievements := make([]domain.Achievement, 0, len(dbAchs))
	for _, dbAch := range dbAchs {
		skills, err := r.queries.ListSkillsForAchievement(ctx, dbAch.ID)
		if err != nil {
			return nil, err
		}

		ach := toDomainAchievement(dbAch)
		ach.Skills = toDomainSkills(skills)
		achievements = append(achievements, ach)
	}

	return achievements, nil
}

// CreateAchievement adds a new achievement to the database and returns its ID.
func (r *AchievementRepo) CreateAchievement(ctx context.Context, a domain.Achievement) (int32, error) {
	params := CreateAchievementParams{
		Title:       a.Title,
		Description: a.Description,
	}
	if a.Date != nil {
		params.Date = pgtype.Date{Time: *a.Date, Valid: true}
	}
	if a.ExperienceID != nil {
		params.ExperienceID = pgtype.Int4{Int32: *a.ExperienceID, Valid: true}
	}
	if a.ProjectID != nil {
		params.ProjectID = pgtype.Int4{Int32: *a.ProjectID, Valid: true}
	}

	ach, err := r.queries.CreateAchievement(ctx, params)
	if err != nil {
		return 0, err
	}
	return ach.ID, nil
}

// AddSkillToAchievement links a skill to an achievement.
func (r *AchievementRepo) AddSkillToAchievement(ctx context.Context, achievementID, skillID int32) error {
	return r.queries.AddSkillToAchievement(ctx, AddSkillToAchievementParams{
		AchievementID: achievementID,
		SkillID:       skillID,
	})
}

// GetAchievementByTitle retrieves an achievement by its title.
func (r *AchievementRepo) GetAchievementByTitle(ctx context.Context, title string) (domain.Achievement, error) {
	dbAch, err := r.queries.GetAchievementByTitle(ctx, title)
	if err != nil {
		return domain.Achievement{}, err
	}
	return toDomainAchievement(dbAch), nil
}

// UpdateAchievement updates an existing achievement in the database.
func (r *AchievementRepo) UpdateAchievement(ctx context.Context, a domain.Achievement) error {
	params := UpdateAchievementParams{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
	}
	if a.Date != nil {
		params.Date = pgtype.Date{Time: *a.Date, Valid: true}
	}
	if a.ExperienceID != nil {
		params.ExperienceID = pgtype.Int4{Int32: *a.ExperienceID, Valid: true}
	}
	if a.ProjectID != nil {
		params.ProjectID = pgtype.Int4{Int32: *a.ProjectID, Valid: true}
	}

	_, err := r.queries.UpdateAchievement(ctx, params)
	return err
}

// ClearSkillsFromAchievement removes all skill links from an achievement.
func (r *AchievementRepo) ClearSkillsFromAchievement(ctx context.Context, achievementID int32) error {
	return r.queries.ClearSkillsFromAchievement(ctx, achievementID)
}
