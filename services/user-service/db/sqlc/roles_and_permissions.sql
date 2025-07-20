-- name: CreateRole :one
INSERT INTO roles (name, description, is_active)
VALUES ($1, $2, COALESCE($3, TRUE))
RETURNING *;

-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = $1;

-- name: GetRoleByName :one
SELECT * FROM roles WHERE name = $1;

-- name: ListRolesPaginatedSearch :many
SELECT id,name,description,is_active,created_at,updated_at,count(*) OVER() as total_count
FROM roles
WHERE
  (LOWER(name) LIKE LOWER('%' || sqlc.arg('search_by') || '%') OR sqlc.arg('search_by') = '')
ORDER BY
  CASE 
    WHEN sqlc.arg('order_by') = 'name' AND sqlc.arg('order_dir') = 'asc' THEN name 
  END ASC,
  CASE 
    WHEN sqlc.arg('order_by') = 'name' AND sqlc.arg('order_dir') = 'desc' THEN name 
  END DESC,
  id DESC -- Default sort
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');


-- name: UpdateRole :one
UPDATE roles SET
  name = COALESCE(sqlc.narg('name'), name),
  description = COALESCE(sqlc.narg('description'), description),
  is_active = COALESCE(sqlc.narg('is_active'), is_active),
  updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;


-- name: CreatePermission :one
INSERT INTO permissions (name)
VALUES ($1)
RETURNING *;

-- name: GetPermissionByID :one
SELECT * FROM permissions WHERE id = $1;

-- name: GetPermissionByName :one
SELECT * FROM permissions WHERE name = $1;

-- name: ListPermissionsPaginatedSearch :many
SELECT id,name,count(*) OVER() as total_count
FROM permissions
WHERE
  (LOWER(name) LIKE LOWER(sqlc.arg('search_by')) OR sqlc.arg('search_by') = '')
ORDER BY
  CASE WHEN sqlc.narg('order_by_name_asc')::bool THEN name END ASC,
  CASE WHEN sqlc.narg('order_by_name_desc')::bool THEN name END DESC,
  id DESC -- Default sort
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: DeletePermission :exec
DELETE FROM permissions WHERE id = $1;

-- name: AssignPermissionToRole :exec
INSERT INTO role_permissions (role_id, permission_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemovePermissionFromRole :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2;

-- name: ListPermissionsByRole :many
SELECT p.* FROM permissions p
JOIN role_permissions rp ON rp.permission_id = p.id
WHERE rp.role_id = $1;
