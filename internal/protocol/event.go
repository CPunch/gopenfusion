package protocol

import "bytes"

const (
	EVENT_CLIENT_DISCONNECT = iota
	EVENT_CLIENT_CONNECT
	EVENT_CLIENT_PACKET
)

type Event struct {
	Type  int
	Peer  *CNPeer
	Pkt   *bytes.Buffer
	PktID uint32
}
