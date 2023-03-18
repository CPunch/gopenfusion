package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/CPunch/gopenfusion/protocol/pool"
)

const (
	USE_E = iota
	USE_FE
)

type PeerHandler interface {
	HandlePacket(client *CNPeer, typeID uint32, pkt Packet) error
	Disconnect(client *CNPeer)
}

// CNPeer is a simple wrapper for net.Conn connections to send/recv packets over the Fusionfall packet protocol.
type CNPeer struct {
	conn      net.Conn
	handler   PeerHandler
	SzID      string
	E_key     []byte
	FE_key    []byte
	AccountID int
	whichKey  int
	alive     bool
}

func NewCNPeer(handler PeerHandler, conn net.Conn) *CNPeer {
	return &CNPeer{
		conn:      conn,
		handler:   handler,
		SzID:      "",
		E_key:     []byte(DEFAULT_KEY),
		FE_key:    nil,
		AccountID: -1,
		whichKey:  USE_E,
		alive:     true,
	}
}

func (peer *CNPeer) Send(typeID uint32, data ...interface{}) error {
	// grab buffer from pool
	buf := pool.Get()
	defer pool.Put(buf)

	// body start
	pkt := NewPacket(buf)

	// encode type id
	if err := pkt.Encode(typeID); err != nil {
		return err
	}

	// encode data
	for _, trailer := range data {
		if err := pkt.Encode(trailer); err != nil {
			return err
		}
	}

	// encrypt body
	switch peer.whichKey {
	case USE_E:
		EncryptData(buf.Bytes(), peer.E_key)
	case USE_FE:
		EncryptData(buf.Bytes(), peer.FE_key)
	}

	// write packet size
	if err := binary.Write(peer.conn, binary.LittleEndian, uint32(buf.Len())); err != nil {
		return err
	}

	// write packet body
	log.Printf("Sending %#v, sizeof: %d", data, buf.Len())
	if _, err := peer.conn.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("[FATAL] failed to write packet body! %v", err)
	}

	return nil
}

func (peer *CNPeer) Kill() {
	if !peer.alive {
		return
	}

	peer.alive = false
	peer.conn.Close()
	peer.handler.Disconnect(peer)
}

// meant to be invoked as a goroutine
func (peer *CNPeer) Handler() {
	defer peer.Kill()

	for {
		// read packet size, the goroutine spends most of it's time parked here
		var sz uint32
		if err := binary.Read(peer.conn, binary.LittleEndian, &sz); err != nil {
			log.Printf("[FATAL] failed to read packet size! %v\n", err)
			return
		}

		// client should never send a packet size outside of this range
		if sz > CN_PACKET_BUFFER_SIZE || sz < 4 {
			log.Printf("[FATAL] malicious packet size received! %d", sz)
			return
		}

		// grab buffer && read packet body
		if err := func() error { // we wrap this in a closure so we can easily defer the buffer return to pool
			buf := pool.Get()
			defer pool.Put(buf)
			if _, err := buf.ReadFrom(io.LimitReader(peer.conn, int64(sz))); err != nil {
				return fmt.Errorf("failed to read packet body! %v", err)
			}

			// decrypt
			DecryptData(buf.Bytes(), peer.E_key)
			pkt := NewPacket(buf)

			// create packet && read typeID
			var typeID uint32
			if err := pkt.Decode(&typeID); err != nil {
				return fmt.Errorf("failed to read packet type! %v", err)
			}

			// dispatch packet
			log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, sz)
			if err := peer.handler.HandlePacket(peer, typeID, pkt); err != nil {
				return err
			}

			return nil
		}(); err != nil {
			log.Printf("[FATAL] %v", err)
			return
		}
	}
}
