-- name: ListSkills :many
SELECT * FROM skills ORDER BY category, name;

-- name: CreateSkill :one
INSERT INTO skills (name, category, proficiency)
VALUES ($1, $2, $3)
RETURNING *;