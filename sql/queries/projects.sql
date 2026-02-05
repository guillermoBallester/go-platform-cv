-- name: ListProjects :many
SELECT * FROM projects ORDER BY start_date DESC NULLS LAST;

-- name: GetProject :one
SELECT * FROM projects WHERE id = $1;

-- name: CreateProject :one
INSERT INTO projects (name, description, start_date, end_date)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET name = $2, description = $3, start_date = $4, end_date = $5, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;

-- Skill linking
-- name: AddSkillToProject :exec
INSERT INTO project_skills (project_id, skill_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: RemoveSkillFromProject :exec
DELETE FROM project_skills WHERE project_id = $1 AND skill_id = $2;

-- name: ListSkillsForProject :many
SELECT s.* FROM skills s
JOIN project_skills ps ON s.id = ps.skill_id
WHERE ps.project_id = $1
ORDER BY s.category, s.name;

-- name: ListProjectsForSkill :many
SELECT p.* FROM projects p
JOIN project_skills ps ON p.id = ps.project_id
WHERE ps.skill_id = $1
ORDER BY p.start_date DESC NULLS LAST;

-- Experience linking
-- name: ListExperiencesForProject :many
SELECT e.* FROM experiences e
JOIN experience_projects ep ON e.id = ep.experience_id
WHERE ep.project_id = $1
ORDER BY e.start_date DESC;

-- Full project with skills (for display/RAG)
-- name: GetProjectWithSkills :many
SELECT
    p.id, p.name, p.description, p.start_date, p.end_date,
    p.created_at, p.updated_at,
    s.id as skill_id, s.name as skill_name, s.category as skill_category
FROM projects p
LEFT JOIN project_skills ps ON p.id = ps.project_id
LEFT JOIN skills s ON ps.skill_id = s.id
WHERE p.id = $1
ORDER BY s.category, s.name;
