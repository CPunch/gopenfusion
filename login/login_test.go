package login_test

import (
	"context"
	"encoding/binary"
	"os"
	"testing"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/CPunch/gopenfusion/internal/testutil"
	"github.com/CPunch/gopenfusion/login"
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

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup environment
	var closer func()
	testDB, rh, closer = testutil.SetupEnvironment(ctx)
	defer closer()

	var err error
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

	dummy := testutil.MakeDummyPeer(ctx, is, loginPort)
	defer dummy.Kill()

	// send login request (this should create an account)
	var resp protocol.SP_LS2CL_REP_LOGIN_SUCC
	dummy.SendAndRecv(is, protocol.P_CL2LS_REQ_LOGIN, protocol.P_LS2CL_REP_LOGIN_SUCC,
		protocol.SP_CL2LS_REQ_LOGIN{
			SzID:       "testLoginSequence",
			SzPassword: "test",
		}, &resp)

	// verify response
	is.Equal(resp.SzID, "testLoginSequence") // should have the same ID
	is.Equal(resp.ICharCount, int8(0))       // should have 0 characters

	// verify account was created
	_, err := testDB.TryLogin("testLoginSequence", "test")
	is.NoErr(err) // TryLogin() should not return an error
}

// This test tries a typical login sequence, but with an invalid password.
func TestLoginFailSequence(t *testing.T) {
	is := is.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dummy := testutil.MakeDummyPeer(ctx, is, loginPort)
	defer dummy.Kill()

	// send login request (this should not create an account)
	var resp protocol.SP_LS2CL_REP_LOGIN_FAIL
	dummy.SendAndRecv(is, protocol.P_CL2LS_REQ_LOGIN, protocol.P_LS2CL_REP_LOGIN_FAIL,
		protocol.SP_CL2LS_REQ_LOGIN{
			SzID:       "",
			SzPassword: "",
		}, &resp)

	// verify response
	is.Equal(resp.SzID, "")                                                    // should have the same ID
	is.Equal(resp.IErrorCode, int32(login.LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH)) // should respond with LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH
}

// This test tries a typical login sequence, and creates a character
func TestCharacterSequence(t *testing.T) {
	is := is.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dummy := testutil.MakeDummyPeer(ctx, is, loginPort)
	defer dummy.Kill()

	// send login request (this should create an account)
	var resp protocol.SP_LS2CL_REP_LOGIN_SUCC
	dummy.SendAndRecv(is, protocol.P_CL2LS_REQ_LOGIN, protocol.P_LS2CL_REP_LOGIN_SUCC,
		protocol.SP_CL2LS_REQ_LOGIN{
			SzID:       "testCharacterSequence",
			SzPassword: "test",
		}, &resp)

	// verify response
	is.Equal(resp.SzID, "testCharacterSequence") // should have the same ID
	is.Equal(resp.ICharCount, int8(0))           // should have 0 characters

	// perform key swap
	dummy.Peer.E_key = protocol.CreateNewKey(
		resp.UiSvrTime,
		uint64(resp.ICharCount+1),
		uint64(resp.ISlotNum+1),
	)
	dummy.Peer.FE_key = protocol.CreateNewKey(
		binary.LittleEndian.Uint64([]byte(protocol.DEFAULT_KEY)),
		0,
		1,
	)

	// send character name check request
	var charResp protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_SUCC
	dummy.SendAndRecv(is, protocol.P_CL2LS_REQ_SAVE_CHAR_NAME, protocol.P_LS2CL_REP_SAVE_CHAR_NAME_SUCC,
		protocol.SP_CL2LS_REQ_SAVE_CHAR_NAME{
			ISlotNum:    1,
			IGender:     1,
			IFNCode:     260,
			ILNCode:     551,
			IMNCode:     33,
			SzFirstName: testCharCreate.PCStyle.SzFirstName,
			SzLastName:  testCharCreate.PCStyle.SzLastName,
		}, &charResp)

	// verify response
	is.Equal(charResp.ISlotNum, int8(1))                               // should have the same slot number
	is.Equal(charResp.IGender, int8(1))                                // should have the same gender
	is.Equal(charResp.SzFirstName, testCharCreate.PCStyle.SzFirstName) // should have the same first name
	is.Equal(charResp.SzLastName, testCharCreate.PCStyle.SzLastName)   // should have the same last name

	// send character create request
	charCreate := testCharCreate
	charCreate.PCStyle.IPC_UID = charResp.IPC_UID
	var charCreateResp protocol.SP_LS2CL_REP_CHAR_CREATE_SUCC
	dummy.SendAndRecv(is, protocol.P_CL2LS_REQ_CHAR_CREATE, protocol.P_LS2CL_REP_CHAR_CREATE_SUCC,
		charCreate, &charCreateResp)

	// verify response
	is.Equal(charCreate.PCStyle, charCreateResp.SPC_Style) // should have the same PCStyle
	is.Equal(charCreate.SOn_Item, charCreateResp.SOn_Item) // should have the same SOn_Item
}
