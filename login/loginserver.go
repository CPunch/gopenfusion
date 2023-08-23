package login

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/CPunch/gopenfusion/internal/protocol/pool"
	"github.com/CPunch/gopenfusion/internal/redis"
)

type PacketHandler func(peer *protocol.CNPeer, pkt protocol.Packet) error

func stubbedPacket(_ *protocol.CNPeer, _ protocol.Packet) error { /* stubbed */ return nil }

type LoginServer struct {
	listener       net.Listener
	port           int
	dbHndlr        *db.DBHandler
	redisHndlr     *redis.RedisHandler
	eRecv          chan *protocol.Event
	peers          map[*protocol.CNPeer]bool
	packetHandlers map[uint32]PacketHandler
	peerLock       sync.Mutex
}

func NewLoginServer(dbHndlr *db.DBHandler, redisHndlr *redis.RedisHandler, port int) (*LoginServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &LoginServer{
		listener:   listener,
		port:       port,
		dbHndlr:    dbHndlr,
		redisHndlr: redisHndlr,
		peers:      make(map[*protocol.CNPeer]bool),
		eRecv:      make(chan *protocol.Event),
	}

	server.packetHandlers = map[uint32]PacketHandler{
		protocol.P_CL2LS_REQ_LOGIN:             server.Login,
		protocol.P_CL2LS_REQ_CHECK_CHAR_NAME:   server.CheckCharacterName,
		protocol.P_CL2LS_REQ_SAVE_CHAR_NAME:    server.SaveCharacterName,
		protocol.P_CL2LS_REQ_CHAR_CREATE:       server.CharacterCreate,
		protocol.P_CL2LS_REQ_CHAR_SELECT:       server.ShardSelect,
		protocol.P_CL2LS_REQ_CHAR_DELETE:       server.CharacterDelete,
		protocol.P_CL2LS_REQ_SHARD_SELECT:      stubbedPacket,
		protocol.P_CL2LS_REQ_SHARD_LIST_INFO:   stubbedPacket,
		protocol.P_CL2LS_CHECK_NAME_LIST:       stubbedPacket,
		protocol.P_CL2LS_REQ_SAVE_CHAR_TUTOR:   server.FinishTutorial,
		protocol.P_CL2LS_REQ_PC_EXIT_DUPLICATE: stubbedPacket,
		protocol.P_CL2LS_REP_LIVE_CHECK:        stubbedPacket,
		protocol.P_CL2LS_REQ_CHANGE_CHAR_NAME:  stubbedPacket,
		protocol.P_CL2LS_REQ_SERVER_SELECT:     stubbedPacket,
	}

	return server, nil
}

func (server *LoginServer) Start() {
	log.Printf("Login service hosted on %s:%d\n", config.GetAnnounceIP(), server.port)

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

func (server *LoginServer) handleEvents() {
	for event := range server.eRecv {
		switch event.Type {
		case protocol.EVENT_CLIENT_DISCONNECT:
			server.disconnect(event.Peer)
		case protocol.EVENT_CLIENT_PACKET:
			if err := server.handlePacket(event.Peer, event.PktID, protocol.NewPacket(event.Pkt)); err != nil {
				log.Printf("Error handling packet: %v", err)
				event.Peer.Kill()
			}

			// the packet is given to us by the event, so we'll need to make sure to return it to the pool
			pool.Put(event.Pkt)
		}
	}
}

func (server *LoginServer) handlePacket(peer *protocol.CNPeer, typeID uint32, pkt protocol.Packet) error {
	if hndlr, ok := server.packetHandlers[typeID]; ok {
		if err := hndlr(peer, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] unknown packet ID: %x\n", typeID)
	}

	return nil
}

func (server *LoginServer) disconnect(peer *protocol.CNPeer) {
	server.peerLock.Lock()
	defer server.peerLock.Unlock()

	log.Printf("Peer %p disconnected from LOGIN\n", peer)
	delete(server.peers, peer)
}

func (server *LoginServer) connect(peer *protocol.CNPeer) {
	server.peerLock.Lock()
	defer server.peerLock.Unlock()

	log.Printf("New peer %p connected to LOGIN\n", peer)
	server.peers[peer] = true
}
