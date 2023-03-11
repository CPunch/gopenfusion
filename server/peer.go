package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/CPunch/gopenfusion/db"
	"github.com/CPunch/gopenfusion/protocol"
	"github.com/CPunch/gopenfusion/protocol/pool"
)

const (
	USE_E = iota
	USE_FE
)

type PeerHandler interface {
	HandlePacket(client *Peer, typeID uint32, pkt protocol.Packet)
	Connect(client *Peer)
	Disconnect(client *Peer)
}

type Peer struct {
	E_key     []byte
	FE_key    []byte
	SzID      string
	AccountID int
	Player    *db.Player
	handler   PeerHandler
	conn      net.Conn
	alive     bool
	whichKey  int
}

func NewPeer(handler PeerHandler, conn net.Conn) *Peer {
	return &Peer{
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

func (client *Peer) Send(data interface{}, typeID uint32) {
	buf := pool.Get()
	defer func() { // always return the buffer to the pool
		pool.Put(buf)
	}()

	// encode
	pkt := protocol.NewPacket(buf)

	// write the typeID and packet body
	pkt.Encode(uint32(typeID))
	pkt.Encode(data)

	// write the packet size
	binary.Write(client.conn, binary.LittleEndian, uint32(buf.Len()))

	// encrypt typeID & body
	switch client.whichKey {
	case USE_E:
		protocol.EncryptData(buf.Bytes(), client.E_key)
	case USE_FE:
		protocol.EncryptData(buf.Bytes(), client.FE_key)
	}

	// write packet type && packet body
	log.Printf("Sending %#v, sizeof: %d", data, buf.Len())
	if _, err := client.conn.Write(buf.Bytes()); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet body! %v", err))
	}
}

func (client *Peer) Kill() {
	if !client.alive {
		return
	}

	client.alive = false
	client.conn.Close()
	client.handler.Disconnect(client)
}

func (client *Peer) ClientHandler() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %p panic'd! %v", client, err)
		}
		client.Kill()
	}()

	for {
		// read packet size
		var sz uint32
		if err := binary.Read(client.conn, binary.LittleEndian, &sz); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet size! %v", err))
		}

		// client should never send a packet size outside of this range
		if sz > protocol.CN_PACKET_BUFFER_SIZE || sz < 4 {
			panic(fmt.Errorf("[FATAL] malicious packet size received! %d", sz))
		}

		// read packet body
		buf := pool.Get()
		if _, err := buf.ReadFrom(io.LimitReader(client.conn, int64(sz))); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet body! %v", err))
		}

		// decrypt
		protocol.DecryptData(buf.Bytes(), client.E_key)

		// create packet && read typeID
		var typeID uint32
		pkt := protocol.NewPacket(buf)
		pkt.Decode(&typeID)

		// dispatch packet
		log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, sz)
		client.handler.HandlePacket(client, typeID, pkt)

		// restore buffer to pool
		pool.Put(buf)
	}
}
