package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
)

type Store struct {
	*Queries // untuk import  ke controller bisa dicontoh
	db       *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransactionParams struct {
	AccountID   int64  `json:"account_id"`
	PaymentById int64  `json:"payment_by_id"`
	Desc        string `json:"desc"`
	Amount      int64  `json:"amount"`
}

type TransactionResult struct {
	Transaction Transaction `json:"transaction"`
	Account     Account     `json:"account"`
	PaymentBy   PaymentBy   `json:"payment_by"`
}

func (store *Store) TransactionTx(ctx context.Context, arg TransactionParams) (TransactionResult, error) {
	var result TransactionResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			AccountID: arg.AccountID,
			PaymentID: arg.PaymentById,
			Desc:      arg.Desc,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		acc, err := q.GetAccountById(ctx, arg.AccountID) // untuk ambil balance
		if err != nil {
			return err
		}
		currentBalanceInt, _ := strconv.Atoi(acc.Balance)

		result.Account, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.AccountID,
			Balance: strconv.Itoa(currentBalanceInt - int(arg.Amount)),
		})
		if err != nil {
			return err
		}

		result.PaymentBy, err = q.GetPaymentById(ctx, arg.PaymentById)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
