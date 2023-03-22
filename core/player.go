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
