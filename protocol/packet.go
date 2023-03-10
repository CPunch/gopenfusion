package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"unicode/utf16"
)

/*
	this file handles serializing (and deserializing) structs to alignment-strict c structures generated via `tools/genstructs.py`.
	see script for details on usage!
*/

type Packet struct {
	ByteOrder binary.ByteOrder
	Buf       []byte
}

func NewPacket(buf []byte) *Packet {
	pkt := &Packet{
		ByteOrder: binary.LittleEndian,
		Buf:       buf,
	}
	return pkt
}

func (pkt *Packet) writeRaw(data []byte) {
	pkt.Buf = append(pkt.Buf, data...)
}

func (pkt *Packet) Write(data []byte) (int, error) {
	pkt.writeRaw(data)

	if len(pkt.Buf) > CN_PACKET_BUFFER_SIZE {
		return 0, fmt.Errorf("Failed to write to packet, invalid size!")
	}

	return len(data), nil
}

func (pkt *Packet) writeByte(data byte) {
	pkt.Write([]byte{data})
}

func (pkt *Packet) readRaw(data []byte) (int, error) {
	sz := copy(data, pkt.Buf)
	pkt.Buf = pkt.Buf[sz:]

	if sz != len(data) {
		return sz, io.EOF
	}

	return sz, nil
}

func (pkt *Packet) Read(data []byte) (int, error) {
	if len(data) > len(pkt.Buf) {
		return 0, fmt.Errorf("Failed to read from packet, invalid size!")
	}

	return pkt.readRaw(data)
}

func (pkt *Packet) readByte() byte {
	data := pkt.Buf[0]
	pkt.Buf = pkt.Buf[1:]
	return data
}

func (pkt *Packet) encodeStructField(field reflect.StructField, value reflect.Value) {
	log.Printf("Encoding '%s'", field.Name)

	switch field.Type.Kind() {
	case reflect.String: // all strings in fusionfall packets are encoded as utf16, we'll need to encode it
		sz, err := strconv.Atoi(field.Tag.Get("size"))
		if err != nil {
			panic(fmt.Errorf("Failed to grab string 'size' tag!!"))
		}

		buf16 := utf16.Encode([]rune(value.String()))

		// len(buf16) needs to be the same size as sz
		if len(buf16) > sz {
			// truncate
			buf16 = buf16[:sz]
		} else {
			// grow
			for len(buf16) < sz {
				buf16 = append(buf16, 0)
			}
		}

		// write
		binary.Write(pkt, pkt.ByteOrder, buf16)
	default:
		pkt.Encode(value.Addr().Interface())
	}

	// write padding bytes
	pad, err := strconv.Atoi(field.Tag.Get("pad"))
	if err == nil {
		for i := 0; i < pad; i++ {
			pkt.writeByte(0)
		}
	}
}

func (pkt *Packet) Encode(data interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(data))

	switch rv.Kind() {
	case reflect.Struct:
		// walk through each struct fields
		sz := rv.NumField()
		for i := 0; i < sz; i++ {
			pkt.encodeStructField(rv.Type().Field(i), rv.Field(i))
		}
	default:
		// we pass everything else to go's binary package
		binary.Write(pkt, pkt.ByteOrder, data)
	}
}

func (pkt *Packet) decodeStructField(field reflect.StructField, value reflect.Value) {
	log.Printf("Decoding '%s'", field.Name)

	switch field.Type.Kind() {
	case reflect.String: // all strings in fusionfall packets are encoded as utf16, we'll need to decode it
		sz, err := strconv.Atoi(field.Tag.Get("size"))
		if err != nil {
			panic(fmt.Errorf("Failed to grab string 'size' tag!!"))
		}

		buf16 := make([]uint16, sz)
		binary.Read(pkt, pkt.ByteOrder, buf16)

		// find null terminator
		var realSize int
		for ; realSize < len(buf16); realSize++ {
			if buf16[realSize] == 0 {
				break
			}
		}

		value.SetString(string(utf16.Decode(buf16[:realSize])))
	default:
		pkt.Decode(value.Addr().Interface())
	}

	// read padding bytes
	pad, err := strconv.Atoi(field.Tag.Get("pad"))
	if err == nil {
		for i := 0; i < pad; i++ {
			pkt.readByte()
		}
	}
}

func (pkt *Packet) Decode(data interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(data))

	switch rv.Kind() {
	case reflect.Struct:
		// walk through each struct fields
		sz := rv.NumField()
		for i := 0; i < sz; i++ {
			pkt.decodeStructField(rv.Type().Field(i), rv.Field(i))
		}
	default:
		binary.Read(pkt, pkt.ByteOrder, data)
	}
}
