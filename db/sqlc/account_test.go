package db

import (
	"context"
	"testing"
	"time"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		ID:       util.RandomString(24),
		Owner:    util.RandomOwner(),
		Balance:  int64(util.RandomInt(0, 1000)),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createTestAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
	require.WithinDuration(t, account.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createTestAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: int64(util.RandomInt(0, 1000)),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
}
