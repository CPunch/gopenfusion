package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/CPunch/GopenFusion/protocol"
)

const (
	USE_E = iota
	USE_FE
)

type Client struct {
	server   *Server
	conn     net.Conn
	e_key    []byte
	fe_key   []byte
	whichKey int
}

func newClient(server *Server, conn net.Conn) *Client {
	return &Client{
		server:   server,
		conn:     conn,
		e_key:    []byte(protocol.DEFAULT_KEY),
		whichKey: USE_E,
	}
}

func (client *Client) Send(data interface{}, typeID uint32) {
	// encode
	pkt := protocol.NewPacket(make([]byte, 0))
	pkt.Encode(data)
	log.Printf("Sending %#v, sizeof: %d", data, len(pkt.Buf))

	// write packet size
	tmp := make([]byte, 4)
	binary.LittleEndian.PutUint32(tmp, uint32(len(pkt.Buf)+4))
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet size! %v", err))
	}

	// prepend the typeID to the packet body
	binary.LittleEndian.PutUint32(tmp, uint32(typeID))
	tmp = append(tmp, pkt.Buf...)

	// encrypt typeID & body
	switch client.whichKey {
	case USE_E:
		protocol.EncryptData(tmp, client.e_key)
	case USE_FE:
		protocol.EncryptData(tmp, client.fe_key)
	}

	// write packet body
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet body! %v", err))
	}
	log.Printf("sent!")
}

func (client *Client) Login(pkt *protocol.Packet) {
	var loginPkt protocol.SP_CL2LS_REQ_LOGIN
	pkt.Decode(&loginPkt)
	log.Printf("Got packet: %#v", loginPkt)

	// !! TODO
	resp := &protocol.SP_LS2CL_REP_LOGIN_SUCC{
		SZID:          loginPkt.SZID,
		ICharCount:    0,
		ISlotNum:      1,
		IPaymentFlag:  1,
		IOpenBetaFlag: 0,
		UISvrTime:     uint64(time.Now().Unix()),
	}

	client.Send(resp, protocol.P_LS2CL_REP_LOGIN_SUCC)
	client.e_key = protocol.CreateNewKey(
		resp.UISvrTime,
		uint64(resp.ICharCount+1),
		uint64(resp.ISlotNum+1),
	)
	client.fe_key = protocol.CreateNewKey(
		binary.LittleEndian.Uint64([]byte(protocol.DEFAULT_KEY)),
		uint64(loginPkt.IClientVerC),
		1,
	)
}

func (client *Client) ClientHandler() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %p panic'd! %v", client, err)
		}
		client.conn.Close()
		client.server.unregister <- client
	}()

	tmp := make([]byte, 4, protocol.CN_PACKET_BUFFER_SIZE)
	for {
		// read packet size
		if _, err := client.conn.Read(tmp); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet size! %v", err))
		}
		sz := int(binary.LittleEndian.Uint32(tmp))

		// client should never send a packet size outside of this range
		if sz > protocol.CN_PACKET_BUFFER_SIZE || sz < 4 {
			panic(fmt.Errorf("[FATAL] malicious packet size received! %d", sz))
		}

		// read packet body
		if _, err := client.conn.Read(tmp[:sz]); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet body! %v", err))
		}

		// decrypt && grab typeID
		protocol.DecryptData(tmp[:sz], client.e_key)
		typeID := int(binary.LittleEndian.Uint32(tmp[:4]))

		log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, sz)
		pkt := protocol.NewPacket(tmp[4:sz])
		switch typeID {
		case protocol.P_CL2LS_REQ_LOGIN:
			client.Login(pkt)
		default:
			log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
		}

		// reset tmp
		tmp = tmp[:4]
	}
}
