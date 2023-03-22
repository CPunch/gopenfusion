package core

import (
	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/core/protocol"
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
}

func (plr *Player) ToPCLoadData2CL() protocol.SPCLoadData2CL {
	return protocol.SPCLoadData2CL{
		IUserLevel:         int16(plr.AccountLevel),
		PCStyle:            plr.PCStyle,
		PCStyle2:           plr.PCStyle2,
		IMentor:            int16(plr.Mentor),
		IMentorCount:       1,
		IHP:                int32(plr.HP),
		IBatteryW:          int32(plr.BatteryW),
		IBatteryN:          int32(plr.BatteryN),
		ICandy:             int32(plr.Taros),
		IFusionMatter:      int32(plr.FusionMatter),
		ISpecialState:      0,
		IMapNum:            0,
		IX:                 int32(plr.X),
		IY:                 int32(plr.Y),
		IZ:                 int32(plr.Z),
		IAngle:             int32(plr.Angle),
		AEquip:             plr.Equip,
		AInven:             plr.Inven,
		ANanoSlots:         [3]int16{int16(plr.EquippedNanos[0]), int16(plr.EquippedNanos[1]), int16(plr.EquippedNanos[2])},
		IActiveNanoSlotNum: int16(plr.ActiveNanoSlotNum),
		IWarpLocationFlag:  int32(plr.WarpLocationFlag),
		IBuddyWarpTime:     60,
		IFatigue:           50,
	}
}
