package db

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/georgysavva/scany/v2/sqlscan"
)

type Account struct {
	AccountID    int    `db:"accountid"`
	Login        string `db:"login"`
	Password     string `db:"password"`
	Selected     int    `db:"selected"`
	AccountLevel int    `db:"accountlevel"`
	Created      int    `db:"created"`
	LastLogin    int    `db:"lastlogin"`
	BannedUntil  int    `db:"banneduntil"`
	BannedSince  int    `db:"bannedsince"`
	BanReason    string `db:"banreason"`
}

func (db *DBHandler) NewAccount(Login, Password string) (*Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 12)
	if err != nil {
		return nil, err
	}

	row, err := db.Query("INSERT INTO Accounts (Login, Password, AccountLevel) VALUES($1, $2, $3) RETURNING *", Login, hash, protocol.CN_ACCOUNT_LEVEL__USER)
	if err != nil {
		return nil, err
	}

	var account Account
	row.Next()
	if err := sqlscan.ScanRow(&account, row); err != nil {
		return nil, err
	}

	return &account, nil
}

var (
	LoginErrorInvalidID       = fmt.Errorf("Invalid Login ID!")
	LoginErrorInvalidPassword = fmt.Errorf("Invalid ID && Password combo!")
)

func (db *DBHandler) TryLogin(Login, Password string) (*Account, error) {
	row, err := db.Query("SELECT * FROM Accounts WHERE Login=$1", Login)
	if err != nil {
		return nil, err
	}

	var account Account
	row.Next()
	if err := sqlscan.ScanRow(&account, row); err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, LoginErrorInvalidID
	}

	if bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(Password)) != nil {
		return nil, LoginErrorInvalidPassword
	}

	// else, login was a success
	return &account, nil
}
