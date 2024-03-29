package shard

import (
	"fmt"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/shard/entity"
)

func (server *ShardServer) updatePlayerPosition(plr *entity.Player, X, Y, Z, Angle int) {
	plr.X = X
	plr.Y = Y
	plr.Z = Z
	plr.Angle = Angle
	server.updateEntityChunk(plr, plr.GetChunkPos(), entity.MakeChunkPosition(X, Y))
}

func (server *ShardServer) playerMove(peer *cnet.Peer, pkt protocol.Packet) error {
	var move protocol.SP_CL2FE_REQ_PC_MOVE
	pkt.Decode(&move)

	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("playerMove: plr is nil")
	}

	// update chunking
	server.updatePlayerPosition(plr, int(move.IX), int(move.IY), int(move.IZ), int(move.IAngle))

	return server.sendOthersPacket(plr, protocol.P_FE2CL_PC_MOVE, protocol.SP_FE2CL_PC_MOVE{
		ICliTime:  move.ICliTime,
		IX:        move.IX,
		IY:        move.IY,
		IZ:        move.IZ,
		FVX:       move.FVX,
		FVY:       move.FVY,
		FVZ:       move.FVZ,
		IAngle:    move.IAngle,
		CKeyValue: move.CKeyValue,
		ISpeed:    move.ISpeed,
		IID:       int32(plr.PlayerID),
		ISvrTime:  protocol.GetTime(),
	})
}

func (server *ShardServer) playerStop(peer *cnet.Peer, pkt protocol.Packet) error {
	var stop protocol.SP_CL2FE_REQ_PC_STOP
	pkt.Decode(&stop)

	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("playerStop: plr is nil")
	}

	// update chunking
	server.updatePlayerPosition(plr, int(stop.IX), int(stop.IY), int(stop.IZ), plr.Angle)

	return server.sendOthersPacket(plr, protocol.P_FE2CL_PC_STOP, protocol.SP_FE2CL_PC_STOP{
		ICliTime: stop.ICliTime,
		IX:       stop.IX,
		IY:       stop.IY,
		IZ:       stop.IZ,
		IID:      int32(plr.PlayerID),
		ISvrTime: protocol.GetTime(),
	})
}

func (server *ShardServer) playerJump(peer *cnet.Peer, pkt protocol.Packet) error {
	var jump protocol.SP_CL2FE_REQ_PC_JUMP
	pkt.Decode(&jump)

	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("playerJump: _plr is nil")
	}

	// update chunking
	server.updatePlayerPosition(plr, int(jump.IX), int(jump.IY), int(jump.IZ), plr.Angle)

	return server.sendOthersPacket(plr, protocol.P_FE2CL_PC_JUMP, protocol.SP_FE2CL_PC_JUMP{
		ICliTime:  jump.ICliTime,
		IX:        jump.IX,
		IY:        jump.IY,
		IZ:        jump.IZ,
		IVX:       jump.IVX,
		IVY:       jump.IVY,
		IVZ:       jump.IVZ,
		IAngle:    jump.IAngle,
		CKeyValue: jump.CKeyValue,
		ISpeed:    jump.ISpeed,
		IID:       int32(plr.PlayerID),
		ISvrTime:  protocol.GetTime(),
	})
}
