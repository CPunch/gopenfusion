package server

import "github.com/CPunch/gopenfusion/protocol"

func (server *ShardServer) RequestEnter(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var enter protocol.SP_CL2FE_REQ_PC_ENTER
	pkt.Decode(&enter)

	return nil
}
