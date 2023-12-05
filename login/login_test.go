package login_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/CPunch/gopenfusion/login"
	"github.com/alicebob/miniredis/v2"
	"github.com/bitcomplete/sqltestutil"
	"github.com/matryer/is"
)

var (
	loginSrv  *login.LoginServer
	loginPort int
	testDB    *db.DBHandler
	rh        *redis.RedisHandler
)

func makeDummyPeer(ctx context.Context, is *is.I, recv chan<- *cnet.PacketEvent) *cnet.Peer {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", loginPort))
	is.NoErr(err)

	peer := cnet.NewPeer(context.Background(), conn)
	go func() {
		peer.Handler(recv)
	}()

	return peer
}

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// spawn postgres container
	psql, err := sqltestutil.StartPostgresContainer(ctx, "15")
	if err != nil {
		panic(err)
	}

	// open db handler
	testDB, err = db.OpenFromConnectionString("postgres", psql.ConnectionString()+"?sslmode=disable")
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
	defer r.Close()

	// open redis handler
	rh, err = redis.OpenRedis(r.Addr())
	if err != nil {
		panic(err)
	}
	defer rh.Close()

	loginPort, err = cnet.RandomPort()
	if err != nil {
		panic(err)
	}

	// start login server
	loginSrv, err = login.NewLoginServer(ctx, testDB, rh, loginPort)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := loginSrv.Start(); err != nil {
			panic(err)
		}
	}()

	// wait for login server to start, then start tests
	<-loginSrv.Service().Started()
	ret := m.Run()
	cancel()
	<-loginSrv.Service().Stopped()
	os.Exit(ret)
}

func TestLoginSuccSequence(t *testing.T) {
	is := is.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	recv := make(chan *cnet.PacketEvent)
	peer := makeDummyPeer(ctx, is, recv)
	defer func() {
		peer.Kill()
	}()

	// send login request (this should create an account)
	err := peer.Send(protocol.P_CL2LS_REQ_LOGIN, protocol.SP_CL2LS_REQ_LOGIN{
		SzID:       "testLoginSequence",
		SzPassword: "test",
	})
	is.NoErr(err) // peer.Send() should not return an error

	// receive login response
	evnt := <-recv
	is.Equal(int(evnt.PktID), protocol.P_LS2CL_REP_LOGIN_SUCC) // should receive P_LS2CL_REP_LOGIN_SUCC

	var resp protocol.SP_LS2CL_REP_LOGIN_SUCC
	err = protocol.NewPacket(evnt.Pkt).Decode(&resp)
	is.NoErr(err) // packet.Decode() should not return an error

	// verify response
	is.Equal(resp.SzID, "testLoginSequence") // should have the same ID
	is.Equal(resp.ICharCount, int8(0))       // should have 0 characters

	// verify account was created
	acc, err := testDB.TryLogin("testLoginSequence", "test")
	is.NoErr(err)                            // GetAccount() should not return an error
	is.Equal(acc.Login, "testLoginSequence") // should have the same ID
}

func TestLoginFailSequence(t *testing.T) {
	is := is.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	recv := make(chan *cnet.PacketEvent)
	peer := makeDummyPeer(ctx, is, recv)
	defer peer.Kill()

	// send login request (this should create an account)
	err := peer.Send(protocol.P_CL2LS_REQ_LOGIN, protocol.SP_CL2LS_REQ_LOGIN{
		SzID:       "",
		SzPassword: "",
	})
	is.NoErr(err) // peer.Send() should not return an error

	// receive login response
	evnt := <-recv
	is.Equal(int(evnt.PktID), protocol.P_LS2CL_REP_LOGIN_FAIL) // should receive P_LS2CL_REP_LOGIN_FAIL

	var resp protocol.SP_LS2CL_REP_LOGIN_FAIL
	err = protocol.NewPacket(evnt.Pkt).Decode(&resp)
	is.NoErr(err) // packet.Decode() should not return an error

	// verify response
	is.Equal(resp.SzID, "")                                                    // should have the same ID
	is.Equal(resp.IErrorCode, int32(login.LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH)) // should respond with LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH
}
