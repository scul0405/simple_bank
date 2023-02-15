-- name: CreateUser :one
INSERT INTO users
    (
    username, hashed_password, full_name, email
    )
VALUES
    (
        $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
 full_name = coalesce(sqlc.narg('full_name'), full_name),
 email = coalesce(sqlc.narg('email'), email),
 hashed_password = coalesce(sqlc.narg('hashed_password'), hashed_password),
 password_changed_at = coalesce(sqlc.narg('password_changed_at'), password_changed_at)
WHERE username = sqlc.arg('username')
RETURNING *;