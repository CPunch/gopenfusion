package testutil

import (
	"context"
	"fmt"
	"net"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/matryer/is"
)

type DummyPeer struct {
	Recv chan *cnet.PacketEvent
	Peer *cnet.Peer
	is   *is.I
}

// MakeDummyPeer creates a new dummy peer and returns it
func MakeDummyPeer(ctx context.Context, is *is.I, port int) *DummyPeer {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	is.NoErr(err)

	recv := make(chan *cnet.PacketEvent)
	peer := cnet.NewPeer(ctx, conn)
	go func() {
		peer.Handler(recv)
	}()

	return &DummyPeer{Recv: recv, Peer: peer, is: is}
}

// SendAndRecv sends a packet (sID & out), waits for the expected response (rID) and decodes it into in
func (dp *DummyPeer) SendAndRecv(sID, rID uint32, out, in interface{}) {
	// send out packet
	err := dp.Peer.Send(sID, out)
	dp.is.NoErr(err) // peer.Send() should not return an error

	// receive response
	evnt := <-dp.Recv
	defer protocol.PutBuffer(evnt.Pkt)

	dp.is.Equal(evnt.PktID, rID)                         // should receive expected type
	dp.is.NoErr(protocol.NewPacket(evnt.Pkt).Decode(in)) // packet.Decode() should not return an error
}

// Kill closes the peer's connection
func (dp *DummyPeer) Kill() {
	dp.Peer.Kill()
}
