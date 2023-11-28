package shard

import (
	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/entity"
	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/CPunch/gopenfusion/internal/service"
)

type PacketHandler func(peer *protocol.CNPeer, pkt protocol.Packet) error

type ShardServer struct {
	service    *service.Service
	dbHndlr    *db.DBHandler
	redisHndlr *redis.RedisHandler
	chunks     map[entity.ChunkPosition]*entity.Chunk
}

func NewShardServer(dbHndlr *db.DBHandler, redisHndlr *redis.RedisHandler, port int) (*ShardServer, error) {
	srvc, err := service.NewService("SHARD", port)
	if err != nil {
		return nil, err
	}

	server := &ShardServer{
		service:    srvc,
		dbHndlr:    dbHndlr,
		redisHndlr: redisHndlr,
		chunks:     make(map[entity.ChunkPosition]*entity.Chunk),
	}

	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_PC_ENTER, server.RequestEnter)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_PC_LOADING_COMPLETE, server.LoadingComplete)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_PC_MOVE, server.playerMove)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_PC_STOP, server.playerStop)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_PC_JUMP, server.playerJump)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_SEND_FREECHAT_MESSAGE, server.freeChat)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_SEND_MENUCHAT_MESSAGE, server.menuChat)
	srvc.AddPacketHandler(protocol.P_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT, server.emoteChat)

	srvc.OnConnect = server.onConnect
	srvc.OnDisconnect = server.onDisconnect

	redisHndlr.RegisterShard(redis.ShardMetadata{
		IP:   config.GetAnnounceIP(),
		Port: port,
	})

	return server, nil
}

func (server *ShardServer) Start() {
	server.LoadNPCs()
	server.service.Start()
}

func (server *ShardServer) onDisconnect(peer *protocol.CNPeer, _plr interface{}) {
	// remove from chunks
	if _plr != nil {
		server.removeEntity(_plr.(*entity.Player))
	}
}

func (server *ShardServer) onConnect(peer *protocol.CNPeer) interface{} {
	return nil
}
