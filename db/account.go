package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/CPunch/GopenFusion/protocol"
	"github.com/blockloop/scan"
)

type Account struct {
	AccountID    int
	Login        string
	Password     string
	Selected     int
	AccountLevel int
	Created      int
	LastLogin    int
	BannedUntil  int
	BannedSince  int
	BanReason    string
}

func (db *DBHandler) NewAccount(Login, Password string) (*Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 12)
	if err != nil {
		return nil, err
	}

	row, err := db.Query("INSERT INTO Accounts (Login, Password, AccountLevel) VALUES(?, ?, ?) RETURNING *", Login, hash, protocol.CN_ACCOUNT_LEVEL__USER)
	if err != nil {
		return nil, err
	}

	var account Account
	if err := scan.Row(&account, row); err != nil {
		return nil, err
	}

	return &account, nil
}

var (
	LoginErrorInvalidID       = fmt.Errorf("Invalid Login ID!")
	LoginErrorInvalidPassword = fmt.Errorf("Invalid ID && Password combo!")
)

func (db *DBHandler) TryLogin(Login, Password string) (*Account, error) {
	row, err := db.Query("SELECT * FROM Accounts WHERE Login=?", Login)
	if err != nil {
		return nil, err
	}

	var account Account
	if err := scan.Row(&account, row); err != nil {
		return nil, LoginErrorInvalidID
	}

	if bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(Password)) != nil {
		return nil, LoginErrorInvalidPassword
	}

	// else, login was a success
	return &account, nil
}
