package shard

import (
	"context"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/config"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/CPunch/gopenfusion/shard/entity"
)

type PacketHandler func(peer *cnet.Peer, pkt protocol.Packet) error

type ShardServer struct {
	service    *cnet.Service
	dbHndlr    *db.DBHandler
	redisHndlr *redis.RedisHandler
	chunks     map[entity.ChunkPosition]*entity.Chunk
}

func NewShardServer(ctx context.Context, dbHndlr *db.DBHandler, redisHndlr *redis.RedisHandler, port int) (*ShardServer, error) {
	srvc := cnet.NewService(ctx, "SHARD", port)

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

func (server *ShardServer) Start() error {
	server.LoadNPCs()
	return server.service.Start()
}

func (server *ShardServer) onDisconnect(peer *cnet.Peer) {
	// remove from chunks
	plr, ok := peer.UserData().(*entity.Player)
	if ok && plr != nil {
		server.removeEntity(plr)
	}
}

func (server *ShardServer) onConnect(peer *cnet.Peer) {

}

func (server *ShardServer) Service() *cnet.Service {
	return server.service
}
