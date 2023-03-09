package server

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/CPunch/GopenFusion/config"
	"github.com/CPunch/GopenFusion/db"
	"github.com/CPunch/GopenFusion/protocol"
	"github.com/CPunch/GopenFusion/util"
)

const (
	LOGIN_DATABASE_ERROR                        = 0
	LOGIN_ID_DOESNT_EXIST                       = 1
	LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH          = 2
	LOGIN_ID_ALREADY_IN_USE                     = 3
	LOGIN_ERROR                                 = 4
	LOGIN_CLIENT_VERSION_OUTDATED               = 6
	LOGIN_YOU_ARE_NOT_AN_AUTHORIZED_BETA_TESTER = 7
	LOGIN_AUTHENTICATION_CONNECTION_ERROR       = 8
	LOGIN_UPDATED_EUALA_REQUIRED                = 9
)

func (server *LoginServer) AcceptLogin(client *Client, SzID string, IClientVerC int32, ISlotNum int8, data []protocol.SP_LS2CL_REP_CHAR_INFO) {
	client.SzID = SzID

	resp := &protocol.SP_LS2CL_REP_LOGIN_SUCC{
		SzID:          SzID,
		ICharCount:    int8(len(data)),
		ISlotNum:      ISlotNum,
		IPaymentFlag:  1,
		IOpenBetaFlag: 0,
		UiSvrTime:     uint64(time.Now().Unix()),
	}

	client.Send(resp, protocol.P_LS2CL_REP_LOGIN_SUCC)
	client.E_key = protocol.CreateNewKey(
		resp.UiSvrTime,
		uint64(resp.ICharCount+1),
		uint64(resp.ISlotNum+1),
	)
	client.FE_key = protocol.CreateNewKey(
		binary.LittleEndian.Uint64([]byte(protocol.DEFAULT_KEY)),
		uint64(IClientVerC),
		1,
	)

	// send characters (if any)
	for i := 0; i < len(data); i++ {
		client.Send(&data[i], protocol.P_LS2CL_REP_CHAR_INFO)
	}
}

func (server *LoginServer) Login(client *Client, pkt *protocol.Packet) {
	var loginPkt protocol.SP_CL2LS_REQ_LOGIN
	pkt.Decode(&loginPkt)

	SendError := func(e int32) {
		client.Send(&protocol.SP_LS2CL_REP_LOGIN_FAIL{
			IErrorCode: e,
			SzID:       loginPkt.SzID,
		}, protocol.P_LS2CL_REP_LOGIN_FAIL)
	}

	// client is resending a login packet??
	if client.AccountID != -1 {
		SendError(LOGIN_ERROR)
		panic(fmt.Errorf("Out of order P_CL2LS_REQ_LOGIN!"))
	}

	// attempt login
	account, err := db.DefaultDB.TryLogin(loginPkt.SzID, loginPkt.SzPassword)
	if err == db.LoginErrorInvalidID {
		// this is the default behavior, auto create the account if the ID isn't in use
		account, err = db.DefaultDB.NewAccount(loginPkt.SzID, loginPkt.SzPassword)
		if err != nil {
			// respond with a dummy login error packet so the client doesn't get softlocked
			SendError(LOGIN_DATABASE_ERROR)
			panic(err)
		}
	} else if err == db.LoginErrorInvalidPassword {
		// respond with invalid password
		SendError(LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH)
		return
	} else if err != nil { // wtf?
		SendError(LOGIN_DATABASE_ERROR)
		panic(err)
	}

	// grab player data
	client.AccountID = account.AccountID
	plrs, err := db.DefaultDB.GetPlayers(account.AccountID)
	if err != nil {
		SendError(LOGIN_DATABASE_ERROR)
		panic(err)
	}

	// build character list
	charInfo := make([]protocol.SP_LS2CL_REP_CHAR_INFO, 0)
	for _, plr := range plrs {
		PCStyle, PCStyle2 := util.Player2PCStyle(&plr)
		info := protocol.SP_LS2CL_REP_CHAR_INFO{
			ISlot:      int8(plr.Slot),
			ILevel:     int16(plr.Level),
			SPC_Style:  PCStyle,
			SPC_Style2: PCStyle2,
			IX:         int32(plr.XCoordinate),
			IY:         int32(plr.YCoordinate),
			IZ:         int32(plr.ZCoordinate),
		}

		AEquip, err := db.DefaultDB.GetPlayerInventorySlots(plr.PlayerID, 0, config.AEQUIP_COUNT-1)
		if err != nil {
			SendError(LOGIN_DATABASE_ERROR)
			panic(err)
		}

		copy(info.AEquip[:], AEquip)
		charInfo = append(charInfo, info)
	}

	server.AcceptLogin(client, loginPkt.SzID, loginPkt.IClientVerC, 1, charInfo)
}

