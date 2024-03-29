package cnet

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync/atomic"

	"github.com/CPunch/gopenfusion/cnet/protocol"
)

const (
	USE_E = iota
	USE_FE
)

type PacketEvent struct {
	Type  int
	Pkt   *bytes.Buffer
	PktID uint32
}

// Peer is a simple wrapper for net.Conn connections to send/recv packets over the Fusionfall packet protocol.
type Peer struct {
	uData    interface{}
	conn     net.Conn
	ctx      context.Context
	whichKey int
	alive    *atomic.Bool

	// May not be set while Send() or Handler() are concurrently running.
	E_key []byte

	// May not be set while Send() or Handler() are concurrently running.
	FE_key []byte
}

func NewPeer(ctx context.Context, conn net.Conn) *Peer {
	p := &Peer{
		conn:     conn,
		ctx:      ctx,
		whichKey: USE_E,
		alive:    &atomic.Bool{},

		E_key:  []byte(protocol.DEFAULT_KEY),
		FE_key: nil,
	}

	return p
}

func (peer *Peer) SetUserData(uData interface{}) {
	peer.uData = uData
}

func (peer *Peer) UserData() interface{} {
	return peer.uData
}

func (peer *Peer) Send(typeID uint32, data ...interface{}) error {
	// grab buffer from pool
	buf := protocol.GetBuffer()
	defer protocol.PutBuffer(buf)

	// allocate space for packet size
	buf.Write(make([]byte, 4))

	// body start
	pkt := protocol.NewPacket(buf)

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

	// prepend the packet size
	binary.LittleEndian.PutUint32(buf.Bytes()[:4], uint32(buf.Len()-4))

	// encrypt body
	var key []byte
	switch peer.whichKey {
	case USE_E:
		key = peer.E_key
	case USE_FE:
		key = peer.FE_key
	}
	protocol.EncryptData(buf.Bytes()[4:], key)

	// send full packet
	// log.Printf("Sending %#v, sizeof: %d, buffer: %v", data, buf.Len(), buf.Bytes())
	if _, err := peer.conn.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write packet body! %v", err)
	}
	return nil
}

func (peer *Peer) SetActiveKey(whichKey int) {
	peer.whichKey = whichKey
}

func (peer *Peer) Kill() {
	// de-bounce: only kill if alive
	if !peer.alive.CompareAndSwap(true, false) {
		return
	}

	peer.conn.Close()
}

// meant to be invoked as a goroutine
func (peer *Peer) Handler(eRecv chan<- *PacketEvent) error {
	defer func() {
		close(eRecv)
		peer.Kill()
	}()

	peer.alive.Store(true)
	for {
		select {
		case <-peer.ctx.Done():
			return nil
		default:
			// read packet size, the goroutine spends most of it's time parked here
			var sz uint32
			if err := binary.Read(peer.conn, binary.LittleEndian, &sz); err != nil {
				return err
			}

			// client should never send a packet size outside of this range
			if sz > protocol.CN_PACKET_BUFFER_SIZE || sz < 4 {
				return fmt.Errorf("invalid packet size: %d", sz)
			}

			// grab buffer && read packet body
			buf := protocol.GetBuffer()
			if _, err := buf.ReadFrom(io.LimitReader(peer.conn, int64(sz))); err != nil {
				return fmt.Errorf("failed to read packet body: %v", err)
			}

			// decrypt
			protocol.DecryptData(buf.Bytes(), peer.E_key)
			pkt := protocol.NewPacket(buf)

			// create packet && read pktID
			var pktID uint32
			if err := pkt.Decode(&pktID); err != nil {
				return fmt.Errorf("failed to read packet type! %v", err)
			}

			// dispatch packet
			// log.Printf("Got packet ID: %x, with a sizeof: %d\n", pktID, sz)
			eRecv <- &PacketEvent{Pkt: buf, PktID: pktID}
		}
	}
}
