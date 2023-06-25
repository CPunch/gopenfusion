package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/CPunch/gopenfusion/core/protocol/pool"
)

const (
	USE_E = iota
	USE_FE
)

// CNPeer is a simple wrapper for net.Conn connections to send/recv packets over the Fusionfall packet protocol.
type CNPeer struct {
	conn      net.Conn
	eRecv     chan *Event
	SzID      string
	E_key     []byte
	FE_key    []byte
	AccountID int
	PlayerID  int32
	whichKey  int
	alive     bool
}

func NewCNPeer(eRecv chan *Event, conn net.Conn) *CNPeer {
	return &CNPeer{
		conn:      conn,
		eRecv:     eRecv,
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

func (peer *CNPeer) SetActiveKey(whichKey int) {
	peer.whichKey = whichKey
}

func (peer *CNPeer) Kill() {
	log.Printf("Killing peer %p", peer)
	if !peer.alive {
		return
	}

	peer.alive = false
	peer.conn.Close()
	peer.eRecv <- &Event{Type: EVENT_CLIENT_DISCONNECT, Peer: peer}
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
		if err := func() error {
			buf := pool.Get()
			if _, err := buf.ReadFrom(io.LimitReader(peer.conn, int64(sz))); err != nil {
				return fmt.Errorf("failed to read packet body! %v", err)
			}

			// decrypt
			DecryptData(buf.Bytes(), peer.E_key)
			pkt := NewPacket(buf)

			// create packet && read pktID
			var pktID uint32
			if err := pkt.Decode(&pktID); err != nil {
				return fmt.Errorf("failed to read packet type! %v", err)
			}

			// dispatch packet
			log.Printf("Got packet ID: %x, with a sizeof: %d\n", pktID, sz)
			peer.eRecv <- &Event{Type: EVENT_CLIENT_PACKET, Peer: peer, Pkt: buf, PktID: pktID}
			return nil
		}(); err != nil {
			log.Printf("[FATAL] %v", err)
			return
		}
	}
}
