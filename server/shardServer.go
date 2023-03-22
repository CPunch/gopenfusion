package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/CPunch/gopenfusion/core"
	"github.com/CPunch/gopenfusion/core/db"
	"github.com/CPunch/gopenfusion/core/protocol"
)

type LoginMetadata struct {
	FEKey     []byte
	Timestamp time.Time
	PlayerID  int32
}

type ShardServer struct {
	listener           net.Listener
	port               int
	dbHndlr            *db.DBHandler
	packetHandlers     map[uint32]PacketHandler
	peers              sync.Map // [*protocol.CNPeer]*core.Player
	loginMetadataQueue sync.Map // [int64]*LoginMetadata w/ int64 = serialKey
}

func NewShardServer(dbHndlr *db.DBHandler, port int) (*ShardServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &ShardServer{
		listener:       listener,
		port:           port,
		dbHndlr:        dbHndlr,
		packetHandlers: make(map[uint32]PacketHandler),
	}

	server.packetHandlers = map[uint32]PacketHandler{
		protocol.P_CL2FE_REQ_PC_ENTER:            server.RequestEnter,
		protocol.P_CL2FE_REQ_PC_LOADING_COMPLETE: server.LoadingComplete,
	}

	return server, nil
}

func (server *ShardServer) RegisterPacketHandler(typeID uint32, hndlr PacketHandler) {
	server.packetHandlers[typeID] = hndlr
}

func (server *ShardServer) Start() {
	log.Printf("Server hosted on 127.0.0.1:%d\n", server.port)

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
	server.peers.Delete(peer)
}

func (server *ShardServer) Connect(peer *protocol.CNPeer) {
	log.Printf("New peer %p connected to SHARD\n", peer)
	server.peers.Store(peer, nil)
}

func (server *ShardServer) GetPlayer(peer *protocol.CNPeer) *core.Player {
	val, ok := server.peers.Load(peer)
	if !ok {
		return nil
	}

	plr, ok := val.(*core.Player)
	if !ok {
		return nil
	}

	return plr
}

func (server *ShardServer) JoinPlayer(peer *protocol.CNPeer, player *core.Player) {
	server.peers.Store(peer, player)
}

// Simple wrapper for server.peers.Range, if f returns false the iteration is stopped.
func (server *ShardServer) RangePeers(f func(peer *protocol.CNPeer, player *core.Player) bool) {
	server.peers.Range(func(key, value any) bool { // simple wrapper to cast the datatypes
		peer, ok := key.(*protocol.CNPeer)
		if !ok { // this should never happen
			panic(fmt.Errorf("ShardServer.peers has an invalid key: peers[%#v] = %#v", key, value))
		}

		player, ok := value.(*core.Player)
		if !ok { // this should also never happen
			panic(fmt.Errorf("ShardServer.peers has an invalid value: peers[%#v] = %#v", key, value))
		}

		return f(peer, player)
	})
}

func (server *ShardServer) QueueLogin(serialKey int64, meta *LoginMetadata) {
	server.loginMetadataQueue.Store(serialKey, meta)
}

func (server *ShardServer) CheckLogin(serialKey int64) (*LoginMetadata, error) {
	value, ok := server.loginMetadataQueue.Load(serialKey)
	if !ok {
		return nil, fmt.Errorf("serialKey %x is not valid!", serialKey)
	}

	meta, ok := value.(*LoginMetadata)
	if !ok { // should never happen
		panic(fmt.Errorf("ShardServer.loginMetadataQueue has an invalid value: loginMetadataQueue[%x] = %#v", serialKey, value))
	}

	return meta, nil
}
