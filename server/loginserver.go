package server

import (
	"log"
	"net"
	"sync"

	"github.com/CPunch/gopenfusion/protocol"
)

type LoginServer struct {
	listener  net.Listener
	peers     map[*Peer]bool
	peersLock sync.Mutex
}

func NewLoginServer() *LoginServer {
	listener, err := net.Listen("tcp", ":23000")
	if err != nil {
		log.Fatal(err)
	}

	return &LoginServer{
		listener: listener,
		peers:    make(map[*Peer]bool),
	}
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

func (server *LoginServer) HandlePacket(peer *Peer, typeID uint32, pkt protocol.Packet) {
	switch typeID {
	case protocol.P_CL2LS_REQ_LOGIN:
		server.Login(peer, pkt)
	case protocol.P_CL2LS_REQ_CHECK_CHAR_NAME:
		server.CheckCharacterName(peer, pkt)
	case protocol.P_CL2LS_REQ_SAVE_CHAR_NAME:
		server.SaveCharacterName(peer, pkt)
	case protocol.P_CL2LS_REQ_CHAR_CREATE:
		server.CharacterCreate(peer, pkt)
	case protocol.P_CL2LS_REQ_CHAR_SELECT:
		/* stubbed */
	case protocol.P_CL2LS_REQ_CHAR_DELETE:
		server.CharacterDelete(peer, pkt)
	case protocol.P_CL2LS_REQ_SHARD_SELECT:
		/* stubbed */
	case protocol.P_CL2LS_REQ_SHARD_LIST_INFO:
		/* stubbed */
	case protocol.P_CL2LS_CHECK_NAME_LIST:
		/* stubbed */
	case protocol.P_CL2LS_REQ_SAVE_CHAR_TUTOR:
		server.FinishTutorial(peer, pkt)
	case protocol.P_CL2LS_REQ_PC_EXIT_DUPLICATE:
		/* stubbed */
	case protocol.P_CL2LS_REP_LIVE_CHECK:
		/* stubbed */
	case protocol.P_CL2LS_REQ_CHANGE_CHAR_NAME:
		/* stubbed */
	case protocol.P_CL2LS_REQ_SERVER_SELECT:
		/* stubbed */
	default:
		log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
	}
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
