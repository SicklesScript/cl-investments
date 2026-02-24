-- name: GetHoldings :one

SELECT
COALESCE(SUM (
    CASE 
        WHEN type = 'BUY' THEN (transaction_price * shares)
        WHEN type = 'SELL' THEN -(transaction_price * shares)
        WHEN type = 'DIV' THEN (transaction_price)
        ELSE 0
    END
    ), 0)::FLOAT AS total_holdings_value
FROM transactions
WHERE username = $1;