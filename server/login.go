package server

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/db"
	"github.com/CPunch/gopenfusion/protocol"
	"github.com/CPunch/gopenfusion/util"
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

func (server *LoginServer) AcceptLogin(peer *Peer, SzID string, IClientVerC int32, ISlotNum int8, data []protocol.SP_LS2CL_REP_CHAR_INFO) error {
	peer.SzID = SzID

	resp := protocol.SP_LS2CL_REP_LOGIN_SUCC{
		SzID:          SzID,
		ICharCount:    int8(len(data)),
		ISlotNum:      ISlotNum,
		IPaymentFlag:  1,
		IOpenBetaFlag: 0,
		UiSvrTime:     uint64(time.Now().Unix()),
	}

	if err := peer.Send(protocol.P_LS2CL_REP_LOGIN_SUCC, resp); err != nil {
		return err
	}

	// swap keys
	peer.E_key = protocol.CreateNewKey(
		resp.UiSvrTime,
		uint64(resp.ICharCount+1),
		uint64(resp.ISlotNum+1),
	)
	peer.FE_key = protocol.CreateNewKey(
		binary.LittleEndian.Uint64([]byte(protocol.DEFAULT_KEY)),
		uint64(IClientVerC),
		1,
	)

	// send characters (if any)
	for i := 0; i < len(data); i++ {
		if err := peer.Send(protocol.P_LS2CL_REP_CHAR_INFO, &data[i]); err != nil {
			return err
		}
	}

	return nil
}

func (server *LoginServer) Login(peer *Peer, pkt protocol.Packet) error {
	var loginPkt protocol.SP_CL2LS_REQ_LOGIN
	pkt.Decode(&loginPkt)

	SendError := func(e int32) {
		peer.Send(protocol.P_LS2CL_REP_LOGIN_FAIL, protocol.SP_LS2CL_REP_LOGIN_FAIL{
			IErrorCode: e,
			SzID:       loginPkt.SzID,
		})
	}

	// client is resending a login packet??
	if peer.AccountID != -1 {
		SendError(LOGIN_ERROR)
		return fmt.Errorf("Out of order P_CL2LS_REQ_LOGIN!")
	}

	// attempt login
	account, err := db.DefaultDB.TryLogin(loginPkt.SzID, loginPkt.SzPassword)
	if err == db.LoginErrorInvalidID {
		// this is the default behavior, auto create the account if the ID isn't in use
		account, err = db.DefaultDB.NewAccount(loginPkt.SzID, loginPkt.SzPassword)
		if err != nil {
			// respond with a dummy login error packet so the client doesn't get softlocked
			SendError(LOGIN_DATABASE_ERROR)
			return err
		}
	} else if err == db.LoginErrorInvalidPassword {
		// respond with invalid password
		SendError(LOGIN_ID_AND_PASSWORD_DO_NOT_MATCH)
		return nil
	} else if err != nil { // wtf?
		SendError(LOGIN_DATABASE_ERROR)
		return err
	}

	// grab player data
	peer.AccountID = account.AccountID
	plrs, err := db.DefaultDB.GetPlayers(account.AccountID)
	if err != nil {
		SendError(LOGIN_DATABASE_ERROR)
		return err
	}

	// truncate plrs
	if len(plrs) > 3 {
		plrs = plrs[:4]
	}

	// build character list
	charInfo := [4]protocol.SP_LS2CL_REP_CHAR_INFO{}
	for i, plr := range plrs {
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
			return err
		}

		copy(info.AEquip[:], AEquip)
		charInfo[i] = info
	}

	return server.AcceptLogin(peer, loginPkt.SzID, loginPkt.IClientVerC, 1, charInfo[:len(plrs)])
}

func (server *LoginServer) CheckCharacterName(peer *Peer, pkt protocol.Packet) error {
	var charPkt protocol.SP_CL2LS_REQ_CHECK_CHAR_NAME
	pkt.Decode(&charPkt)

	// just auto accept, client resends this data later
	return peer.Send(protocol.P_LS2CL_REP_CHECK_CHAR_NAME_SUCC, protocol.SP_LS2CL_REP_CHECK_CHAR_NAME_SUCC{
		SzFirstName: charPkt.SzFirstName,
		SzLastName:  charPkt.SzLastName,
	})
}

