-- name: CreateRole :one
INSERT INTO roles (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: ListRoles :many
SELECT * FROM roles WHERE is_active = true;