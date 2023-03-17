package server

import (
	"github.com/CPunch/gopenfusion/protocol"
)

type PacketHandler func(peer *Peer, pkt protocol.Packet) error

func stubbedPacket(_ *Peer, _ protocol.Packet) error { /* stubbed */ return nil }
