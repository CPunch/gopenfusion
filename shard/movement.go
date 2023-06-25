package shard

import (
	"time"

	"github.com/CPunch/gopenfusion/core/entity"
	"github.com/CPunch/gopenfusion/core/protocol"
)

func (server *ShardServer) updatePlayerPosition(plr *entity.Player, X, Y, Z, Angle int) error {
	plr.X = X
	plr.Y = Y
	plr.Z = Z
	plr.Angle = Angle
	server.updateEntityChunk(plr, plr.GetChunk(), entity.MakeChunkPosition(X, Y))
	return nil
}

func (server *ShardServer) playerMove(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var move protocol.SP_CL2FE_REQ_PC_MOVE
	pkt.Decode(&move)

	// sanity check
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	// update chunking
	if err := server.updatePlayerPosition(plr, int(move.IX), int(move.IY), int(move.IZ), int(move.IAngle)); err != nil {
		return err
	}

	return server.sendOthersPacket(plr, protocol.P_FE2CL_PC_MOVE, protocol.SP_FE2CL_PC_MOVE{
		ICliTime:  uint64(time.Now().Unix()),
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
		ISvrTime:  uint64(time.Now().Unix()),
	})
}

func (server *ShardServer) playerStop(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var stop protocol.SP_CL2FE_REQ_PC_STOP
	pkt.Decode(&stop)

	// sanity check
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	// update chunking
	if err := server.updatePlayerPosition(plr, int(stop.IX), int(stop.IY), int(stop.IZ), plr.Angle); err != nil {
		return err
	}

	return server.sendOthersPacket(plr, protocol.P_FE2CL_PC_STOP, protocol.SP_FE2CL_PC_STOP{
		ICliTime: uint64(time.Now().Unix()),
		IX:       stop.IX,
		IY:       stop.IY,
		IZ:       stop.IZ,
		IID:      int32(plr.PlayerID),
		ISvrTime: uint64(time.Now().Unix()),
	})
}

func (server *ShardServer) playerJump(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var jump protocol.SP_CL2FE_REQ_PC_JUMP
	pkt.Decode(&jump)

	// sanity check
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	// update chunking
	if err := server.updatePlayerPosition(plr, int(jump.IX), int(jump.IY), int(jump.IZ), plr.Angle); err != nil {
		return err
	}

	return server.sendOthersPacket(plr, protocol.P_FE2CL_PC_JUMP, protocol.SP_FE2CL_PC_JUMP{
		ICliTime:  uint64(time.Now().Unix()),
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
		ISvrTime:  uint64(time.Now().Unix()),
	})
}
