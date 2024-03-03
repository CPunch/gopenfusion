package testutil

import (
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
)

var (
	TestCharCreate = protocol.SP_CL2LS_REQ_CHAR_CREATE{
		PCStyle: protocol.SPCStyle{
			INameCheck: 1, SzFirstName: "Hector",
			SzLastName: "Bannonvenom", IGender: 1, IFaceStyle: 1,
			IHairStyle: 17, IHairColor: 11, ISkinColor: 10, IEyeColor: 2,
			IHeight: 1, IBody: 0, IClass: 0,
		},
		SOn_Item: protocol.SOnItem{
			IEquipHandID: 0, IEquipUBID: 53, IEquipLBID: 17, IEquipFootID: 58,
			IEquipHeadID: 0, IEquipFaceID: 0, IEquipBackID: 0,
		},
		SOn_Item_Index: protocol.SOnItem_Index{
			IEquipUBID_index: 15, IEquipLBID_index: 12, IEquipFootID_index: 17,
			IFaceStyle: 2, IHairStyle: 18,
		},
	}
)

// creates a new account and player in the database
func MakeTestPlayer(db *db.DBHandler, id string, password string) (acc *db.Account, plr *db.Player, err error) {
	acc, err = db.NewAccount(id, password)
	if err != nil {
		return
	}

	var plrID int
	plrID, err = db.NewPlayer(acc.AccountID, TestCharCreate.PCStyle.SzFirstName, TestCharCreate.PCStyle.SzLastName, 1)
	if err != nil {
		return
	}

	charCreate := TestCharCreate
	charCreate.PCStyle.IPC_UID = int64(plrID)
	err = db.FinishPlayer(&charCreate, acc.AccountID)
	if err != nil {
		return
	}

	err = db.FinishTutorial(plrID, acc.AccountID)
	if err != nil {
		return
	}

	plr, err = db.GetPlayer(plrID)
	return
}

func QueueLogin(redisHndlr *redis.RedisHandler, FEKey []byte, plrID, accID int) (int64, error) {
	key, err := protocol.GenSerialKey()
	if err != nil {
		return 0, err
	}

	return key, redisHndlr.QueueLogin(key, redis.LoginMetadata{
		FEKey:     FEKey,
		PlayerID:  int32(plrID),
		AccountID: accID,
	})
}
