-- name: GetUserById :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;
-- name: GetUserByNickName :one
SELECT *
FROM users
WHERE nickname = ?
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT ? OFFSET ?;
-- name: CountUsers :one
SELECT count(*)
FROM users;
-- name: CreateUser :execresult
INSERT INTO users (nickname, gender, age, avatar, password)
VALUES (?, ?, ?, ?, ?);
-- name: UpdateUser :execresult
UPDATE users
SET nickname = ?,
  gender = ?,
  age = ?,
  avatar = ?,
  password = ?
WHERE id = ?;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;