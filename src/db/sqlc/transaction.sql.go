// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: transaction.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, 
    payment_id, "desc", amount
) VALUES (
    $1, $2, $3, $4
) RETURNING id, account_id, payment_id, created_at, "desc", amount
`

type CreateTransactionParams struct {
	AccountID int64  `json:"account_id"`
	PaymentID int64  `json:"payment_id"`
	Desc      string `json:"desc"`
	Amount    int64  `json:"amount"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.AccountID,
		arg.PaymentID,
		arg.Desc,
		arg.Amount,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PaymentID,
		&i.CreatedAt,
		&i.Desc,
		&i.Amount,
	)
	return i, err
}

const deleteTransactionById = `-- name: DeleteTransactionById :exec
DELETE FROM transactions WHERE id = $1
`

func (q *Queries) DeleteTransactionById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransactionById, id)
	return err
}

const getTransactions = `-- name: GetTransactions :many
SELECT id, account_id, payment_id, created_at, "desc", amount FROM transactions LIMIT $1 OFFSET $2
`

type GetTransactionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetTransactions(ctx context.Context, arg GetTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getTransactions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.PaymentID,
			&i.CreatedAt,
			&i.Desc,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionsById = `-- name: GetTransactionsById :one
SELECT id, account_id, payment_id, created_at, "desc", amount FROM transactions WHERE id = $1
`

func (q *Queries) GetTransactionsById(ctx context.Context, id int64) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransactionsById, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PaymentID,
		&i.CreatedAt,
		&i.Desc,
		&i.Amount,
	)
	return i, err
}

const updateTransactionById = `-- name: UpdateTransactionById :one
UPDATE transactions SET "desc" = $1 WHERE id = $2 RETURNING id, account_id, payment_id, created_at, "desc", amount
`

type UpdateTransactionByIdParams struct {
	Desc string `json:"desc"`
	ID   int64  `json:"id"`
}

func (q *Queries) UpdateTransactionById(ctx context.Context, arg UpdateTransactionByIdParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, updateTransactionById, arg.Desc, arg.ID)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PaymentID,
		&i.CreatedAt,
		&i.Desc,
		&i.Amount,
	)
	return i, err
}