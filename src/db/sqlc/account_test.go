package db

import (
	"be-4-rs-crud/utils"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account { // biar independen per test
	arg := CreateAccountParams{
		UserName: utils.RandomUserName(),
		Balance:  utils.RandomNumString(4),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	assert.NoError(t, err) // fail the test kalau err != nil
	assert.NotEmpty(t, account)

	assert.Equal(t, arg.UserName, account.UserName)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccountById(context.Background(), acc1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc2)

	assert.Equal(t, acc1.ID, acc2.ID)
	assert.Equal(t, acc1.UserName, acc2.UserName)
	assert.Equal(t, acc1.Balance, acc2.Balance)
	assert.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: utils.RandomNumString(4),
	}

	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	assert.Equal(t, acc1.UserName, acc2.UserName)
	assert.Equal(t, arg.Balance, acc2.Balance) // dari args
	assert.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	var accounts []Account
	for i := 0; i < 10; i++ {
		createdAcc := createRandomAccount(t)
		accounts = append(accounts, createdAcc)
	}

	t.Run("When Limit or Offset > 0", func(t *testing.T) {
		arg := ListAccountsParams{
			Limit:  5,
			Offset: 5,
		}
		accs, err := testQueries.ListAccounts(context.Background(), arg)
		assert.NoError(t, err)
		assert.Len(t, accs, 5)

		for _, acc := range accs {
			assert.NotEmpty(t, acc)
		}
	})

	t.Run("When Offset < 0", func(t *testing.T) {
		arg := ListAccountsParams{
			Limit:  5,
			Offset: -5,
		}
		accs, err := testQueries.ListAccounts(context.Background(), arg)
		assert.NotEmpty(t, err)
		assert.Len(t, accs, 0)

		for _, acc := range accs {
			assert.NotEmpty(t, acc)
		}
	})
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)
	assert.NotEmpty(t, acc)

	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	assert.NoError(t, err)
}
