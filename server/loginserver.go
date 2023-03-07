package server

import (
	"fmt"
	"log"
	"net"

	"github.com/CPunch/GopenFusion/protocol"
)

type LoginServer struct {
	listener   net.Listener
	clients    map[*protocol.Client]bool
	unregister chan *protocol.Client
}

func NewLoginServer() *LoginServer {
	listener, err := net.Listen("tcp", ":23000")
	if err != nil {
		log.Fatal(err)
	}

	return &LoginServer{
		listener:   listener,
		clients:    make(map[*protocol.Client]bool),
		unregister: make(chan *protocol.Client),
	}
}

func (server *LoginServer) Start() {
	log.Print("Server hosted on 127.0.0.1:23000")

	for {
		select {
		case client := <-server.unregister:
			delete(server.clients, client)
			fmt.Printf("Client %p disconnected\n", client)
		default:
			conn, err := server.listener.Accept()
			if err != nil {
				log.Println("Connection error: ", err)
				return
			}

			client := protocol.NewClient(server, conn)
			server.clients[client] = true
			fmt.Printf("Client %p connected\n", client)
			go client.ClientHandler()
		}
	}
}

func (server *LoginServer) HandlePacket(client *protocol.Client, typeID uint32, pkt *protocol.Packet) {
	switch typeID {
	case protocol.P_CL2LS_REQ_LOGIN:
		var loginPkt protocol.SP_CL2LS_REQ_LOGIN
		pkt.Decode(&loginPkt)

		// TODO: for now, we're a dummy server
		client.AcceptLogin(loginPkt.SzID, loginPkt.IClientVerC, 1, []protocol.SP_LS2CL_REP_CHAR_INFO{})
	case protocol.P_CL2LS_REQ_CHECK_CHAR_NAME:
		var charPkt protocol.SP_CL2LS_REQ_CHECK_CHAR_NAME
		pkt.Decode(&charPkt)

		client.Send(&protocol.SP_LS2CL_REP_CHECK_CHAR_NAME_SUCC{
			SzFirstName: charPkt.SzFirstName,
			SzLastName:  charPkt.SzLastName,
		}, protocol.P_LS2CL_REP_CHECK_CHAR_NAME_SUCC)
	case protocol.P_CL2LS_REQ_SAVE_CHAR_NAME:
		var charPkt protocol.SP_CL2LS_REQ_SAVE_CHAR_NAME
		pkt.Decode(&charPkt)

		client.Send(&protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_SUCC{
			IPC_UID:     1,
			ISlotNum:    charPkt.ISlotNum,
			IGender:     charPkt.IGender,
			SzFirstName: charPkt.SzFirstName,
			SzLastName:  charPkt.SzLastName,
		}, protocol.P_LS2CL_REP_SAVE_CHAR_NAME_SUCC)
	default:
		log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
	}
}

func (server *LoginServer) Disconnect(client *protocol.Client) {
	server.unregister <- client
}
