package shard

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/core/db"
	"github.com/CPunch/gopenfusion/core/entity"
	"github.com/CPunch/gopenfusion/core/protocol"
	"github.com/CPunch/gopenfusion/core/protocol/pool"
	"github.com/CPunch/gopenfusion/core/redis"
)

type PacketHandler func(peer *protocol.CNPeer, pkt protocol.Packet) error

type ShardServer struct {
	listener       net.Listener
	port           int
	dbHndlr        *db.DBHandler
	redisHndlr     *redis.RedisHandler
	eRecv          chan *protocol.Event
	packetHandlers map[uint32]PacketHandler
	peers          map[*protocol.CNPeer]*entity.Player
	chunks         map[entity.ChunkPosition]*entity.Chunk
	peerLock       sync.Mutex
}

func NewShardServer(dbHndlr *db.DBHandler, redisHndlr *redis.RedisHandler, port int) (*ShardServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &ShardServer{
		listener:       listener,
		port:           port,
		dbHndlr:        dbHndlr,
		redisHndlr:     redisHndlr,
		packetHandlers: make(map[uint32]PacketHandler),
		peers:          make(map[*protocol.CNPeer]*entity.Player),
		chunks:         make(map[entity.ChunkPosition]*entity.Chunk),
		eRecv:          make(chan *protocol.Event),
	}

	server.packetHandlers = map[uint32]PacketHandler{
		protocol.P_CL2FE_REQ_PC_ENTER:              server.RequestEnter,
		protocol.P_CL2FE_REQ_PC_LOADING_COMPLETE:   server.LoadingComplete,
		protocol.P_CL2FE_REQ_PC_MOVE:               server.playerMove,
		protocol.P_CL2FE_REQ_PC_STOP:               server.playerStop,
		protocol.P_CL2FE_REQ_PC_JUMP:               server.playerJump,
		protocol.P_CL2FE_REQ_SEND_FREECHAT_MESSAGE: server.freeChat,
		protocol.P_CL2FE_REQ_SEND_MENUCHAT_MESSAGE: server.menuChat,
		protocol.P_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT: server.emoteChat,
	}

	redisHndlr.RegisterShard(redis.ShardMetadata{
		IP:   config.GetAnnounceIP(),
		Port: port,
	})

	return server, nil
}

func (server *ShardServer) handleEvents() {
	for event := range server.eRecv {
		switch event.Type {
		case protocol.EVENT_CLIENT_DISCONNECT:
			server.disconnect(event.Peer)
		case protocol.EVENT_CLIENT_PACKET:
			if err := server.handlePacket(event.Peer, event.PktID, protocol.NewPacket(event.Pkt)); err != nil {
				event.Peer.Kill()
			}

			// the packet is given to us by the event, so we'll need to make sure to return it to the pool
			pool.Put(event.Pkt)
		}
	}
}

func (server *ShardServer) Start() {
	server.LoadNPCs()

	log.Printf("Shard service hosted on %s:%d\n", config.GetAnnounceIP(), server.port)
	go server.handleEvents()
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}

		client := protocol.NewCNPeer(server.eRecv, conn)
		server.connect(client)
		go client.Handler()
	}
}

func (server *ShardServer) GetPort() int {
	return server.port
}

func (server *ShardServer) handlePacket(peer *protocol.CNPeer, typeID uint32, pkt protocol.Packet) error {
	if hndlr, ok := server.packetHandlers[typeID]; ok {
		if err := hndlr(peer, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] unknown packet ID: %x\n", typeID)
	}

	return nil
}

func (server *ShardServer) disconnect(peer *protocol.CNPeer) {
	server.peerLock.Lock()
	defer server.peerLock.Unlock()

	// remove from chunk(s)
	plr, ok := server.peers[peer]
	if ok {
		log.Printf("Player %d (AccountID %d) disconnected\n", plr.PlayerID, plr.AccountID)
		server.removeEntity(plr)
	}

	log.Printf("Peer %p disconnected from SHARD\n", peer)
	delete(server.peers, peer)
}

func (server *ShardServer) connect(peer *protocol.CNPeer) {
	server.peerLock.Lock()
	defer server.peerLock.Unlock()

	log.Printf("New peer %p connected to SHARD\n", peer)
	server.peers[peer] = nil
}

func (server *ShardServer) getPlayer(peer *protocol.CNPeer) (*entity.Player, error) {
	plr, ok := server.peers[peer]
	if !ok {
		return nil, fmt.Errorf("player not found")
	}

	return plr, nil
}

func (server *ShardServer) setPlayer(peer *protocol.CNPeer, plr *entity.Player) {
	server.peers[peer] = plr
}

// If f returns false the iteration is stopped.
func (server *ShardServer) rangePeers(f func(peer *protocol.CNPeer, plr *entity.Player) bool) {
	for peer, plr := range server.peers {
		if f(peer, plr) {
			return
		}
	}
}
