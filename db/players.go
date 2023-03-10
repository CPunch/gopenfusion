package db

import (
	"database/sql"

	"github.com/CPunch/GopenFusion/config"
	"github.com/CPunch/GopenFusion/protocol"
	"github.com/blockloop/scan"
)

type Player struct {
	PlayerID           int
	AccountID          int
	FirstName          string
	LastName           string
	NameCheck          int
	Slot               int
	Created            int
	LastLogin          int
	Level              int
	Nano1              int
	Nano2              int
	Nano3              int
	AppearanceFlag     int
	TutorialFlag       int
	PayZoneFlag        int
	XCoordinate        int
	YCoordinate        int
	ZCoordinate        int
	Angle              int
	HP                 int
	FusionMatter       int
	Taros              int
	BatteryW           int
	BatteryN           int
	Mentor             int
	CurrentMissionID   int
	WarpLocationFlag   int
	SkywayLocationFlag []byte
	FirstUseFlag       []byte
	Quests             []byte
	/* Appearance tbl */
	Body      int `db:"Body"`
	EyeColor  int `db:"EyeColor"`
	FaceStyle int `db:"FaceStyle"`
	Gender    int `db:"Gender"`
	HairColor int `db:"HairColor"`
	HairStyle int `db:"HairStyle"`
	Height    int `db:"Height"`
	SkinColor int `db:"SkinColor"`
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
		row, err := tx.Query(
			"INSERT INTO Players (AccountID, Slot, FirstName, LastName, XCoordinate, YCoordinate, ZCoordinate, Angle, HP, NameCheck, Quests, SkywayLocationFlag, FirstUseFlag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING PlayerID",
			AccountID, slot, FirstName, LastName, config.SPAWN_X, config.SPAWN_Y, config.SPAWN_Z, 0, config.GetMaxHP(1), nameCheck, QuestFlag, SkywayLocationFlag, FirstUseFlag)
		if err != nil {
			return err
		}

		if err := scan.Row(&PlayerID, row); err != nil {
			return err
		}

		// create appearance
		if _, err := tx.Exec("INSERT INTO Appearances (PlayerID) VALUES (?)", PlayerID); err != nil {
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
		_, err := tx.Exec("UPDATE Players SET AppearanceFlag = 1 WHERE PlayerID = ? AND AccountID = ? AND AppearanceFlag = 0", character.PCStyle.IPC_UID, AccountId)
		if err != nil {
			return err
		}

		// update Appearance
		_, err = tx.Exec("UPDATE Appearances SET Body = ?, EyeColor = ?, FaceStyle = ?, Gender = ?, HairColor = ?, HairStyle = ?, Height = ?, SkinColor = ? WHERE PlayerID = ?",
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
			_, err = tx.Exec("INSERT INTO Inventory (PlayerID, Slot, ID, Type, Opt) VALUES (?, ?, ?, ?, 1)", character.PCStyle.IPC_UID, i, items[i], i+1)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (db *DBHandler) FinishTutorial(PlayerID, AccountID int) error {
	_, err := db.Exec("UPDATE Players SET TutorialFlag = 1 WHERE PlayerID = ? AND AccountID = ? AND TutorialFlag = 0", PlayerID, AccountID)
	if err != nil {
		return err
	}

	// TODO: reference openfusion's finishTutorial for their academy specific patches
	return nil
}

// returns the deleted Slot number
func (db *DBHandler) DeletePlayer(PlayerID, AccountID int) (int, error) {
	row, err := db.Query("DELETE FROM Players WHERE AccountID = ? AND PlayerID = ? RETURNING Slot")
	if err != nil {
		return -1, err
	}

	var slot int
	for row.Next() {
		if err := row.Scan(&slot); err != nil {
			return -1, err
		}
	}

	return slot, nil
}

func (db *DBHandler) GetPlayer(PlayerID int) (*Player, error) {
	rows, err := db.Query(`SELECT
			p.PlayerID, p.AccountID, p.Slot, p.FirstName, p.LastName,
			p.Level, p.Nano1, p.Nano2, p.Nano3,
			p.AppearanceFlag, p.TutorialFlag, p.PayZoneFlag,
			p.XCoordinate, p.YCoordinate, p.ZCoordinate, p.NameCheck,
			p.Angle, p.HP, p.FusionMatter, p.Taros, p.Quests,
			p.BatteryW, p.BatteryN, p.Mentor, p.WarpLocationFlag,
			p.SkywayLocationFlag, p.CurrentMissionID, p.FirstUseFlag,
			a.Body, a.EyeColor, a.FaceStyle, a.Gender, a.HairColor, a.HairStyle,
			a.Height, a.SkinColor
		FROM Players as p
		INNER JOIN Appearances as a ON p.PlayerID = a.PlayerID
		WHERE p.PlayerID = ?`, PlayerID)
	if err != nil {
		return nil, err
	}

	var plr Player
	for rows.Next() {
		if err := rows.Scan(
			&plr.PlayerID, &plr.AccountID, &plr.Slot, &plr.FirstName, &plr.LastName,
			&plr.Level, &plr.Nano1, &plr.Nano2, &plr.Nano3,
			&plr.AppearanceFlag, &plr.TutorialFlag, &plr.PayZoneFlag,
			&plr.XCoordinate, &plr.YCoordinate, &plr.ZCoordinate, &plr.NameCheck,
			&plr.Angle, &plr.HP, &plr.FusionMatter, &plr.Taros, &plr.Quests,
			&plr.BatteryW, &plr.BatteryN, &plr.Mentor, &plr.WarpLocationFlag,
			&plr.SkywayLocationFlag, &plr.CurrentMissionID, &plr.FirstUseFlag,
			&plr.Body, &plr.EyeColor, &plr.FaceStyle, &plr.Gender, &plr.HairColor, &plr.HairStyle,
			&plr.Height, &plr.SkinColor); err != nil {
			panic(err)
		}
	}

	return &plr, nil
}

func (db *DBHandler) GetPlayers(AccountID int) ([]Player, error) {
	rows, err := db.Query(`SELECT
			p.PlayerID, p.AccountID, p.Slot, p.FirstName, p.LastName,
			p.Level, p.Nano1, p.Nano2, p.Nano3,
			p.AppearanceFlag, p.TutorialFlag, p.PayZoneFlag,
			p.XCoordinate, p.YCoordinate, p.ZCoordinate, p.NameCheck,
			p.Angle, p.HP, p.FusionMatter, p.Taros, p.Quests,
			p.BatteryW, p.BatteryN, p.Mentor, p.WarpLocationFlag,
			p.SkywayLocationFlag, p.CurrentMissionID, p.FirstUseFlag,
			a.Body, a.EyeColor, a.FaceStyle, a.Gender, a.HairColor, a.HairStyle,
			a.Height, a.SkinColor
		FROM Players as p
		INNER JOIN Appearances as a ON p.PlayerID = a.PlayerID
		WHERE p.AccountID = ?`, AccountID)
	if err != nil {
		return nil, err
	}

	var plrs []Player
	for rows.Next() {
		plr := Player{}

		if err := rows.Scan(
			&plr.PlayerID, &plr.AccountID, &plr.Slot, &plr.FirstName, &plr.LastName,
			&plr.Level, &plr.Nano1, &plr.Nano2, &plr.Nano3,
			&plr.AppearanceFlag, &plr.TutorialFlag, &plr.PayZoneFlag,
			&plr.XCoordinate, &plr.YCoordinate, &plr.ZCoordinate, &plr.NameCheck,
			&plr.Angle, &plr.HP, &plr.FusionMatter, &plr.Taros, &plr.Quests,
			&plr.BatteryW, &plr.BatteryN, &plr.Mentor, &plr.WarpLocationFlag,
			&plr.SkywayLocationFlag, &plr.CurrentMissionID, &plr.FirstUseFlag,
			&plr.Body, &plr.EyeColor, &plr.FaceStyle, &plr.Gender, &plr.HairColor, &plr.HairStyle,
			&plr.Height, &plr.SkinColor); err != nil {
			return nil, err
		}

		plrs = append(plrs, plr)
	}

	return plrs, nil
}
