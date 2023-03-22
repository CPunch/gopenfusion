package server

import "github.com/CPunch/gopenfusion/core/protocol"

func (server *ShardServer) RequestEnter(peer *protocol.CNPeer, pkt protocol.Packet) (retErr error) {
	var enter protocol.SP_CL2FE_REQ_PC_ENTER
	pkt.Decode(&enter)

	loginData, err := server.CheckLogin(enter.IEnterSerialKey)
	if err != nil {
		// the error codes for P_FE2CL_REP_PC_ENTER_FAIL aren't referenced in the client :(
		peer.Send(protocol.P_FE2CL_REP_PC_ENTER_FAIL, protocol.SP_FE2CL_REP_PC_ENTER_FAIL{})
		return err
	}

	plr, err := server.dbHndlr.GetPlayer(int(loginData.PlayerID))
	if err != nil {
		peer.Send(protocol.P_FE2CL_REP_PC_ENTER_FAIL, protocol.SP_FE2CL_REP_PC_ENTER_FAIL{})
		return err
	}

	// TODO
	_ = plr
	return nil
}
