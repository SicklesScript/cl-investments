-- name: GetReturn :many

SELECT ticker,
SUM(CASE
    WHEN type = 'BUY' THEN shares
    WHEN type = 'SELL' THEN -shares
    ELSE 0 END)::DOUBLE PRECISION AS current_shares,

SUM(CASE
    WHEN type = 'BUY' THEN (shares * transaction_price)
    WHEN type = 'SELL' THEN -(shares * transaction_price)
    WHEN type = 'DIV' THEN -(shares * transaction_price)
    ELSE 0 END)::DOUBLE PRECISION AS cost_basis
FROM transactions
WHERE username = $1
GROUP BY ticker;