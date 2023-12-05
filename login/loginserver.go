package login

import (
	"context"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
)

type LoginServer struct {
	service    *cnet.Service
	dbHndlr    *db.DBHandler
	redisHndlr *redis.RedisHandler
}

func NewLoginServer(ctx context.Context, dbHndlr *db.DBHandler, redisHndlr *redis.RedisHandler, port int) (*LoginServer, error) {
	srvc := cnet.NewService(ctx, "LOGIN", port)

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
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SHARD_SELECT, cnet.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SHARD_LIST_INFO, cnet.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_CHECK_NAME_LIST, cnet.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SAVE_CHAR_TUTOR, server.FinishTutorial)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_PC_EXIT_DUPLICATE, cnet.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REP_LIVE_CHECK, cnet.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_CHANGE_CHAR_NAME, cnet.StubbedPacket)
	srvc.AddPacketHandler(protocol.P_CL2LS_REQ_SERVER_SELECT, cnet.StubbedPacket)

	return server, nil
}

func (server *LoginServer) Service() *cnet.Service {
	return server.service
}

func (server *LoginServer) Start() error {
	return server.service.Start()
}
