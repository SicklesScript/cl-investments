-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

ALTER TABLE users
ADD COLUMN hashed_password TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
ALTER TABLE users
DROP COLUMN IF EXISTS hashed_password;