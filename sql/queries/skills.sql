-- name: ListSkills :many
SELECT * FROM skills ORDER BY category, name;

-- name: GetSkillByName :one
SELECT * FROM skills WHERE name = $1;

-- name: CreateSkill :one
INSERT INTO skills (name, category, proficiency, logo_url)
VALUES ($1, $2, $3, $4)
RETURNING *;