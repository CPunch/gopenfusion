package db

import (
	"database/sql"

	"github.com/CPunch/gopenfusion/core/protocol"
)

type Inventory struct {
	PlayerID  int `db:"playerid"`
	Slot      int `db:"slot"`
	ID        int `db:"id"`
	Type      int `db:"type"`
	Opt       int `db:"opt"`
	TimeLimit int `db:"timelimit"`
}

// start && end are both inclusive
func (db *DBHandler) GetPlayerInventorySlots(PlayerID int, start int, end int) ([]protocol.SItemBase, error) {
	rows, err := db.Query("SELECT Slot, Type, ID, Opt, TimeLimit FROM Inventory WHERE Slot BETWEEN $1 AND $2 AND PlayerID = $3", start, end, PlayerID)
	if err != nil {
		return nil, err
	}

	items := make([]protocol.SItemBase, end-start)
	for rows.Next() {
		var slot int
		item := protocol.SItemBase{}

		if err := rows.Scan(
			&slot, &item.IType, &item.IID, &item.IOpt, &item.ITimeLimit); err != nil {
			return nil, err
		}

		items[slot-start] = item
	}

	return items, nil
}

// start is inclusive
func (db *DBHandler) SetPlayerInventorySlots(PlayerID int, start int, items []protocol.SItemBase) error {
	return db.Transaction(func(tx *sql.Tx) error {
		// delete inventory slots
		_, err := db.Exec("DELETE FROM Inventory WHERE Slot BETWEEN $1 AND $2 AND PlayerID = $3", start, start+len(items)-1, PlayerID)
		if err != nil {
			return err
		}

		// insert inventory
		for i, item := range items {
			if item.IID != 0 {
				_, err := db.Exec("INSERT INTO Inventory (PlayerID, Slot, ID, Type, Opt, TimeLimit) VALUES ($1, $2, $3, $4, $5, $6)", PlayerID, start+i, item.IID, item.IType, item.IOpt, item.ITimeLimit)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