func (server *LoginServer) SaveCharacterName(client *Client, pkt *protocol.Packet) {
	var charPkt protocol.SP_CL2LS_REQ_SAVE_CHAR_NAME
	pkt.Decode(&charPkt)

	if client.AccountID == -1 {
		client.Send(&protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_FAIL{}, protocol.P_LS2CL_REP_SAVE_CHAR_NAME_FAIL)
		panic(fmt.Errorf("Out of order P_LS2CL_REP_SAVE_CHAR_NAME_FAIL!"))
	}

	PlayerID, err := db.DefaultDB.NewPlayer(client.AccountID, charPkt.SzFirstName, charPkt.SzLastName, int(charPkt.ISlotNum))
	if err != nil {
		client.Send(&protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_FAIL{}, protocol.P_LS2CL_REP_SAVE_CHAR_NAME_FAIL)
		panic(err)
	}

	client.Send(&protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_SUCC{
		IPC_UID:     int64(PlayerID),
		ISlotNum:    charPkt.ISlotNum,
		IGender:     charPkt.IGender,
		SzFirstName: charPkt.SzFirstName,
		SzLastName:  charPkt.SzLastName,
	}, protocol.P_LS2CL_REP_SAVE_CHAR_NAME_SUCC)
}

func validateCharacterCreation(character *protocol.SP_CL2LS_REQ_CHAR_CREATE) bool {
	// thanks openfusion!
	// all the values have been determined from analyzing client code and xdt
	// and double checked using cheat engine

	// check base parameters
	style := &character.PCStyle
	if !(style.IBody >= 0 && style.IBody <= 2 &&
		style.IEyeColor >= 1 && style.IEyeColor <= 5 &&
		style.IGender >= 1 && style.IGender <= 2 &&
		style.IHairColor >= 1 && style.IHairColor <= 18) &&
		style.IHeight >= 0 && style.IHeight <= 4 &&
		style.INameCheck >= 0 && style.INameCheck <= 2 &&
		style.ISkinColor >= 1 && style.ISkinColor <= 12 {
		return false
	}

	// facestyle and hairstyle are gender dependent
	if !(style.IGender == 1 && style.IFaceStyle >= 1 && style.IFaceStyle <= 5 && style.IHairStyle >= 1 && style.IHairStyle <= 23) &&
		!(style.IGender == 2 && style.IFaceStyle >= 6 && style.IFaceStyle <= 10 && style.IHairStyle >= 25 && style.IHairStyle <= 45) {
		return false
	}

	return true
}

func (server *LoginServer) CharacterCreate(client *Client, pkt *protocol.Packet) {
	var charPkt protocol.SP_CL2LS_REQ_CHAR_CREATE
	pkt.Decode(&charPkt)

	if !validateCharacterCreation(&charPkt) {
		client.Send(&protocol.SP_LS2CL_REP_SHARD_SELECT_FAIL{IErrorCode: 2}, protocol.P_LS2CL_REP_SHARD_SELECT_FAIL)
		panic(fmt.Errorf("invalid SP_CL2LS_REQ_CHAR_CREATE!"))
	}

	if err := db.DefaultDB.FinishPlayer(&charPkt, client.AccountID); err != nil {
		client.Send(&protocol.SP_LS2CL_REP_SHARD_SELECT_FAIL{IErrorCode: 2}, protocol.P_LS2CL_REP_SHARD_SELECT_FAIL)
		panic(err)
	}

	plr, err := db.DefaultDB.GetPlayer(int(charPkt.PCStyle.IPC_UID))
	if err != nil {
		client.Send(&protocol.SP_LS2CL_REP_SHARD_SELECT_FAIL{IErrorCode: 2}, protocol.P_LS2CL_REP_SHARD_SELECT_FAIL)
		panic(err)
	}

	PCStyle, PCStyle2 := util.Player2PCStyle(plr)
	client.Send(&protocol.SP_LS2CL_REP_CHAR_CREATE_SUCC{
		ILevel:     int16(plr.Level),
		SPC_Style:  PCStyle,
		SPC_Style2: PCStyle2,
		SOn_Item:   protocol.SOnItem{ /*TODO*/ },
	}, protocol.P_LS2CL_REP_CHAR_CREATE_SUCC)
}
