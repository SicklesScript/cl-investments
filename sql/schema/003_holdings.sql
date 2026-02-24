-- +goose Up
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker VARCHAR(6) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    transaction_price DECIMAL(10, 4) NOT NULL,
    shares DECIMAL(10, 8) NOT NULL,
    type VARCHAR(4) NOT NULL CHECK (type IN('BUY', 'SELL', 'DIV')),
    username TEXT NOT NULL,

    CONSTRAINT fk_user_transactions
        FOREIGN KEY (username)
        REFERENCES users(name) 
        ON DELETE CASCADE
);

-- Index that allows the db to look up user p
CREATE INDEX idx_transactions_user ON transactions (username);

-- +goose Down
DROP TABLE transactions;