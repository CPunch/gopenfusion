package db

import (
	"database/sql"

	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/config"
	"github.com/blockloop/scan"
)

type Player struct {
	PlayerID           int
	AccountID          int
	AccountLevel       int
	Slot               int
	PCStyle            protocol.SPCStyle
	PCStyle2           protocol.SPCStyle2
	EquippedNanos      [3]int
	Nanos              [config.NANO_COUNT]protocol.SNano
	Equip              [config.AEQUIP_COUNT]protocol.SItemBase
	Inven              [config.AINVEN_COUNT]protocol.SItemBase
	Bank               [config.ABANK_COUNT]protocol.SItemBase
	SkywayLocationFlag []byte
	FirstUseFlag       []byte
	Quests             []byte
	HP                 int
	Level              int
	Taros              int
	FusionMatter       int
	Mentor             int
	X, Y, Z            int
	Angle              int
	BatteryN           int
	BatteryW           int
	WarpLocationFlag   int
	ActiveNanoSlotNum  int
	Fatigue            int
	CurrentMissionID   int
	IPCState           int8
}

// returns PlayerID, error
func (db *DBHandler) NewPlayer(AccountID int, FirstName, LastName string, slot int) (int, error) {
	nameCheck := 1 // for now, we approve all names
	QuestFlag := make([]byte, 128)
	SkywayLocationFlag := make([]byte, 16)
	FirstUseFlag := make([]byte, 16)

	var PlayerID int
	if err := db.Transaction(func(tx *sql.Tx) error {
		// create player
		rows, err := tx.Query(
			"INSERT INTO Players (AccountID, Slot, FirstName, LastName, XCoordinate, YCoordinate, ZCoordinate, Angle, HP, NameCheck, Quests, SkywayLocationFlag, FirstUseFlag) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING PlayerID",
			AccountID, slot, FirstName, LastName, config.SPAWN_X, config.SPAWN_Y, config.SPAWN_Z, 0, config.GetMaxHP(1), nameCheck, QuestFlag, SkywayLocationFlag, FirstUseFlag)
		if err != nil {
			return err
		}

		if err := scan.Row(&PlayerID, rows); err != nil {
			return err
		}

		// create appearance
		if _, err := tx.Exec("INSERT INTO Appearances (PlayerID) VALUES ($1)", PlayerID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return -1, nil
	}

	return PlayerID, nil
}

// TODO: should this operate on the raw packet? should we do validation here or prior?
func (db *DBHandler) FinishPlayer(character *protocol.SP_CL2LS_REQ_CHAR_CREATE, AccountId int) error {
	return db.Transaction(func(tx *sql.Tx) error {
		// update AppearanceFlag
		_, err := tx.Exec("UPDATE Players SET AppearanceFlag = 1 WHERE PlayerID = $1 AND AccountID = $2 AND AppearanceFlag = 0", character.PCStyle.IPC_UID, AccountId)
		if err != nil {
			return err
		}

		// update Appearance
		_, err = tx.Exec("UPDATE Appearances SET Body = $1, EyeColor = $2, FaceStyle = $3, Gender = $4, HairColor = $5, HairStyle = $6, Height = $7, SkinColor = $8 WHERE PlayerID = $9",
			character.PCStyle.IBody,
			character.PCStyle.IEyeColor,
			character.PCStyle.IFaceStyle,
			character.PCStyle.IGender,
			character.PCStyle.IHairColor,
			character.PCStyle.IHairStyle,
			character.PCStyle.IHeight,
			character.PCStyle.ISkinColor,
			character.PCStyle.IPC_UID)
		if err != nil {
			return err
		}

		// update Inventory
		items := [3]int16{character.SOn_Item.IEquipUBID, character.SOn_Item.IEquipLBID, character.SOn_Item.IEquipFootID}
		for i := 0; i < len(items); i++ {
			_, err = tx.Exec("INSERT INTO Inventory (PlayerID, Slot, ID, Type, Opt) VALUES ($1, $2, $3, $4, 1)", character.PCStyle.IPC_UID, i+1, items[i], i+1)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (db *DBHandler) FinishTutorial(PlayerID, AccountID int) error {
	_, err := db.Exec("UPDATE Players SET TutorialFlag = 1 WHERE PlayerID = $1 AND AccountID = $2 AND TutorialFlag = 0", PlayerID, AccountID)
	if err != nil {
		return err
	}

	// TODO: reference openfusion's finishTutorial for their academy specific patches
	return nil
}

// returns the deleted Slot number
func (db *DBHandler) DeletePlayer(PlayerID, AccountID int) (int, error) {
	row, err := db.Query("DELETE FROM Players WHERE AccountID = $1 AND PlayerID = $2 RETURNING Slot")
	if err != nil {
		return -1, err
	}

	var slot int
	if err := row.Scan(&slot); err != nil {
		return -1, err
	}

	return slot, nil
}

const (
	QUERY_PLAYERS = `SELECT
		p.PlayerID, p.AccountID, p.Slot, p.FirstName, p.LastName,
		p.Level, p.Nano1, p.Nano2, p.Nano3,
		p.AppearanceFlag, p.TutorialFlag, p.PayZoneFlag,
		p.XCoordinate, p.YCoordinate, p.ZCoordinate, p.NameCheck,
		p.Angle, p.HP, p.FusionMatter, p.Taros, p.Quests,
		p.BatteryW, p.BatteryN, p.Mentor, p.WarpLocationFlag,
		p.SkywayLocationFlag, p.CurrentMissionID, p.FirstUseFlag,
		a.Body, a.EyeColor, a.FaceStyle, a.Gender, a.HairColor, a.HairStyle,
		a.Height, a.SkinColor, acc.AccountLevel
	FROM Players as p
	INNER JOIN Appearances as a ON p.PlayerID = a.PlayerID
	INNER JOIN Accounts as acc ON p.AccountID = acc.AccountID `
)

func (db *DBHandler) readPlayer(rows *sql.Rows) (*Player, error) {
	plr := Player{ActiveNanoSlotNum: 0}

	if err := rows.Scan(
		&plr.PlayerID, &plr.AccountID, &plr.Slot, &plr.PCStyle.SzFirstName, &plr.PCStyle.SzLastName,
		&plr.Level, &plr.EquippedNanos[0], &plr.EquippedNanos[1], &plr.EquippedNanos[2],
		&plr.PCStyle2.IAppearanceFlag, &plr.PCStyle2.ITutorialFlag, &plr.PCStyle2.IPayzoneFlag,
		&plr.X, &plr.Y, &plr.Z, &plr.PCStyle.INameCheck,
		&plr.Angle, &plr.HP, &plr.FusionMatter, &plr.Taros, &plr.Quests,
		&plr.BatteryW, &plr.BatteryN, &plr.Mentor, &plr.WarpLocationFlag,
		&plr.SkywayLocationFlag, &plr.CurrentMissionID, &plr.FirstUseFlag,
		&plr.PCStyle.IBody, &plr.PCStyle.IEyeColor, &plr.PCStyle.IFaceStyle, &plr.PCStyle.IGender, &plr.PCStyle.IHairColor, &plr.PCStyle.IHairStyle,
		&plr.PCStyle.IHeight, &plr.PCStyle.ISkinColor, &plr.AccountLevel); err != nil {
		return nil, err
	}
	plr.PCStyle.IPC_UID = int64(plr.PlayerID)

	inven, err := db.GetPlayerInventorySlots(plr.PlayerID, 0, config.AEQUIP_COUNT+config.AINVEN_COUNT+config.ABANK_COUNT)
	if err != nil {
		return nil, err
	}

	// populate player inven
	for i, item := range inven {
		if item.IID == 0 { // skip empty slots
			continue
		}

		switch {
		case i < config.AEQUIP_COUNT:
			plr.Equip[i] = item
		case i < config.AEQUIP_COUNT+config.AINVEN_COUNT:
			plr.Inven[i-config.AEQUIP_COUNT] = item
		default:
			plr.Bank[i-config.AEQUIP_COUNT-config.AINVEN_COUNT] = item
		}
	}

	return &plr, nil
}

func (db *DBHandler) GetPlayer(PlayerID int) (*Player, error) {
	rows, err := db.Query(QUERY_PLAYERS+"WHERE p.PlayerID = $1", PlayerID)
	if err != nil {
		return nil, err
	}

	var plr *Player
	for rows.Next() {
		plr, err = db.readPlayer(rows)
		if err != nil {
			return nil, err
		}
	}

	return plr, nil
}

func (db *DBHandler) GetPlayers(AccountID int) ([]Player, error) {
	rows, err := db.Query(QUERY_PLAYERS+"WHERE p.AccountID = $1", AccountID)
	if err != nil {
		return nil, err
	}

	var plrs []Player
	for rows.Next() {
		plr, err := db.readPlayer(rows)
		if err != nil {
			return nil, err
		}

		plrs = append(plrs, *plr)
	}

	return plrs, nil
}
