--name :AuthenticateUser :one
SELECT id
FROM users
WHERE email = $1 AND password_hash = $2
LIMIT 1
RETURNING id;