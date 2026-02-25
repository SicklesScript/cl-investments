-- name: GetAll :many

SELECT * FROM transactions
WHERE username = $1;