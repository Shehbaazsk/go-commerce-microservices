-- name: CreateUserAddress :one
INSERT INTO user_addresses (
    user_id, label, address1, address2, city, state, country, pin_code, is_default
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetUserAddressesByUserID :many
SELECT * FROM user_addresses
WHERE user_id = $1;

-- name: UpdateUserAddress :one
UPDATE user_addresses
SET
    label = $2,
    address1 = $3,
    address2 = $4,
    city = $5,
    state = $6,
    country = $7,
    pin_code = $8,
    is_default = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUserAddress :exec
DELETE FROM user_addresses WHERE id = $1;
