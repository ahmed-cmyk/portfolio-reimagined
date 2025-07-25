-- name: GetBlog :one
SELECT * FROM blogs
WHERE id = ? LIMIT 1;

-- name: ListBlogs :many
SELECT * FROM blogs
ORDER BY created_at;

-- name: CreateBlog :one
INSERT INTO blogs (
  title, body
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = ?
