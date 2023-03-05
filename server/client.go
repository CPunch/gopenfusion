package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/CPunch/GoFusion/protocol"
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
	tmp := make([]byte, 4)

	// encode
	pkt := protocol.NewPacket(0)
	pkt.Encode(data)

	// write packet size
	log.Printf("Sending %#v, sizeof: %d", data, len(pkt.Buf))
	binary.LittleEndian.PutUint32(tmp, uint32(len(pkt.Buf)+4))
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet size! %v", err))
	}

	// prepend the typeID to the packet body
	binary.LittleEndian.PutUint32(tmp, uint32(typeID))
	pkt.Buf = append(tmp, pkt.Buf...)

	// encrypt body
	protocol.EncryptData(pkt.Buf[:], client.key)

	// write packet body
	if _, err := client.conn.Write(pkt.Buf); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet body! %v", err))
	}
	log.Printf("sent!")
}

func (client *Client) ClientHandler() {
	defer func() {
		client.conn.Close()
		client.server.unregister <- client
	}()

	tmp := make([]byte, 4)
	for {
		// read packet size
		if _, err := client.conn.Read(tmp); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet size! %v", err))
		}
		sz := int(binary.LittleEndian.Uint32(tmp))

		// read packet body
		pkt := protocol.NewPacket(sz)
		if _, err := client.conn.Read(pkt.Buf); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet body! %v", err))
		}

		// decrypt && grab typeID
		protocol.DecryptData(pkt.Buf[:], client.key)
		typeID := int(binary.LittleEndian.Uint32(pkt.Read(4)))

		log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, len(pkt.Buf))
		pkt.ResetCursor()
		switch typeID {
		case protocol.P_CL2LS_REQ_LOGIN:
			var loginPkt protocol.SP_CL2LS_REQ_LOGIN
			pkt.Decode(&loginPkt)
			log.Printf("Got packet: %#v", loginPkt)

			client.Send(&protocol.SP_LS2CL_REP_LOGIN_FAIL{ErrorCode: protocol.LOGIN_FAIL_VERSION_ERROR, ID: loginPkt.ID}, protocol.P_LS2CL_REP_LOGIN_FAIL)
		default:
			log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
		}
	}
}
