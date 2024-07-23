-- name: GetByID :one
SELECT * FROM clients
WHERE id = ? LIMIT 1;