package shard

import (
	"fmt"
	"log"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/CPunch/gopenfusion/shard/entity"
	"github.com/CPunch/gopenfusion/util"
)

func (server *ShardServer) attachPlayer(peer *cnet.Peer, meta redis.LoginMetadata) (*entity.Player, error) {
	dbPlr, err := server.dbHndlr.GetPlayer(int(meta.PlayerID))
	if err != nil {
		return nil, err
	}
	plr := entity.NewPlayer(peer, dbPlr)

	// once we create the player, it's memory address is owned by the
	// server.Start() goroutine. the only functions allowed to access
	// it are the packet handlers as no other goroutines will be
	// concurrently accessing it.
	peer.SetUserData(plr)
	return plr, nil
}

func (server *ShardServer) RequestEnter(peer *cnet.Peer, pkt protocol.Packet) error {
	var enter protocol.SP_CL2FE_REQ_PC_ENTER
	pkt.Decode(&enter)

	// resending a shard enter packet?
	_plr, ok := peer.UserData().(*entity.Player)
	if ok && _plr != nil {
		return fmt.Errorf("resent enter packet")
	}

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
		UiSvrTime:     util.GetTime(),
	}

	// setup peer
	peer.E_key = protocol.CreateNewKey(resp.UiSvrTime, uint64(resp.IID+1), uint64(resp.PCLoadData2CL.IFusionMatter+1))
	peer.FE_key = loginData.FEKey
	peer.SetActiveKey(cnet.USE_FE)

	log.Printf("Player %d (AccountID %d) entered\n", resp.IID, loginData.AccountID)
	if err := peer.Send(protocol.P_FE2CL_REP_PC_ENTER_SUCC, resp); err != nil {
		return err
	}

	return nil
}

func (server *ShardServer) LoadingComplete(peer *cnet.Peer, pkt protocol.Packet) error {
	var loadComplete protocol.SP_CL2FE_REQ_PC_LOADING_COMPLETE
	pkt.Decode(&loadComplete)

	// was the peer attached to a player?
	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("loadingComplete: plr is nil")
	}

	err := peer.Send(protocol.P_FE2CL_REP_PC_LOADING_COMPLETE_SUCC, protocol.SP_FE2CL_REP_PC_LOADING_COMPLETE_SUCC{IPC_ID: int32(plr.PlayerID)})
	if err != nil {
		return err
	}

	// we send the chunk updates (PC_NEW, NPC_NEW, etc.) after the enter packet
	chunkPos := entity.MakeChunkPosition(plr.X, plr.Y)
	viewableChunks := server.getViewableChunks(chunkPos)

	plr.SetChunkPos(chunkPos)
	server.getChunk(chunkPos).AddEntity(plr)
	server.addEntityToChunks(plr, viewableChunks)
	return nil
}
