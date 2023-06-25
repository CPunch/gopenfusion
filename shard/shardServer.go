package shard

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/core"
	"github.com/CPunch/gopenfusion/core/db"
	"github.com/CPunch/gopenfusion/core/protocol"
	"github.com/CPunch/gopenfusion/core/redis"
)

type PacketHandler func(peer *protocol.CNPeer, pkt protocol.Packet) error

func stubbedPacket(_ *protocol.CNPeer, _ protocol.Packet) error { /* stubbed */ return nil }

type ShardServer struct {
	listener       net.Listener
	port           int
	dbHndlr        *db.DBHandler
	redisHndlr     *redis.RedisHandler
	packetHandlers map[uint32]PacketHandler
	peersLock      sync.Mutex
	peers          map[*protocol.CNPeer]*core.Player
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
	}

	server.packetHandlers = map[uint32]PacketHandler{
		protocol.P_CL2FE_REQ_PC_ENTER:            server.RequestEnter,
		protocol.P_CL2FE_REQ_PC_LOADING_COMPLETE: server.LoadingComplete,
	}

	redisHndlr.RegisterShard(redis.ShardMetadata{
		IP:   config.GetAnnounceIP(),
		Port: port,
	})

	return server, nil
}

func (server *ShardServer) RegisterPacketHandler(typeID uint32, hndlr PacketHandler) {
	server.packetHandlers[typeID] = hndlr
}

func (server *ShardServer) Start() {
	log.Printf("Shard service hosted on %s:%d\n", config.GetAnnounceIP(), server.port)

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}

		client := protocol.NewCNPeer(server, conn)
		server.Connect(client)
		go client.Handler()
	}
}

func (server *ShardServer) GetPort() int {
	return server.port
}

func (server *ShardServer) HandlePacket(peer *protocol.CNPeer, typeID uint32, pkt protocol.Packet) error {
	if hndlr, ok := server.packetHandlers[typeID]; ok {
		if err := hndlr(peer, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] invalid packet ID: %x\n", typeID)
	}

	return nil
}

func (server *ShardServer) Disconnect(peer *protocol.CNPeer) {
	log.Printf("Peer %p disconnected from SHARD\n", peer)
	delete(server.peers, peer)
}

func (server *ShardServer) Connect(peer *protocol.CNPeer) {
	log.Printf("New peer %p connected to SHARD\n", peer)
	server.peers[peer] = nil
}

// Returns a copy of the player
func (server *ShardServer) LoadPlayer(peer *protocol.CNPeer) (core.Player, error) {
	plr, ok := server.peers[peer]
	if !ok {
		return core.Player{}, fmt.Errorf("Player not found")
	}

	return *plr, nil
}

// UpdatePlayer locks the peers map, and calls the provided callback. The returned new pointer will be stored, however if an error returns it will be passed back.
// Since it is UNSAFE to write to the returned pointer from LoadPlayer, this wrapper is for the cases that state in the player struct needs to be updated.
// The pointers new and old may be the same if you are just updating struct fields. This function should NOT be called recursively.
func (server *ShardServer) UpdatePlayer(peer *protocol.CNPeer, f func(old *core.Player) (new *core.Player, err error)) error {
	server.peersLock.Lock()
	defer server.peersLock.Unlock()

	// on fail, the player should not be stored
	new, err := f(server.peers[peer])
	if err != nil {
		return err
	}

	server.peers[peer] = new
	return nil
}

// If f returns false the iteration is stopped.
func (server *ShardServer) RangePeers(f func(peer *protocol.CNPeer) bool) {
	for peer := range server.peers {
		if f(peer) {
			return
		}
	}
}
