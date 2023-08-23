package shard

import "github.com/CPunch/gopenfusion/internal/protocol"

func (server *ShardServer) freeChat(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_SEND_FREECHAT_MESSAGE
	pkt.Decode(&chat)

	// sanity check
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC, protocol.SP_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC{
		IPC_ID:     int32(plr.PlayerID),
		SzFreeChat: chat.SzFreeChat,
		IEmoteCode: chat.IEmoteCode,
	})
}

func (server *ShardServer) menuChat(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_SEND_MENUCHAT_MESSAGE
	pkt.Decode(&chat)

	// sanity check
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC, protocol.SP_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC{
		IPC_ID:     int32(plr.PlayerID),
		SzFreeChat: chat.SzFreeChat,
		IEmoteCode: chat.IEmoteCode,
	})
}

func (server *ShardServer) emoteChat(peer *protocol.CNPeer, pkt protocol.Packet) error {
	var chat protocol.SP_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT
	pkt.Decode(&chat)

	// sanity check
	plr, err := server.getPlayer(peer)
	if err != nil {
		return err
	}

	// spread message
	return server.sendAllPacket(plr, protocol.P_FE2CL_REP_PC_AVATAR_EMOTES_CHAT, protocol.SP_FE2CL_REP_PC_AVATAR_EMOTES_CHAT{
		IID_From:   int32(plr.PlayerID),
		IEmoteCode: chat.IEmoteCode,
	})
}
