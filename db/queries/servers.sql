
-- name: CreateServer :one
INSERT INTO servers (
    name,
    address,
    environment
)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetServer :one
SELECT *
FROM servers
WHERE id = $1;

-- name: GetListServers :many
SELECT *
FROM servers
ORDER BY created_at DESC, id DESC;

-- name: DeleteServer :execrows
DELETE FROM servers
WHERE id = $1;
