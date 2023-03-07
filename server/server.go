package server

import (
	"github.com/CPunch/GopenFusion/protocol"
)

type Server interface {
	HandlePacket(client *Client, typeID uint32, pkt *protocol.Packet)
	Disconnect(client *Client)
}
