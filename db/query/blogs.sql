-- name: ListBlogs :many
SELECT blogs.id,
    blogs.title,
    blogs.desc,
    blogs.user_id,
    blogs.created_at,
    blogs.updated_at,
    sqlc.embed(users)
FROM blogs
    LEFT JOIN users ON users.id = blogs.user_id -- WHERE blogs.desc LIKE sqlc.narg('desc')
ORDER BY blogs.id DESC
LIMIT ? OFFSET ?;
-- name: CountBlogs :one
SELECT Count(*)
FROM blogs
WHERE blogs.deleted_at IS NULL
    AND blogs.desc LIKE sqlc.narg('desc');