package test

import (
	// "be-4-rs-crud/src/db"
	db "be-4-rs-crud/src/db/sqlc"
	sqlc "be-4-rs-crud/src/db/sqlc"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

type Store struct {
	store *db.Store
}

func TestCreateAccount(t *testing.T) {
	store := sqlc.NewStore(sqlc.TestDB) // kurang connection
	//acc := LoadPackage("../src/db/sqlc/account.sql.go")
	arg := sqlc.CreateAccountParams{
		UserName: "tom",
		Balance:  "100",
	}

	account, err := store.CreateAccount(context.Background(), arg)
	require.NoError(t, err) // fail the test kalau err != nil
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserName, account.UserName)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func LoadPackage(path string) *packages.Package {
	const mode = packages.NeedTypes |
		packages.NeedName |
		packages.NeedSyntax |
		packages.NeedFiles |
		packages.NeedTypesInfo |
		packages.NeedTypesInfo |
		packages.NeedModule

	cfg := &packages.Config{Mode: mode}

	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		panic(err)
	}

	return pkgs[0]
}
