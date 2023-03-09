package db

/*
	This database has been based off of openfusion's. Databases should be completely interchangable between
	openfusion and gopenfusion.
*/

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

type DBHandler struct {
	db *sql.DB
}

type DBQuery interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}

//go:embed migrations/new.sql
var createDBQuery string

var (
	DefaultDB *DBHandler
)

func OpenLiteDB(dbPath string) (*DBHandler, error) {
	sqliteFmt := fmt.Sprintf("%s", dbPath)

	db, err := sql.Open("sqlite", sqliteFmt)
	if err != nil {
		return nil, err
	}

	return &DBHandler{db}, nil
}

func (db *DBHandler) Query(query string, args ...any) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

func (db *DBHandler) Exec(query string, args ...any) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *DBHandler) Close() error {
	return db.db.Close()
}

func (db *DBHandler) Setup() error {
	// create db tables
	_, err := db.db.Exec(createDBQuery)
	return err
}

// calls transaction, if transaction returns a non-nil error the transaction is rolled back. otherwise the transaction is committed
func (db *DBHandler) Transaction(transaction func(*sql.Tx) error) (err error) {
	tx, err := db.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			// we panic'd ??? rollback and rethrow
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = transaction(tx)
	return
}
