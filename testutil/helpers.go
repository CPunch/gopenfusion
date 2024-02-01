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

func SetupEnvironment(ctx context.Context) (*db.DBHandler, *redis.RedisHandler, func()) {
	// spawn postgres container
	psql, err := sqltestutil.StartPostgresContainer(ctx, "15")
	if err != nil {
		panic(err)
	}

	// open db handler
	testDB, err := db.OpenFromConnectionString("postgres", psql.ConnectionString()+"?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = testDB.Setup(); err != nil {
		panic(err)
	}

	// start miniredis
	r, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	// open redis handler
	rh, err := redis.OpenRedis(r.Addr())
	if err != nil {
		panic(err)
	}

	return testDB, rh, func() {
		psql.Shutdown(ctx)
		rh.Close()
		r.Close()
	}
}

func MakeDummyPeer(ctx context.Context, is *is.I, port int, recv chan<- *cnet.PacketEvent) *cnet.Peer {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	is.NoErr(err)

	peer := cnet.NewPeer(ctx, conn)
	go func() {
		peer.Handler(recv)
	}()

	return peer
}

func SendAndRecv(peer *cnet.Peer, recv chan *cnet.PacketEvent, is *is.I, sID, rID uint32, out, in interface{}) {
	// send out packet
	err := peer.Send(sID, out)
	is.NoErr(err) // peer.Send() should not return an error

	// receive response
	evnt := <-recv
	defer protocol.PutBuffer(evnt.Pkt)

	is.Equal(evnt.PktID, rID)                         // should receive expected type
	is.NoErr(protocol.NewPacket(evnt.Pkt).Decode(in)) // packet.Decode() should not return an error
}
