package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func createRandomTransfer(t *testing.T, toAccount Account, fromAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	updateToAmountArgs := UpdateAccountParams{
		ID:      toAccount.ID,
		Balance: toAccount.Balance + transfer.Amount,
	}

	updatedToAccount, err := testQueries.UpdateAccount(context.Background(), updateToAmountArgs)

	require.NoError(t, err)
	require.NotEmpty(t, updatedToAccount)

	updateFromAmountArgs := UpdateAccountParams{
		ID:      fromAccount.ID,
		Balance: fromAccount.Balance - transfer.Amount,
	}

	updatedFromAccount, err := testQueries.UpdateAccount(context.Background(), updateFromAmountArgs)

	require.NoError(t, err)
	require.NotEmpty(t, updatedFromAccount)

	require.Equal(t, transfer.FromAccountID, fromAccount.ID)
	require.Equal(t, transfer.ToAccountID, toAccount.ID)
	require.Equal(t, toAccount.ID, updatedToAccount.ID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, toAccount.Balance+transfer.Amount, updatedToAccount.Balance)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)

	createRandomTransfer(t, toAccount, fromAccount)
}

func TestGetTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)

	transfer := createRandomTransfer(t, toAccount, fromAccount)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, toAccount, fromAccount)
	}

	arg := ListTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        3,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
