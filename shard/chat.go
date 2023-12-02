package shard

import (
	"fmt"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/shard/entity"
)

func (server *ShardServer) freeChat(peer *cnet.Peer, pkt protocol.Packet) error {
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

func (server *ShardServer) menuChat(peer *cnet.Peer, pkt protocol.Packet) error {
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

func (server *ShardServer) emoteChat(peer *cnet.Peer, pkt protocol.Packet) error {
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
