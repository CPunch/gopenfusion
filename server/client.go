package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/CPunch/GopenFusion/protocol"
)

type Client struct {
	server *Server
	conn   net.Conn
	key    []byte
}

func newClient(server *Server, conn net.Conn, key []byte) *Client {
	return &Client{
		server: server,
		conn:   conn,
		key:    key,
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
	protocol.EncryptData(tmp, client.key)

	// write packet body
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet body! %v", err))
	}
	log.Printf("sent!")
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
		protocol.DecryptData(tmp[:sz], client.key)
		typeID := int(binary.LittleEndian.Uint32(tmp[:4]))

		log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, sz)
		pkt := protocol.NewPacket(tmp[4:sz])
		switch typeID {
		case protocol.P_CL2LS_REQ_LOGIN:
			var loginPkt protocol.SP_CL2LS_REQ_LOGIN
			pkt.Decode(&loginPkt)
			log.Printf("Got packet: %#v", loginPkt)

			client.Send(&protocol.SP_LS2CL_REP_LOGIN_FAIL{
				ErrorCode: protocol.LOGIN_FAIL_VERSION_ERROR,
				ID:        loginPkt.ID,
			}, protocol.P_LS2CL_REP_LOGIN_FAIL)
		default:
			log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
		}

		// reset tmp
		tmp = tmp[:4]
	}
}
