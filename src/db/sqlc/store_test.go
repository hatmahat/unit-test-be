package db

import (
	"be-4-rs-crud/utils"
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionTx(t *testing.T) {
	store := NewStore(testDB)

	acc := createRandomAccount(t)
	pymt := createRandomPaymentName(t, "BCA")

	currentBalanceInt, _ := strconv.Atoi(acc.Balance)
	trxRandomAmount := utils.RandomMoney()

	t.Run("Success Create Trx", func(t *testing.T) {
		trx, err := store.TransactionTx(context.Background(), TransactionParams{
			AccountID:   acc.ID,
			PaymentById: pymt.ID,
			Desc:        utils.RandomString(20),
			Amount:      trxRandomAmount,
		})

		newAccBalance, _ := store.GetAccountById(context.Background(), acc.ID)
		newAccBalanceInt, _ := strconv.Atoi(newAccBalance.Balance) // untuk ngecek apakah balance beneran terupdate

		assert.NoError(t, err)
		assert.NotEmpty(t, trx.Transaction)
		assert.NotEmpty(t, trx.Account)
		assert.NotEmpty(t, trx.PaymentBy)
		assert.NotZero(t, trx.Transaction.CreatedAt)

		assert.Equal(t, trx.Transaction.AccountID, acc.ID)
		assert.Equal(t, trx.Transaction.PaymentID, pymt.ID)
		assert.Equal(t, currentBalanceInt-int(trxRandomAmount), newAccBalanceInt)

		// cek beneran terbuat di db atau nggak
		trxGet, err := store.GetTransactionsById(context.Background(), trx.Transaction.ID)
		// queries object embedded inside store jadi bisa ambil juga Get

		assert.NoError(t, err)
		assert.Equal(t, trx.Transaction, trxGet)
	})

	t.Run("No Acc Id", func(t *testing.T) {
		trx, err := store.TransactionTx(context.Background(), TransactionParams{
			AccountID:   0,
			PaymentById: pymt.ID,
			Desc:        utils.RandomString(20),
			Amount:      trxRandomAmount,
		})

		assert.Error(t, err)
		assert.Empty(t, trx)
	})

	t.Run("No Pymt Id", func(t *testing.T) {
		trx, err := store.TransactionTx(context.Background(), TransactionParams{
			AccountID:   acc.ID,
			PaymentById: 0,
			Desc:        utils.RandomString(20),
			Amount:      trxRandomAmount,
		})

		assert.Error(t, err)
		assert.Empty(t, trx)
	})
}
