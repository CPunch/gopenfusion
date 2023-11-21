package db

import (
	"context"
	"os"
	"testing"

	"github.com/bitcomplete/sqltestutil"
)

var (
	testDB *DBHandler
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	psql, err := sqltestutil.StartPostgresContainer(ctx, "latest")
	if err != nil {
		panic(err)
	}
	defer psql.Shutdown(ctx)

	testDB, err = OpenFromConnectionString("postgres", psql.ConnectionString()+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer testDB.Close()

	if err := testDB.Setup(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestDBAccount(t *testing.T) {
	if _, err := testDB.NewAccount("test", "test"); err != nil {
		t.Error(err)
	}

	// now try to retrieve account data
	acc, err := testDB.TryLogin("test", "test")
	if err != nil {
		t.Error(err)
	}

	if acc.Login != "test" {
		t.Error("account username is not test")
	}
}
