package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/CPunch/gopenfusion/db"
	"github.com/CPunch/gopenfusion/protocol"
)

type LoginServer struct {
	listener       net.Listener
	port           int
	dbHndlr        *db.DBHandler
	packetHandlers map[uint32]PacketHandler
	peers          map[*Peer]bool
	peersLock      sync.Mutex
	shard          *ShardServer
}

func NewLoginServer(dbHndlr *db.DBHandler, port int) (*LoginServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &LoginServer{
		listener: listener,
		port:     port,
		dbHndlr:  dbHndlr,
		peers:    make(map[*Peer]bool),
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
	log.Printf("Server hosted on 127.0.0.1:%d\n", server.port)

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}

		client := NewPeer(server, conn)
		server.Connect(client)
		go client.Handler()
	}
}

func (server *LoginServer) HandlePacket(peer *Peer, typeID uint32, pkt protocol.Packet) error {
	if hndlr, ok := server.packetHandlers[typeID]; ok {
		if err := hndlr(peer, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] invalid packet ID: %x\n", typeID)
	}

	return nil
}

func (server *LoginServer) Disconnect(peer *Peer) {
	server.peersLock.Lock()
	log.Printf("Peer %p disconnected from LOGIN\n", peer)
	delete(server.peers, peer)
	server.peersLock.Unlock()
}

func (server *LoginServer) Connect(peer *Peer) {
	server.peersLock.Lock()
	log.Printf("New peer %p connected to LOGIN\n", peer)
	server.peers[peer] = true
	server.peersLock.Unlock()
}

func (server *LoginServer) AddShard(shard *ShardServer) {
	server.shard = shard
}
