package shard

import (
	"fmt"

	"github.com/CPunch/gopenfusion/cnpeer"
	"github.com/CPunch/gopenfusion/internal/entity"
	"github.com/CPunch/gopenfusion/internal/protocol"
)

func (server *ShardServer) freeChat(peer *cnpeer.CNPeer, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_SEND_FREECHAT_MESSAGE
	pkt.Decode(&chat)

	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("freeChat: plr is nil")
	}

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC, protocol.SP_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC{
		IPC_ID:     int32(plr.PlayerID),
		SzFreeChat: chat.SzFreeChat,
		IEmoteCode: chat.IEmoteCode,
	})
}

func (server *ShardServer) menuChat(peer *cnpeer.CNPeer, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_SEND_MENUCHAT_MESSAGE
	pkt.Decode(&chat)

	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("menuChat: plr is nil")
	}

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC, protocol.SP_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC{
		IPC_ID:     int32(plr.PlayerID),
		SzFreeChat: chat.SzFreeChat,
		IEmoteCode: chat.IEmoteCode,
	})
}

func (server *ShardServer) emoteChat(peer *cnpeer.CNPeer, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT
	pkt.Decode(&chat)

	plr, ok := peer.UserData().(*entity.Player)
	if !ok || plr == nil {
		return fmt.Errorf("emoteChat: plr is nil")
	}

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_PC_AVATAR_EMOTES_CHAT, protocol.SP_FE2CL_REP_PC_AVATAR_EMOTES_CHAT{
		IID_From:   int32(plr.PlayerID),
		IEmoteCode: chat.IEmoteCode,
	})
}
