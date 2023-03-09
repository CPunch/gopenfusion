package db

import (
	"github.com/CPunch/GopenFusion/protocol"
	"github.com/blockloop/scan"
)

type Inventory struct {
	PlayerID  int
	Slot      int
	ID        int
	Type      int
	Opt       int
	TimeLimit int
}

// start && end are both inclusive
func (db *DBHandler) GetPlayerInventorySlots(PlayerID int, start int, end int) ([]protocol.SItemBase, error) {
	rows, err := db.Query("SELECT * FROM Inventory WHERE Slot BETWEEN ? AND ? AND PlayerID = ?", start, end, PlayerID)
	if err != nil {
		return nil, err
	}

	var inven []Inventory
	if err := scan.Row(&inven, rows); err != nil {
		return make([]protocol.SItemBase, end-start), nil
	}

	// populate an SItemBase array
	items := make([]protocol.SItemBase, end-start)
	for _, item := range inven {
		items[item.Slot-start] = protocol.SItemBase{
			IID:        int16(item.ID),
			IType:      int16(item.Type),
			IOpt:       int32(item.Opt),
			ITimeLimit: int32(item.TimeLimit),
		}
	}
	return items, nil
}

// start is inclusive TODO: needs to be a transaction !!
func (db *DBHandler) SetPlayerInventorySlots(PlayerID int, start int, items []protocol.SItemBase) error {
	// delete inventory slots
	_, err := db.Query("DELETE FROM Inventory WHERE Slot BETWEEN ? AND ? AND PlayerID = ?", start, start+len(items)-1, PlayerID)
	if err != nil {
		return err
	}

	// insert inventory
	for i, item := range items {
		if item.IID != 0 {
			_, err := db.Query("INSERT INTO Inventory (PlayerID, Slot, ID, Type, Opt, TimeLimit) VALUES (?, ?, ?, ?, ?, ?)", PlayerID, start+i, item.IID, item.IType, item.IOpt, item.ITimeLimit)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
