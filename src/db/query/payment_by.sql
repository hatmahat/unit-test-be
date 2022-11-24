-- name: CreatePaymentName :one
INSERT INTO payment_by (
    payment_name 
) VALUES ($1) RETURNING *;

-- name: GetPaymentBy :many 
SELECT * FROM payment_by LIMIT $1 OFFSET $2;

-- name: GetPaymentById :one
SELECT * FROM payment_by WHERE id = $1;

-- name: DeletePaymenBy :exec
DELETE FROM payment_by WHERE id = $1;

-- name: UpdatePaymentNameById :one
UPDATE payment_by SET payment_name = $1 WHERE id = $2 RETURNING *;