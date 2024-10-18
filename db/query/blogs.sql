-- name: ListBlogs :many
SELECT blogs.id,
    blogs.title,
    -- blogs.desc,
    blogs.user_id,
    blogs.created_at,
    blogs.updated_at,
    sqlc.embed(users)
FROM blogs
    LEFT JOIN users ON users.id = blogs.user_id
ORDER BY blogs.id
LIMIT ? OFFSET ?;