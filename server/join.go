package server

import (
	"fmt"
	"time"

	"github.com/CPunch/gopenfusion/core"
	"github.com/CPunch/gopenfusion/core/protocol"
)

func (server *ShardServer) RequestEnter(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var enter protocol.SP_CL2FE_REQ_PC_ENTER
	pkt.Decode(&enter)

	loginData, err := server.CheckLogin(enter.IEnterSerialKey)
	if err != nil {
		// the error codes for P_FE2CL_REP_PC_ENTER_FAIL aren't referenced in the client :(
		peer.Send(protocol.P_FE2CL_REP_PC_ENTER_FAIL, protocol.SP_FE2CL_REP_PC_ENTER_FAIL{})
		return err
	}

	// attach player
	var resp *protocol.SP_FE2CL_REP_PC_ENTER_SUCC
	if err := server.UpdatePlayer(peer, func(old *core.Player) (*core.Player, error) {
		if old != nil { // resending a shard enter packet?
			return nil, fmt.Errorf("resent enter packet!")
		}

		plr, err := server.dbHndlr.GetPlayer(int(loginData.PlayerID))
		if err != nil {
			return nil, err
		}

		resp = &protocol.SP_FE2CL_REP_PC_ENTER_SUCC{
			IID:           int32(plr.PlayerID),
			PCLoadData2CL: plr.ToPCLoadData2CL(),
			UiSvrTime:     uint64(time.Now().Unix()),
		}

		return plr, nil
	}); err != nil {
		peer.Send(protocol.P_FE2CL_REP_PC_ENTER_FAIL, protocol.SP_FE2CL_REP_PC_ENTER_FAIL{})
		return err
	}

	// setup key
	peer.E_key = protocol.CreateNewKey(resp.UiSvrTime, uint64(resp.IID+1), uint64(resp.PCLoadData2CL.IFusionMatter+1))
	peer.FE_key = loginData.FEKey
	peer.SetActiveKey(protocol.USE_FE)

	return peer.Send(protocol.P_FE2CL_REP_PC_ENTER_SUCC, resp)
}

func (server *ShardServer) LoadingComplete(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var loadComplete protocol.SP_CL2FE_REQ_PC_LOADING_COMPLETE
	pkt.Decode(&loadComplete)

	plr := server.LoadPlayer(peer)
	if plr == nil {
		return fmt.Errorf("peer has no player attached!")
	}

	return peer.Send(protocol.P_FE2CL_REP_PC_LOADING_COMPLETE_SUCC, protocol.SP_FE2CL_REP_PC_LOADING_COMPLETE_SUCC{IPC_ID: int32(plr.PlayerID)})
}
