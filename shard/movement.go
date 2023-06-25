package shard

import (
	"github.com/CPunch/gopenfusion/core/protocol"
)

func (server *ShardServer) updatePlayerPosition(peer *protocol.CNPeer, X, Y, Z, Angle int) error {
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	plr.X = X
	plr.Y = Y
	plr.Z = Z
	plr.Angle = Angle

	return nil
}

func (server *ShardServer) playerMove(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var move protocol.SP_CL2FE_REQ_PC_MOVE
	pkt.Decode(&move)

	return server.updatePlayerPosition(peer, int(move.IX), int(move.IY), int(move.IZ), int(move.IAngle))
}
