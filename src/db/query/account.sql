-- name: CreateAccount :one
INSERT INTO accounts (
    user_name,
    balance
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAccountById :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;