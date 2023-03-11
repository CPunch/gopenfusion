package server

import (
	"log"
	"net"
	"sync"

	"github.com/CPunch/gopenfusion/protocol"
)

type PacketHandler func(peer *Peer, pkt protocol.Packet) error

type LoginServer struct {
	listener       net.Listener
	packetHandlers map[uint32]PacketHandler
	peers          map[*Peer]bool
	peersLock      sync.Mutex
}

func stubbedPacket(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil }

func NewLoginServer() *LoginServer {
	listener, err := net.Listen("tcp", ":23000")
	if err != nil {
		log.Fatal(err)
	}

	loginServer := &LoginServer{
		listener: listener,
		peers:    make(map[*Peer]bool),
	}

	loginServer.packetHandlers = map[uint32]PacketHandler{
		protocol.P_CL2LS_REQ_LOGIN:             loginServer.Login,
		protocol.P_CL2LS_REQ_CHECK_CHAR_NAME:   loginServer.CheckCharacterName,
		protocol.P_CL2LS_REQ_SAVE_CHAR_NAME:    loginServer.SaveCharacterName,
		protocol.P_CL2LS_REQ_CHAR_CREATE:       loginServer.CharacterCreate,
		protocol.P_CL2LS_REQ_CHAR_SELECT:       stubbedPacket,
		protocol.P_CL2LS_REQ_CHAR_DELETE:       loginServer.CharacterDelete,
		protocol.P_CL2LS_REQ_SHARD_SELECT:      stubbedPacket,
		protocol.P_CL2LS_REQ_SHARD_LIST_INFO:   stubbedPacket,
		protocol.P_CL2LS_CHECK_NAME_LIST:       stubbedPacket,
		protocol.P_CL2LS_REQ_SAVE_CHAR_TUTOR:   loginServer.FinishTutorial,
		protocol.P_CL2LS_REQ_PC_EXIT_DUPLICATE: stubbedPacket,
		protocol.P_CL2LS_REP_LIVE_CHECK:        stubbedPacket,
		protocol.P_CL2LS_REQ_CHANGE_CHAR_NAME:  stubbedPacket,
		protocol.P_CL2LS_REQ_SERVER_SELECT:     stubbedPacket,
	}

	return loginServer
}

func (server *LoginServer) RegisterPacketHandler(typeID uint32, hndlr PacketHandler) {
	server.packetHandlers[typeID] = hndlr
}

func (server *LoginServer) Start() {
	log.Print("Server hosted on 127.0.0.1:23000")

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
	delete(server.peers, peer)
	server.peersLock.Unlock()
}

func (server *LoginServer) Connect(peer *Peer) {
	server.peersLock.Lock()
	server.peers[peer] = true
	server.peersLock.Unlock()
}
