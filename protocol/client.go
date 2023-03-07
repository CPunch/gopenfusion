package protocol

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	USE_E = iota
	USE_FE
)

type ClientHandler interface {
	HandlePacket(client *Client, typeID uint32, pkt *Packet)
	Disconnect(client *Client)
}

type Client struct {
	handler  ClientHandler
	conn     net.Conn
	e_key    []byte
	fe_key   []byte
	whichKey int
}

func NewClient(handler ClientHandler, conn net.Conn) *Client {
	return &Client{
		handler:  handler,
		conn:     conn,
		e_key:    []byte(DEFAULT_KEY),
		whichKey: USE_E,
	}
}

func (client *Client) Send(data interface{}, typeID uint32) {
	// encode
	pkt := NewPacket(make([]byte, 0))
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
		EncryptData(tmp, client.e_key)
	case USE_FE:
		EncryptData(tmp, client.fe_key)
	}

	// write packet body
	if _, err := client.conn.Write(tmp); err != nil {
		panic(fmt.Errorf("[FATAL] failed to write packet body! %v", err))
	}
}

func (client *Client) AcceptLogin(SZID string, IClientVerC int32, ISlotNum int8, data []SP_LS2CL_REP_CHAR_INFO) {
	resp := &SP_LS2CL_REP_LOGIN_SUCC{
		SzID:          SZID,
		ICharCount:    int8(len(data)),
		ISlotNum:      ISlotNum,
		IPaymentFlag:  1,
		IOpenBetaFlag: 0,
		UiSvrTime:     uint64(time.Now().Unix()),
	}

	client.Send(resp, P_LS2CL_REP_LOGIN_SUCC)
	client.e_key = CreateNewKey(
		resp.UiSvrTime,
		uint64(resp.ICharCount+1),
		uint64(resp.ISlotNum+1),
	)
	client.fe_key = CreateNewKey(
		binary.LittleEndian.Uint64([]byte(DEFAULT_KEY)),
		uint64(IClientVerC),
		1,
	)

	// send characters (if any)
	for i := 0; i < len(data); i++ {
		client.Send(data[i], P_LS2CL_REP_CHAR_INFO)
	}
}

func (client *Client) ClientHandler() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Client %p panic'd! %v", client, err)
		}
		client.conn.Close()
		client.handler.Disconnect(client)
	}()

	tmp := make([]byte, 4, CN_PACKET_BUFFER_SIZE)
	for {
		// read packet size
		if _, err := client.conn.Read(tmp); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet size! %v", err))
		}
		sz := int(binary.LittleEndian.Uint32(tmp))

		// client should never send a packet size outside of this range
		if sz > CN_PACKET_BUFFER_SIZE || sz < 4 {
			panic(fmt.Errorf("[FATAL] malicious packet size received! %d", sz))
		}

		// read packet body
		if _, err := client.conn.Read(tmp[:sz]); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet body! %v", err))
		}

		// decrypt && grab typeID
		DecryptData(tmp[:sz], client.e_key)
		typeID := uint32(binary.LittleEndian.Uint32(tmp[:4]))

		// dispatch packet
		log.Printf("Got packet ID: %x, with a sizeof: %d\n", typeID, sz)
		pkt := NewPacket(tmp[4:sz])
		client.handler.HandlePacket(client, typeID, pkt)

		// reset tmp
		tmp = tmp[:4]
	}
}
