package login

import (
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/CPunch/gopenfusion/internal/service"
)

type LoginServer struct {
	service    *service.Service
	dbHndlr    *db.DBHandler
	redisHndlr *redis.RedisHandler
}

func NewLoginServer(dbHndlr *db.DBHandler, redisHndlr *redis.RedisHandler, port int) (*LoginServer, error) {
	srvc, err := service.NewService("LOGIN", port)
	if err != nil {
		return nil, err
	}

	server := &LoginServer{
		service:    srvc,
		dbHndlr:    dbHndlr,
		redisHndlr: redisHndlr,
	}

	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_LOGIN, server.Login)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_CHECK_CHAR_NAME, server.CheckCharacterName)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SAVE_CHAR_NAME, server.SaveCharacterName)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_CHAR_CREATE, server.CharacterCreate)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_CHAR_SELECT, server.ShardSelect)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_CHAR_DELETE, server.CharacterDelete)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SHARD_SELECT, service.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SHARD_LIST_INFO, service.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_CHECK_NAME_LIST, service.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SAVE_CHAR_TUTOR, server.FinishTutorial)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_PC_EXIT_DUPLICATE, service.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REP_LIVE_CHECK, service.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_CHANGE_CHAR_NAME, service.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SERVER_SELECT, service.StubbedPacket)

	srvc.OnConnect = func(peer *protocol.CNPeer) interface{} {
		return nil
	}

	return server, nil
}

func (server *LoginServer) Start() {
	server.service.Start()
}
