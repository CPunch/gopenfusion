package testutil

import (
	"context"
	"fmt"
	"net"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/alicebob/miniredis/v2"
	"github.com/bitcomplete/sqltestutil"
	"github.com/matryer/is"
)

type DummyPeer struct {
	Recv chan *cnet.PacketEvent
	Peer *cnet.Peer
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

	return &DummyPeer{Recv: recv, Peer: peer}
}

// SendAndRecv sends a packet (sID & out), waits for the expected response (rID) and decodes it into in
func (dp *DummyPeer) SendAndRecv(is *is.I, sID, rID uint32, out, in interface{}) {
	// send out packet
	err := dp.Peer.Send(sID, out)
	is.NoErr(err) // peer.Send() should not return an error

	// receive response
	evnt := <-dp.Recv
	defer protocol.PutBuffer(evnt.Pkt)

	is.Equal(evnt.PktID, rID)                         // should receive expected type
	is.NoErr(protocol.NewPacket(evnt.Pkt).Decode(in)) // packet.Decode() should not return an error
}

// Kill closes the peer's connection
func (dp *DummyPeer) Kill() {
	dp.Peer.Kill()
}

// SetupEnvironment spawns a postgres container and returns a db and redis handler
// along with a cleanup function
func SetupEnvironment(ctx context.Context) (*db.DBHandler, *redis.RedisHandler, func()) {
	// spawn postgres container
	psql, err := sqltestutil.StartPostgresContainer(ctx, "15")
	if err != nil {
		panic(err)
	}

	// open db handler
	testDB, err := db.OpenFromConnectionString("postgres", psql.ConnectionString()+"?sslmode=disable")
	if err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	if err = testDB.Setup(); err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	// start miniredis
	r, err := miniredis.Run()
	if err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	// open redis handler
	rh, err := redis.OpenRedis(r.Addr())
	if err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	return testDB, rh, func() {
		psql.Shutdown(ctx)
		rh.Close()
		r.Close()
	}
}
