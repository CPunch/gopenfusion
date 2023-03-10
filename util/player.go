package util

import (
	"github.com/CPunch/gopenfusion/db"
	"github.com/CPunch/gopenfusion/protocol"
)

func Player2PCStyle(plr *db.Player) (protocol.SPCStyle, protocol.SPCStyle2) {
	return protocol.SPCStyle{
			IPC_UID:     int64(plr.PlayerID),
			INameCheck:  int8(plr.NameCheck),
			SzFirstName: plr.FirstName,
			SzLastName:  plr.LastName,
			IGender:     int8(plr.Gender),
			IFaceStyle:  int8(plr.FaceStyle),
			IHairStyle:  int8(plr.HairStyle),
			IHairColor:  int8(plr.HairColor),
			ISkinColor:  int8(plr.SkinColor),
			IEyeColor:   int8(plr.EyeColor),
			IHeight:     int8(plr.Height),
			IBody:       int8(plr.Body),
		},
		protocol.SPCStyle2{
			IAppearanceFlag: int8(plr.AppearanceFlag),
			ITutorialFlag:   int8(plr.TutorialFlag),
			IPayzoneFlag:    int8(plr.PayZoneFlag),
		}
}
