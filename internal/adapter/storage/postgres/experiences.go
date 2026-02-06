package postgres

import (
	"context"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

// ExperienceRepo represents a repository for managing experiences.
type ExperienceRepo struct {
	queries *Queries
}

// NewExperienceRepository creates a new instance of ExperienceRepo.
func NewExperienceRepository(q *Queries) *ExperienceRepo {
	return &ExperienceRepo{queries: q}
}

// GetExperiences retrieves all experiences ordered by start date.
func (r *ExperienceRepo) GetExperiences(ctx context.Context) ([]domain.Experience, error) {
	dbExps, err := r.queries.ListExperiences(ctx)
	if err != nil {
		return nil, err
	}
	return toDomainExperiences(dbExps), nil
}

// GetExperienceWithSkills retrieves a single experience with its associated skills.
func (r *ExperienceRepo) GetExperienceWithSkills(ctx context.Context, id int32) (domain.Experience, error) {
	rows, err := r.queries.GetExperienceWithSkills(ctx, id)
	if err != nil {
		return domain.Experience{}, err
	}

	if len(rows) == 0 {
		return domain.Experience{}, nil
	}

	// First row contains the experience data
	first := rows[0]
	exp := domain.Experience{
		ID:          first.ID,
		CompanyName: first.CompanyName,
		JobTitle:    first.JobTitle,
		Location:    first.Location.String,
		StartDate:   first.StartDate.Time,
		Description: first.Description,
		Highlights:  first.Highlights.String,
		CreatedAt:   first.CreatedAt.Time,
		UpdatedAt:   first.UpdatedAt.Time,
	}
	if first.EndDate.Valid {
		endDate := first.EndDate.Time
		exp.EndDate = &endDate
	}

	// Collect skills from all rows
	for _, row := range rows {
		if row.SkillID.Valid {
			exp.Skills = append(exp.Skills, domain.Skill{
				ID:       row.SkillID.Int32,
				Name:     row.SkillName.String,
				Category: row.SkillCategory.String,
			})
		}
	}

	return exp, nil
}

// GetAllExperiencesWithSkills retrieves all experiences, each with their associated skills.
func (r *ExperienceRepo) GetAllExperiencesWithSkills(ctx context.Context) ([]domain.Experience, error) {
	// First get all experiences
	dbExps, err := r.queries.ListExperiences(ctx)
	if err != nil {
		return nil, err
	}

	experiences := make([]domain.Experience, 0, len(dbExps))
	for _, dbExp := range dbExps {
		// Get skills for each experience
		skills, err := r.queries.ListSkillsForExperience(ctx, dbExp.ID)
		if err != nil {
			return nil, err
		}

		exp := toDomainExperience(dbExp)
		exp.Skills = toDomainSkills(skills)
		experiences = append(experiences, exp)
	}

	return experiences, nil
}

// CreateExperience adds a new experience to the database and returns its ID.
func (r *ExperienceRepo) CreateExperience(ctx context.Context, e domain.Experience) (int32, error) {
	params := CreateExperienceParams{
		CompanyName: e.CompanyName,
		JobTitle:    e.JobTitle,
		Location:    pgtype.Text{String: e.Location, Valid: e.Location != ""},
		StartDate:   pgtype.Date{Time: e.StartDate, Valid: true},
		Description: e.Description,
		Highlights:  pgtype.Text{String: e.Highlights, Valid: e.Highlights != ""},
	}
	if e.EndDate != nil {
		params.EndDate = pgtype.Date{Time: *e.EndDate, Valid: true}
	}

	exp, err := r.queries.CreateExperience(ctx, params)
	if err != nil {
		return 0, err
	}
	return exp.ID, nil
}

// AddSkillToExperience links a skill to an experience.
func (r *ExperienceRepo) AddSkillToExperience(ctx context.Context, experienceID, skillID int32) error {
	return r.queries.AddSkillToExperience(ctx, AddSkillToExperienceParams{
		ExperienceID: experienceID,
		SkillID:      skillID,
	})
}

// GetExperienceByCompanyAndTitle retrieves an experience by company name and job title.
func (r *ExperienceRepo) GetExperienceByCompanyAndTitle(ctx context.Context, companyName, jobTitle string) (domain.Experience, error) {
	dbExp, err := r.queries.GetExperienceByCompanyAndTitle(ctx, GetExperienceByCompanyAndTitleParams{
		CompanyName: companyName,
		JobTitle:    jobTitle,
	})
	if err != nil {
		return domain.Experience{}, err
	}
	return toDomainExperience(dbExp), nil
}

// UpdateExperience updates an existing experience in the database.
func (r *ExperienceRepo) UpdateExperience(ctx context.Context, e domain.Experience) error {
	params := UpdateExperienceParams{
		ID:          e.ID,
		CompanyName: e.CompanyName,
		JobTitle:    e.JobTitle,
		Location:    pgtype.Text{String: e.Location, Valid: e.Location != ""},
		StartDate:   pgtype.Date{Time: e.StartDate, Valid: true},
		Description: e.Description,
		Highlights:  pgtype.Text{String: e.Highlights, Valid: e.Highlights != ""},
	}
	if e.EndDate != nil {
		params.EndDate = pgtype.Date{Time: *e.EndDate, Valid: true}
	}

	_, err := r.queries.UpdateExperience(ctx, params)
	return err
}

// ClearSkillsFromExperience removes all skill links from an experience.
func (r *ExperienceRepo) ClearSkillsFromExperience(ctx context.Context, experienceID int32) error {
	return r.queries.ClearSkillsFromExperience(ctx, experienceID)
}
