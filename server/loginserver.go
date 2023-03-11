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

func NewLoginServer() *LoginServer {
	listener, err := net.Listen("tcp", ":23000")
	if err != nil {
		log.Fatal(err)
	}

	loginServer := &LoginServer{
		listener:       listener,
		packetHandlers: make(map[uint32]PacketHandler),
		peers:          make(map[*Peer]bool),
	}

	loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_LOGIN, loginServer.Login)
	loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_CHECK_CHAR_NAME, loginServer.CheckCharacterName)
	loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_SAVE_CHAR_NAME, loginServer.SaveCharacterName)
	loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_CHAR_CREATE, loginServer.CharacterCreate)
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_CHAR_SELECT, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_CHAR_DELETE, loginServer.CharacterDelete)
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_SHARD_SELECT, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_SHARD_LIST_INFO, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_CHECK_NAME_LIST, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_SAVE_CHAR_TUTOR, loginServer.FinishTutorial)
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_PC_EXIT_DUPLICATE, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REP_LIVE_CHECK, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_CHANGE_CHAR_NAME, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })
	// loginServer.RegisterPacketHandler(protocol.P_CL2LS_REQ_SERVER_SELECT, func(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil })

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
		go client.ClientHandler()
	}
}

func (server *LoginServer) HandlePacket(peer *Peer, typeID uint32, pkt protocol.Packet) error {
	if hndlr, ok := server.packetHandlers[typeID]; ok {
		if err := hndlr(peer, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
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
