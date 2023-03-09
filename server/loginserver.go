package server

import (
	"log"
	"net"
	"sync"

	"github.com/CPunch/GopenFusion/protocol"
)

type LoginServer struct {
	listener net.Listener
	peers    map[*Peer]bool
	lock     sync.Mutex
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

func (server *LoginServer) HandlePacket(peer *Peer, typeID uint32, pkt *protocol.Packet) {
	switch typeID {
	case protocol.P_CL2LS_REQ_LOGIN:
		server.Login(peer, pkt)
	case protocol.P_CL2LS_REQ_CHECK_CHAR_NAME:
		var charPkt protocol.SP_CL2LS_REQ_CHECK_CHAR_NAME
		pkt.Decode(&charPkt)

		peer.Send(&protocol.SP_LS2CL_REP_CHECK_CHAR_NAME_SUCC{
			SzFirstName: charPkt.SzFirstName,
			SzLastName:  charPkt.SzLastName,
		}, protocol.P_LS2CL_REP_CHECK_CHAR_NAME_SUCC)
	case protocol.P_CL2LS_REQ_SAVE_CHAR_NAME:
		server.SaveCharacterName(peer, pkt)
	case protocol.P_CL2LS_REQ_CHAR_CREATE:
		server.CharacterCreate(peer, pkt)
	default:
		log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
	}
}

func (server *LoginServer) Disconnect(peer *Peer) {
	server.lock.Lock()
	delete(server.peers, peer)
	server.lock.Unlock()
}

func (server *LoginServer) Connect(peer *Peer) {
	server.lock.Lock()
	server.peers[peer] = true
	server.lock.Unlock()
}
