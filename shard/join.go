package shard

import (
	"fmt"
	"log"
	"time"

	"github.com/CPunch/gopenfusion/core/entity"
	"github.com/CPunch/gopenfusion/core/protocol"
	"github.com/CPunch/gopenfusion/core/redis"
)

func (server *ShardServer) attachPlayer(peer *protocol.CNPeer, meta redis.LoginMetadata) (*entity.Player, error) {
	// resending a shard enter packet?
	old, err := server.getPlayer(peer)
	if old != nil {
		return nil, fmt.Errorf("resent enter packet!")
	}

	// attach player
	plr, err := server.dbHndlr.GetPlayer(int(meta.PlayerID))
	if err != nil {
		return nil, err
	}
	plr.Peer = peer

	server.setPlayer(peer, plr)
	return plr, nil
}

func (server *ShardServer) RequestEnter(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var enter protocol.SP_CL2FE_REQ_PC_ENTER
	pkt.Decode(&enter)

	loginData, err := server.redisHndlr.GetLogin(enter.IEnterSerialKey)
	if err != nil {
		// the error codes for P_FE2CL_REP_PC_ENTER_FAIL aren't referenced in the client :(
		peer.Send(protocol.P_FE2CL_REP_PC_ENTER_FAIL, protocol.SP_FE2CL_REP_PC_ENTER_FAIL{})
		return err
	}

	plr, err := server.attachPlayer(peer, loginData)
	if err != nil {
		return err
	}

	resp := &protocol.SP_FE2CL_REP_PC_ENTER_SUCC{
		IID:           int32(plr.PlayerID),
		PCLoadData2CL: plr.ToPCLoadData2CL(),
		UiSvrTime:     uint64(time.Now().Unix()),
	}

	// setup peer
	peer.E_key = protocol.CreateNewKey(resp.UiSvrTime, uint64(resp.IID+1), uint64(resp.PCLoadData2CL.IFusionMatter+1))
	peer.FE_key = loginData.FEKey
	peer.PlayerID = loginData.PlayerID
	peer.AccountID = loginData.AccountID
	peer.SetActiveKey(protocol.USE_FE)

	log.Printf("Player %d (AccountID %d) entered\n", resp.IID, loginData.AccountID)

	server.updatePlayerPosition(peer, int(plr.X), int(plr.Y), int(plr.Z), int(plr.Angle))
	return peer.Send(protocol.P_FE2CL_REP_PC_ENTER_SUCC, resp)
}

func (server *ShardServer) LoadingComplete(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var loadComplete protocol.SP_CL2FE_REQ_PC_LOADING_COMPLETE
	pkt.Decode(&loadComplete)

	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	return peer.Send(protocol.P_FE2CL_REP_PC_LOADING_COMPLETE_SUCC, protocol.SP_FE2CL_REP_PC_LOADING_COMPLETE_SUCC{IPC_ID: int32(plr.PlayerID)})
}
