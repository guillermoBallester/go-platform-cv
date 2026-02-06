-- name: ListSkills :many
SELECT * FROM skills ORDER BY category, name;

-- name: GetSkillByName :one
SELECT * FROM skills WHERE name = $1;

-- name: CreateSkill :one
INSERT INTO skills (name, category, proficiency, logo_url)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateSkill :one
UPDATE skills SET category = $2, proficiency = $3, logo_url = $4, updated_at = NOW()
WHERE id = $1 RETURNING *;