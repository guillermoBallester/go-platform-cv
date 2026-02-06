-- name: ListExperiences :many
SELECT * FROM experiences ORDER BY start_date DESC;

-- name: GetExperience :one
SELECT * FROM experiences WHERE id = $1;

-- name: CreateExperience :one
INSERT INTO experiences (company_name, job_title, location, start_date, end_date, description, highlights)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateExperience :one
UPDATE experiences
SET company_name = $2, job_title = $3, location = $4, start_date = $5, end_date = $6,
    description = $7, highlights = $8, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteExperience :exec
DELETE FROM experiences WHERE id = $1;

-- Skill linking
-- name: AddSkillToExperience :exec
INSERT INTO experience_skills (experience_id, skill_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: RemoveSkillFromExperience :exec
DELETE FROM experience_skills WHERE experience_id = $1 AND skill_id = $2;

-- name: ListSkillsForExperience :many
SELECT s.* FROM skills s
JOIN experience_skills es ON s.id = es.skill_id
WHERE es.experience_id = $1
ORDER BY s.category, s.name;

-- name: ListExperiencesForSkill :many
SELECT e.* FROM experiences e
JOIN experience_skills es ON e.id = es.experience_id
WHERE es.skill_id = $1
ORDER BY e.start_date DESC;

-- Project linking
-- name: AddProjectToExperience :exec
INSERT INTO experience_projects (experience_id, project_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: RemoveProjectFromExperience :exec
DELETE FROM experience_projects WHERE experience_id = $1 AND project_id = $2;

-- name: ListProjectsForExperience :many
SELECT p.* FROM projects p
JOIN experience_projects ep ON p.id = ep.project_id
WHERE ep.experience_id = $1
ORDER BY p.start_date DESC;

-- name: GetExperienceByCompanyAndTitle :one
SELECT * FROM experiences WHERE company_name = $1 AND job_title = $2;

-- name: ClearSkillsFromExperience :exec
DELETE FROM experience_skills WHERE experience_id = $1;

-- Full experience with skills (for display/RAG)
-- name: GetExperienceWithSkills :many
SELECT
    e.id, e.company_name, e.job_title, e.location, e.start_date, e.end_date,
    e.description, e.highlights, e.created_at, e.updated_at,
    s.id as skill_id, s.name as skill_name, s.category as skill_category
FROM experiences e
LEFT JOIN experience_skills es ON e.id = es.experience_id
LEFT JOIN skills s ON es.skill_id = s.id
WHERE e.id = $1
ORDER BY s.category, s.name;
