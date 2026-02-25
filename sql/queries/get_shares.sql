-- name: GetShares :one

SELECT shares FROM transactions
WHERE ticker = $1;