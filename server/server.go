package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/CPunch/GoFusion/api/protocol"
)

const (
	E_KEY      = "m@rQn~W#"
	KEY_LENGTH = 8
)

func encrypt_byte_change_A(ERSize int, data []byte) int {
	var num, num2, num3 int

	for num+ERSize <= len(data) {
		num4 := num + num3
		num5 := num + (ERSize - 1 - num3)

		b := data[num4]
		data[num4] = data[num5]
		data[num5] = b
		num += ERSize
		num3++
		if num3 > ERSize/2 {
			num3 = 0
		}
	}

	num2 = ERSize - (num + ERSize - len(data))
	return num + num2
}

func xorData(buff, key []byte, size int) {
	for i := 0; i < size; i++ {
		buff[i] ^= key[i%KEY_LENGTH]
	}
}

func decryptData(buff, key []byte) {
	ERSize := len(buff)%(KEY_LENGTH/2+1)*2 + KEY_LENGTH
	size := encrypt_byte_change_A(ERSize, buff)
	xorData(buff, key, size)
}

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	tmp := make([]byte, 4)
	for {
		// grab packet size
		if _, err := conn.Read(tmp); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet size! %v", err))
		}

		sz := int(binary.LittleEndian.Uint32(tmp))
		pkt := protocol.NewPacket(sz)

		if _, err := conn.Read(pkt.Buf); err != nil {
			panic(fmt.Errorf("[FATAL] failed to read packet body! %v", err))
		}

		decryptData(pkt.Buf, []byte(E_KEY))
		typeID := int(binary.LittleEndian.Uint32(pkt.Read(4)))
		sizeof := len(pkt.Buf)
		log.Printf("Got packet ID: %x\n", typeID)

		pkt.ResetCursor()
		switch typeID {
		case protocol.P_CL2LS_REQ_LOGIN:
			var loginPkt protocol.SP_CL2LS_REQ_LOGIN
			pkt.Decode(&loginPkt)
			log.Printf("Got packet: %#v, with a sizeof: %d", loginPkt, sizeof)
		default:
			log.Printf("[WARN] unsupported packet ID: %x\n", typeID)
		}
	}
}

func main() {
	server, err := net.Listen("tcp", ":23000")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Server hosted on 127.0.0.1:23000")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}
		go handle(conn)
	}
}
