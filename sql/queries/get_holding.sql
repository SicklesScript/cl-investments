-- name: GetHolding :one

SELECT 
COALESCE(SUM (
    CASE 
        WHEN type = 'BUY' THEN shares
        WHEN type = 'SELL' THEN -shares
        ELSE 0
    END
    ), 0)::FLOAT AS total_holding
FROM transactions
WHERE ticker = $1 AND username = $2;