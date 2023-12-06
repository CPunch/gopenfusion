package login_test

import (
	"context"
	"encoding/binary"
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

var (
	testCharCreate = protocol.SP_CL2LS_REQ_CHAR_CREATE{
		PCStyle: protocol.SPCStyle{
			INameCheck: 1, SzFirstName: "Hector",
			SzLastName: "Bannonvenom", IGender: 1, IFaceStyle: 1,
			IHairStyle: 17, IHairColor: 11, ISkinColor: 10, IEyeColor: 2,
			IHeight: 1, IBody: 0, IClass: 0,
		},
		SOn_Item: protocol.SOnItem{
			IEquipHandID: 0, IEquipUBID: 53, IEquipLBID: 17, IEquipFootID: 58,
			IEquipHeadID: 0, IEquipFaceID: 0, IEquipBackID: 0,
		},
		SOn_Item_Index: protocol.SOnItem_Index{
			IEquipUBID_index: 15, IEquipLBID_index: 12, IEquipFootID_index: 17,
			IFaceStyle: 2, IHairStyle: 18,
		},
	}
)

func makeDummyPeer(ctx context.Context, is *is.I, recv chan<- *cnet.PacketEvent) *cnet.Peer {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", loginPort))
	is.NoErr(err)

	peer := cnet.NewPeer(ctx, conn)
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

// This test tries a typical login sequence.
func TestLoginSuccSequence(t *testing.T) {
	is := is.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	recv := make(chan *cnet.PacketEvent)
	peer := makeDummyPeer(ctx, is, recv)
	defer peer.Kill()

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
	_, err = testDB.TryLogin("testLoginSequence", "test")
	is.NoErr(err) // TryLogin() should not return an error
}

// This test tries a typical login sequence, but with an invalid password.
func TestLoginFailSequence(t *testing.T) {
	is := is.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	recv := make(chan *cnet.PacketEvent)
	peer := makeDummyPeer(ctx, is, recv)
	defer peer.Kill()

	// send login request (this should not create an account)
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

// This test tries a typical login sequence, and creates a character
func TestCharacterSequence(t *testing.T) {
	is := is.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	recv := make(chan *cnet.PacketEvent)
	peer := makeDummyPeer(ctx, is, recv)
	defer peer.Kill()

	// send login request (this should create an account)
	err := peer.Send(protocol.P_CL2LS_REQ_LOGIN, protocol.SP_CL2LS_REQ_LOGIN{
		SzID:       "TestCharacterSequence",
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
	is.Equal(resp.SzID, "TestCharacterSequence") // should have the same ID
	is.Equal(resp.ICharCount, int8(0))           // should have 0 characters

	// perform key swap
	peer.E_key = protocol.CreateNewKey(
		resp.UiSvrTime,
		uint64(resp.ICharCount+1),
		uint64(resp.ISlotNum+1),
	)
	peer.FE_key = protocol.CreateNewKey(
		binary.LittleEndian.Uint64([]byte(protocol.DEFAULT_KEY)),
		0,
		1,
	)

	// send character name check request
	err = peer.Send(protocol.P_CL2LS_REQ_SAVE_CHAR_NAME, protocol.SP_CL2LS_REQ_SAVE_CHAR_NAME{
		ISlotNum:    1,
		IGender:     1,
		IFNCode:     260,
		ILNCode:     551,
		IMNCode:     33,
		SzFirstName: testCharCreate.PCStyle.SzFirstName,
		SzLastName:  testCharCreate.PCStyle.SzLastName,
	})
	is.NoErr(err) // peer.Send() should not return an error

	// receive character creation response
	evnt = <-recv
	is.Equal(int(evnt.PktID), protocol.P_LS2CL_REP_SAVE_CHAR_NAME_SUCC) // should receive P_LS2CL_REP_SAVE_CHAR_NAME_SUCC

	var charResp protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_SUCC
	err = protocol.NewPacket(evnt.Pkt).Decode(&charResp)
	is.NoErr(err) // packet.Decode() should not return an error

	// verify response
	is.Equal(charResp.ISlotNum, int8(1))                               // should have the same slot number
	is.Equal(charResp.IGender, int8(1))                                // should have the same gender
	is.Equal(charResp.SzFirstName, testCharCreate.PCStyle.SzFirstName) // should have the same first name
	is.Equal(charResp.SzLastName, testCharCreate.PCStyle.SzLastName)   // should have the same last name

	// send character create request
	charCreate := testCharCreate
	charCreate.PCStyle.IPC_UID = charResp.IPC_UID
	err = peer.Send(protocol.P_CL2LS_REQ_CHAR_CREATE, charCreate)
	is.NoErr(err) // peer.Send() should not return an error

	// receive character create response
	evnt = <-recv
	is.Equal(int(evnt.PktID), protocol.P_LS2CL_REP_CHAR_CREATE_SUCC) // should receive P_LS2CL_REP_CHAR_CREATE_SUCC

	var charCreateResp protocol.SP_LS2CL_REP_CHAR_CREATE_SUCC
	err = protocol.NewPacket(evnt.Pkt).Decode(&charCreateResp)
	is.NoErr(err) // packet.Decode() should not return an error

	// verify response
	is.Equal(charCreate.PCStyle, charCreateResp.SPC_Style) // should have the same PCStyle
	is.Equal(charCreate.SOn_Item, charCreateResp.SOn_Item) // should have the same SOn_Item
}
