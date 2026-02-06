package port

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// AchievementRepository specifies methods for accessing and manipulating achievements in a repository.
type AchievementRepository interface {
	GetAchievements(ctx context.Context) ([]domain.Achievement, error)
	GetAllAchievementsWithSkills(ctx context.Context) ([]domain.Achievement, error)
	GetAchievementByTitle(ctx context.Context, title string) (domain.Achievement, error)
	CreateAchievement(ctx context.Context, ach domain.Achievement) (int32, error)
	UpdateAchievement(ctx context.Context, ach domain.Achievement) error
	AddSkillToAchievement(ctx context.Context, achievementID, skillID int32) error
	ClearSkillsFromAchievement(ctx context.Context, achievementID int32) error
}
