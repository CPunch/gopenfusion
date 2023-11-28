package shard

import (
	"fmt"

	"github.com/CPunch/gopenfusion/internal/entity"
	"github.com/CPunch/gopenfusion/internal/protocol"
)

func (server *ShardServer) freeChat(peer *protocol.CNPeer, _plr interface{}, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_SEND_FREECHAT_MESSAGE
	pkt.Decode(&chat)

	if _plr == nil {
		return fmt.Errorf("freeChat: _plr is nil")
	}
	plr := _plr.(*entity.Player)

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC, protocol.SP_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC{
		IPC_ID:     int32(plr.PlayerID),
		SzFreeChat: chat.SzFreeChat,
		IEmoteCode: chat.IEmoteCode,
	})
}

func (server *ShardServer) menuChat(peer *protocol.CNPeer, _plr interface{}, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_SEND_MENUCHAT_MESSAGE
	pkt.Decode(&chat)

	if _plr == nil {
		return fmt.Errorf("menuChat: _plr is nil")
	}
	plr := _plr.(*entity.Player)

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC, protocol.SP_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC{
		IPC_ID:     int32(plr.PlayerID),
		SzFreeChat: chat.SzFreeChat,
		IEmoteCode: chat.IEmoteCode,
	})
}

func (server *ShardServer) emoteChat(peer *protocol.CNPeer, _plr interface{}, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT
	pkt.Decode(&chat)

	if _plr == nil {
		return fmt.Errorf("emoteChat: _plr is nil")
	}
	plr := _plr.(*entity.Player)

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_PC_AVATAR_EMOTES_CHAT, protocol.SP_FE2CL_REP_PC_AVATAR_EMOTES_CHAT{
		IID_From:   int32(plr.PlayerID),
		IEmoteCode: chat.IEmoteCode,
	})
}
