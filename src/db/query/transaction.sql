-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, 
    payment_id, "desc", amount
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTransactions :many
SELECT * FROM transactions LIMIT $1 OFFSET $2;

-- name: GetTransactionsById :one
SELECT * FROM transactions WHERE id = $1;

-- name: DeleteTransactionById :exec
DELETE FROM transactions WHERE id = $1;

-- name: UpdateTransactionById :one 
UPDATE transactions SET "desc" = $1 WHERE id = $2 RETURNING *;