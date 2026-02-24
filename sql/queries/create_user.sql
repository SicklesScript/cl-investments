-- name: CreateUser :exec
INSERT INTO users(id, created_at, updated_at, name, hashed_password)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2
);