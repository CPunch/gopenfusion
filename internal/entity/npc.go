package entity

import (
	"sync/atomic"

	"github.com/CPunch/gopenfusion/internal/protocol"
)

type NPC struct {
	ID      int
	X       int `json:"iX"`
	Y       int `json:"iY"`
	Z       int `json:"iZ"`
	Angle   int `json:"iAngle"`
	NPCType int `json:"iNPCType"`
	Chunk   ChunkPosition
}

var nextNPCID = &atomic.Int32{}

func NewNPC(X, Y, Z, Angle int, npcType int) *NPC {
	return &NPC{
		ID:      int(nextNPCID.Add(1)),
		X:       X,
		Y:       Y,
		Z:       Z,
		Angle:   Angle,
		NPCType: npcType,
		Chunk:   MakeChunkPosition(X, Y),
	}
}

// ==================== Entity interface ====================

func (npc *NPC) GetKind() EntityKind {
	return ENTITY_KIND_NPC
}

func (npc *NPC) GetChunk() ChunkPosition {
	return npc.Chunk
}

func (npc *NPC) GetPosition() (x int, y int, z int) {
	return npc.X, npc.Y, npc.Z
}

func (npc *NPC) GetAngle() int {
	return npc.Angle
}

func (npc *NPC) SetChunk(chunk ChunkPosition) {
	npc.Chunk = chunk
}

func (npc *NPC) SetPosition(x, y, z int) {
	npc.X = x
	npc.Y = y
	npc.Z = z
}

func (npc *NPC) SetAngle(angle int) {
	npc.Angle = angle
}

func (npc *NPC) DisappearFromViewOf(peer *protocol.CNPeer) {
	peer.Send(protocol.P_FE2CL_NPC_EXIT, protocol.SP_FE2CL_NPC_EXIT{
		INPC_ID: int32(npc.ID),
	})
}

func (npc *NPC) EnterIntoViewOf(peer *protocol.CNPeer) {
	peer.Send(protocol.P_FE2CL_NPC_NEW, protocol.SP_FE2CL_NPC_NEW{
		NPCAppearanceData: npc.GetAppearanceData(),
	})
}

func (npc *NPC) GetAppearanceData() protocol.SNPCAppearanceData {
	return protocol.SNPCAppearanceData{
		INPC_ID:  int32(npc.ID),
		INPCType: int32(npc.NPCType),
		IHP:      100,
		IX:       int32(npc.X),
		IY:       int32(npc.Y),
		IZ:       int32(npc.Z),
		IAngle:   int32(npc.Angle),
	}
}
