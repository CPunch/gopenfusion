package entity

import (
	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/protocol"
)

type Player struct {
	db.Player
	Peer  *cnet.Peer
	Chunk ChunkPosition
}

func NewPlayer(peer *cnet.Peer, player *db.Player) *Player {
	return &Player{
		Player: *player,
		Peer:   peer,
		Chunk:  MakeChunkPosition(player.X, player.Y),
	}
}

// ==================== Entity interface ====================

func (plr *Player) GetKind() EntityKind {
	return ENTITY_KIND_PLAYER
}

func (plr *Player) GetChunkPos() ChunkPosition {
	return plr.Chunk
}

func (plr *Player) GetPosition() (x int, y int, z int) {
	return plr.X, plr.Y, plr.Z
}

func (plr *Player) GetAngle() int {
	return plr.Angle
}

func (plr *Player) SetChunkPos(chunk ChunkPosition) {
	plr.Chunk = chunk
}

func (plr *Player) SetPosition(x, y, z int) {
	plr.X = x
	plr.Y = y
	plr.Z = z
}

func (plr *Player) SetAngle(angle int) {
	plr.Angle = angle
}

func (plr *Player) DisappearFromViewOf(peer *cnet.Peer) {
	peer.Send(protocol.P_FE2CL_PC_EXIT, protocol.SP_FE2CL_PC_EXIT{
		IID: int32(plr.PlayerID),
	})
}

func (plr *Player) EnterIntoViewOf(peer *cnet.Peer) {
	peer.Send(protocol.P_FE2CL_PC_NEW, protocol.SP_FE2CL_PC_NEW{
		PCAppearanceData: plr.GetAppearanceData(),
	})
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

func (plr *Player) GetAppearanceData() protocol.SPCAppearanceData {
	return protocol.SPCAppearanceData{
		IID:       int32(plr.PlayerID),
		IHP:       int32(plr.HP),
		ILv:       int16(plr.Level),
		IX:        int32(plr.X),
		IY:        int32(plr.Y),
		IZ:        int32(plr.Z),
		IAngle:    int32(plr.Angle),
		PCStyle:   plr.PCStyle,
		IPCState:  plr.IPCState,
		ItemEquip: plr.Equip,
		Nano:      plr.Nanos[plr.ActiveNanoSlotNum],
	}
}
