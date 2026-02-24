-- name: AddTransaction :one

INSERT INTO transactions (
    id, 
    ticker, 
    transaction_date, 
    transaction_price, 
    shares, 
    type, 
    username
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;