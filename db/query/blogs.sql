-- name: ListBlogs :many
SELECT blogs.id,
    blogs.title,
    blogs.desc,
    blogs.user_id,
    blogs.created_at,
    blogs.updated_at,
    sqlc.embed(users) -- sqlc.embed(connects)
FROM blogs
    LEFT JOIN users ON users.id = blogs.user_id -- LEFT JOIN connects ON users.id = connects.user_id
WHERE (
        blogs.desc LIKE sqlc.narg('desc')
        OR sqlc.narg('desc') IS NULL
    )
ORDER BY blogs.id DESC
LIMIT ? OFFSET ?;
-- name: CountBlogs :one
SELECT Count(*)
FROM blogs
WHERE blogs.deleted_at IS NULL
    AND (
        blogs.desc LIKE sqlc.narg('desc')
        OR sqlc.narg('desc') IS NULL
    );