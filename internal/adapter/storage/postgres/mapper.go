package postgres

import (
	"time"

	"github.com/guillermoBallester/go-platform-cv/internal/core/domain"
)

// toDomainSkill converts a Skill object to a domain.Skill object with modified field names.
func toDomainSkill(s Skill) domain.Skill {
	return domain.Skill{
		ID:          s.ID,
		Name:        s.Name,
		Category:    s.Category,
		Proficiency: s.Proficiency.Int32,
		LogoPath:    s.LogoUrl.String,
	}
}

// toDomainSkills converts a slice of Skill objects to a slice of domain.Skill objects using toDomainSkill for each element.
func toDomainSkills(dbSkills []Skill) []domain.Skill {
	skills := make([]domain.Skill, len(dbSkills))
	for i, s := range dbSkills {
		skills[i] = toDomainSkill(s)
	}
	return skills
}

// toDomainExperience converts an Experience object to a domain.Experience object.
func toDomainExperience(e Experience) domain.Experience {
	exp := domain.Experience{
		ID:          e.ID,
		CompanyName: e.CompanyName,
		JobTitle:    e.JobTitle,
		Location:    e.Location.String,
		StartDate:   e.StartDate.Time,
		Description: e.Description,
		Highlights:  e.Highlights.String,
		CreatedAt:   e.CreatedAt.Time,
		UpdatedAt:   e.UpdatedAt.Time,
	}
	if e.EndDate.Valid {
		endDate := e.EndDate.Time
		exp.EndDate = &endDate
	}
	return exp
}

// toDomainExperiences converts a slice of Experience objects to domain.Experience slice.
func toDomainExperiences(dbExps []Experience) []domain.Experience {
	exps := make([]domain.Experience, len(dbExps))
	for i, e := range dbExps {
		exps[i] = toDomainExperience(e)
	}
	return exps
}

// toDomainProject converts a Project object to a domain.Project object.
func toDomainProject(p Project) domain.Project {
	proj := domain.Project{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		URL:         p.Url.String,
		RepoURL:     p.RepoUrl.String,
		CreatedAt:   p.CreatedAt.Time,
		UpdatedAt:   p.UpdatedAt.Time,
	}
	if p.StartDate.Valid {
		startDate := p.StartDate.Time
		proj.StartDate = &startDate
	}
	if p.EndDate.Valid {
		endDate := p.EndDate.Time
		proj.EndDate = &endDate
	}
	return proj
}

// toDomainProjects converts a slice of Project objects to domain.Project slice.
func toDomainProjects(dbProjs []Project) []domain.Project {
	projs := make([]domain.Project, len(dbProjs))
	for i, p := range dbProjs {
		projs[i] = toDomainProject(p)
	}
	return projs
}

// toDomainAchievement converts an Achievement object to a domain.Achievement object.
func toDomainAchievement(a Achievement) domain.Achievement {
	ach := domain.Achievement{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		CreatedAt:   a.CreatedAt.Time,
		UpdatedAt:   a.UpdatedAt.Time,
	}
	if a.Date.Valid {
		date := a.Date.Time
		ach.Date = &date
	}
	if a.ExperienceID.Valid {
		expID := a.ExperienceID.Int32
		ach.ExperienceID = &expID
	}
	if a.ProjectID.Valid {
		projID := a.ProjectID.Int32
		ach.ProjectID = &projID
	}
	return ach
}

// toDomainAchievements converts a slice of Achievement objects to domain.Achievement slice.
func toDomainAchievements(dbAchs []Achievement) []domain.Achievement {
	achs := make([]domain.Achievement, len(dbAchs))
	for i, a := range dbAchs {
		achs[i] = toDomainAchievement(a)
	}
	return achs
}

// Helper to convert time.Time to pgtype-compatible values for inserts
func toDate(t time.Time) time.Time {
	return t
}

func toNullableDate(t *time.Time) any {
	if t == nil {
		return nil
	}
	return *t
}
