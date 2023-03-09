package server

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/CPunch/GopenFusion/db"
	"github.com/CPunch/GopenFusion/protocol"
)

const (
	USE_E = iota
	USE_FE
)

type ClientHandler interface {
	HandlePacket(client *Client, typeID uint32, pkt *protocol.Packet)
	Connect(client *Client)
	Disconnect(client *Client)
}

type Client struct {
	E_key     []byte
	FE_key    []byte
	SzID      string
	AccountID int
	Player    *db.Player
	handler   ClientHandler
	conn      net.Conn
	alive     bool
	whichKey  int
}

func NewClient(handler ClientHandler, conn net.Conn) *Client {
	return &Client{
		E_key:     []byte(protocol.DEFAULT_KEY),
		FE_key:    nil,
		SzID:      "",
		AccountID: -1,
		Player:    nil,
		handler:   handler,
		conn:      conn,
		alive:     true,
		whichKey:  USE_E,
	}
}

func (client *Client) Send(data interface{}, typeID uint32) {
	// encode
	pkt := protocol.NewPacket(make([]byte, 0))
	pkt.Encode(data)
	log.Printf("Sending %#v, sizeof: %d", data, len(pkt.Buf))

	// write packet size
	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, uint32(len(pkt.Buf)+4))
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet size! %v", err))
	}

	// prepend the typeID to the packet body
	binary.LittleEndian.PutUint32(tmp, uint32(typeID))
	tmp = append(tmp, pkt.Buf...)

	// encrypt typeID & body
	switch client.whichKey {
	case USE_E:
		protocol.EncryptData(tmp, client.E_key)
	case USE_FE:
		protocol.EncryptData(tmp, client.FE_key)
	}

	// write packet body
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet body! %v", err))
	}
}

func (client *Client) Kill() {
	if !client.alive {
		return
	}

	client.alive = false
	client.conn.Close()
	client.handler.Disconnect(client)
}

func (client *Client) ClientHandler() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %p panic'd! %v", client, err)
		}
		client.Kill()
	}()

	tmp := make([]byte, 4, protocol.CN_PACKET_BUFFER_SIZE)
	for {
		// read packet size
		if _, err := client.conn.Read(tmp); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet size! %v", err))
		}
		sz := int(binary.LittleEndian.Uint32(tmp))

		// client should never send a packet size outside of this range
		if sz > protocol.CN_PACKET_BUFFER_SIZE || sz < 4 {
			panic(fmt.Errorf("[FATAL] malicious packet size received! %d", sz))
		}

		// read packet body
		if _, err := client.conn.Read(tmp[:sz]); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet body! %v", err))
		}

		// decrypt && grab typeID
		protocol.DecryptData(tmp[:sz], client.E_key)
		typeID := uint32(binary.LittleEndian.Uint32(tmp[:4]))

		// dispatch packet
		log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, sz)
		pkt := protocol.NewPacket(tmp[4:sz])
		client.handler.HandlePacket(client, typeID, pkt)

		// reset tmp
		tmp = tmp[:4]
	}
}
