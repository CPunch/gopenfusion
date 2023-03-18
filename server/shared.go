package server

import (
	"github.com/CPunch/gopenfusion/protocol"
)

type PacketHandler func(peer *protocol.CNPeer, pkt protocol.Packet) error

func stubbedPacket(_ *protocol.CNPeer, _ protocol.Packet) error { /* stubbed */ return nil }
