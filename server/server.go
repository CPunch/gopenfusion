package main

import (
	"fmt"
	"log"
	"net"

	"github.com/CPunch/GoFusion/protocol"
)

type Server struct {
	listener   net.Listener
	clients    map[*Client]bool
	unregister chan *Client
}

func newServer() *Server {
	listener, err := net.Listen("tcp", ":23000")
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		listener:   listener,
		clients:    make(map[*Client]bool),
		unregister: make(chan *Client),
	}
}

func (server *Server) Start() {
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

			client := newClient(server, conn, []byte(protocol.E_KEY))
			server.clients[client] = true
			go client.ClientHandler()
			fmt.Printf("Client %p connected\n", client)
		}
	}
}

func main() {
	server := newServer()
	server.Start()
}
