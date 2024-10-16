-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByNickName :one
SELECT * FROM users
WHERE nickname = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
  nickname, gender, age, avatar
) VALUES (
  ?, ?, ?, ?
);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;