-- name: CreateUser :one
INSERT INTO users (first_name ,last_name ,email ,password_hash ,phone_no ,dob ,is_active)
VALUES ($1, $2, $3, $4, $5, $6, COALESCE($7, TRUE))
RETURNING *;


-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsersPaginated :many
SELECT * FROM users
ORDER BY id DESC
LIMIT $1 OFFSET $2;


-- name: UpdateUser :one
UPDATE users 
SET
  email = COALESCE(sqlc.narg('email'), email),
  first_name = COALESCE(sqlc.narg('first_name'), first_name),
  last_name = COALESCE(sqlc.narg('last_name'), last_name),
  phone_no = COALESCE(sqlc.narg('phone_no'), phone_no),
  dob = COALESCE(sqlc.narg('dob'), dob),
  is_active = COALESCE(sqlc.narg('is_active'), is_active)
WHERE id = sqlc.arg('id')
RETURNING *;




-- name: DeleteUser :exec
UPDATE users 
SET
  is_deleted = TRUE
WHERE id = $1;

-- name: IsEmailTakenByOtherUser :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE email = $1 AND id <> $2
);

