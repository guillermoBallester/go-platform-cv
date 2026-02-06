-- name: ListAchievements :many
SELECT * FROM achievements ORDER BY date DESC NULLS LAST;

-- name: GetAchievement :one
SELECT * FROM achievements WHERE id = $1;

-- name: CreateAchievement :one
INSERT INTO achievements (title, description, date, experience_id, project_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateAchievement :one
UPDATE achievements
SET title = $2, description = $3, date = $4, experience_id = $5,
    project_id = $6, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAchievement :exec
DELETE FROM achievements WHERE id = $1;

-- Skill linking
-- name: AddSkillToAchievement :exec
INSERT INTO achievement_skills (achievement_id, skill_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: RemoveSkillFromAchievement :exec
DELETE FROM achievement_skills WHERE achievement_id = $1 AND skill_id = $2;

-- name: ListSkillsForAchievement :many
SELECT s.* FROM skills s
JOIN achievement_skills aks ON s.id = aks.skill_id
WHERE aks.achievement_id = $1
ORDER BY s.category, s.name;

-- name: ListAchievementsForSkill :many
SELECT a.* FROM achievements a
JOIN achievement_skills aks ON a.id = aks.achievement_id
WHERE aks.skill_id = $1
ORDER BY a.date DESC NULLS LAST;

-- name: GetAchievementByTitle :one
SELECT * FROM achievements WHERE title = $1;

-- name: ClearSkillsFromAchievement :exec
DELETE FROM achievement_skills WHERE achievement_id = $1;

-- Filter by context
-- name: ListAchievementsForExperience :many
SELECT * FROM achievements WHERE experience_id = $1 ORDER BY date DESC NULLS LAST;

-- name: ListAchievementsForProject :many
SELECT * FROM achievements WHERE project_id = $1 ORDER BY date DESC NULLS LAST;

-- Full achievement with skills (for display/RAG)
-- name: GetAchievementWithSkills :many
SELECT
    a.id, a.title, a.description, a.date, a.experience_id, a.project_id,
    a.created_at, a.updated_at,
    s.id as skill_id, s.name as skill_name, s.category as skill_category
FROM achievements a
LEFT JOIN achievement_skills aks ON a.id = aks.achievement_id
LEFT JOIN skills s ON aks.skill_id = s.id
WHERE a.id = $1
ORDER BY s.category, s.name;

-- RAG context query: Get all achievements with their related experience/project context
-- name: ListAchievementsWithContext :many
SELECT
    a.id, a.title, a.description, a.date,
    e.company_name, e.job_title,
    p.name as project_name
FROM achievements a
LEFT JOIN experiences e ON a.experience_id = e.id
LEFT JOIN projects p ON a.project_id = p.id
ORDER BY a.date DESC NULLS LAST;