func (server *LoginServer) SaveCharacterName(peer *Peer, pkt protocol.Packet) error {
	var charPkt protocol.SP_CL2LS_REQ_SAVE_CHAR_NAME
	pkt.Decode(&charPkt)

	if peer.AccountID == -1 {
		peer.Send(protocol.P_LS2CL_REP_SAVE_CHAR_NAME_FAIL, protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_FAIL{})
		return fmt.Errorf("Out of order P_LS2CL_REP_SAVE_CHAR_NAME_FAIL!")
	}

	// TODO: sanity check SzFirstName && SzLastName
	PlayerID, err := db.DefaultDB.NewPlayer(peer.AccountID, charPkt.SzFirstName, charPkt.SzLastName, int(charPkt.ISlotNum))
	if err != nil {
		peer.Send(protocol.P_LS2CL_REP_SAVE_CHAR_NAME_FAIL, protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_FAIL{})
		return err
	}

	return peer.Send(protocol.P_LS2CL_REP_SAVE_CHAR_NAME_SUCC, protocol.SP_LS2CL_REP_SAVE_CHAR_NAME_SUCC{
		IPC_UID:     int64(PlayerID),
		ISlotNum:    charPkt.ISlotNum,
		IGender:     charPkt.IGender,
		SzFirstName: charPkt.SzFirstName,
		SzLastName:  charPkt.SzLastName,
	})
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

	// TODO: sanity check items in SOn_Item; see db.FinishPlayer()
	return true
}

func SendFail(peer *Peer) error {
	if err := peer.Send(protocol.P_LS2CL_REP_SHARD_SELECT_FAIL, protocol.SP_LS2CL_REP_SHARD_SELECT_FAIL{
		IErrorCode: 2,
	}); err != nil {
		return err
	}

	return nil
}

func (server *LoginServer) CharacterCreate(peer *Peer, pkt protocol.Packet) error {
	var charPkt protocol.SP_CL2LS_REQ_CHAR_CREATE
	pkt.Decode(&charPkt)

	if !validateCharacterCreation(&charPkt) {
		return SendFail(peer)
	}

	if err := db.DefaultDB.FinishPlayer(&charPkt, peer.AccountID); err != nil {
		return SendFail(peer)
	}

	plr, err := db.DefaultDB.GetPlayer(int(charPkt.PCStyle.IPC_UID))
	if err != nil {
		return SendFail(peer)
	}

	PCStyle, PCStyle2 := util.Player2PCStyle(plr)
	return peer.Send(protocol.P_LS2CL_REP_CHAR_CREATE_SUCC, protocol.SP_LS2CL_REP_CHAR_CREATE_SUCC{
		ILevel:     int16(plr.Level),
		SPC_Style:  PCStyle,
		SPC_Style2: PCStyle2,
		SOn_Item:   charPkt.SOn_Item, // if items were faked, we don't really care since the db only stores the sanitized fields
	})
}

func (server *LoginServer) CharacterDelete(peer *Peer, pkt protocol.Packet) error {
	var charPkt protocol.SP_CL2LS_REQ_CHAR_DELETE
	pkt.Decode(&charPkt)

	slot, err := db.DefaultDB.DeletePlayer(int(charPkt.IPC_UID), peer.AccountID)
	if err != nil {
		return SendFail(peer)
	}

	return peer.Send(protocol.P_LS2CL_REP_CHAR_DELETE_SUCC, protocol.SP_LS2CL_REP_CHAR_DELETE_SUCC{
		ISlotNum: int8(slot),
	})
}

func (server *LoginServer) FinishTutorial(peer *Peer, pkt protocol.Packet) error {
	var charPkt protocol.SP_CL2LS_REQ_SAVE_CHAR_TUTOR
	pkt.Decode(&charPkt)

	if err := db.DefaultDB.FinishTutorial(int(charPkt.IPC_UID), peer.AccountID); err != nil {
		return SendFail(peer)
	}

	// no response
	return nil
}
