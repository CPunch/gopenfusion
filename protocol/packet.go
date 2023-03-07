package protocol

import (
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"unicode/utf16"
	"unsafe"
)

/*
	this file handles serializing (and deserializing) structs to alignment-strict c structures
*/

type Packet struct {
	ByteOrder binary.ByteOrder
	Buf       []byte
	cursor    int // to keep track of things like member alignment
}

const PACK_ALIGN = 4

func NewPacket(buf []byte) *Packet {
	pkt := &Packet{
		ByteOrder: binary.LittleEndian,
		Buf:       buf,
		cursor:    0,
	}
	return pkt
}

func (pkt *Packet) ResetCursor() {
	pkt.cursor = 0
}

func (pkt *Packet) CompletePacket(typeID uint32) {
	tmp := make([]byte, 4)
	pkt.ByteOrder.PutUint32(tmp, typeID)
	pkt.Buf = append(tmp, pkt.Buf...)
}

func (pkt *Packet) writeRaw(data []byte) {
	pkt.Buf = append(pkt.Buf, data...)
	pkt.cursor += len(data)
}

func (pkt *Packet) Write(data []byte) {
	pkt.writeRaw(data)

	if len(pkt.Buf) > CN_PACKET_BUFFER_SIZE {
		panic(fmt.Errorf("Failed to write to packet, invalid size!"))
	}
}

func (pkt *Packet) writeByte(data byte) {
	pkt.Write([]byte{data})
}

func (pkt *Packet) readRaw(sz int) []byte {
	data := pkt.Buf[:sz]
	pkt.Buf = pkt.Buf[sz:]
	pkt.cursor += sz
	return data
}

func (pkt *Packet) Read(sz int) []byte {
	if sz > len(pkt.Buf) {
		panic(fmt.Errorf("Failed to read from packet, invalid size!"))
	}

	return pkt.readRaw(sz)
}

func (pkt *Packet) readByte() byte {
	return pkt.Read(1)[0]
}

func (pkt *Packet) encodeStructField(field reflect.StructField, value reflect.Value) {
	log.Printf("Encoding '%s', current cursor: %d", field.Name, len(pkt.Buf))

	switch field.Type.Kind() {
	case reflect.String: // all strings in fusionfall packets are encoded as utf16, we'll need to encode it
		sz, err := strconv.Atoi(field.Tag.Get("size"))
		if err != nil {
			panic(fmt.Errorf("Failed to grab string 'size' tag!!"))
		}

		sz *= 2
		buf16 := utf16.Encode([]rune(value.String()))
		buf := *(*[]byte)(unsafe.Pointer(&buf16))

		// len(buf) needs to be the same size as sz
		if len(buf) > sz {
			// truncate
			buf = buf[:sz]
		} else {
			// grow
			for len(buf) < sz {
				buf = append(buf, 0)
			}
		}

		log.Printf("sending %d", len(buf))
		pkt.Write(buf)
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
	case reflect.Slice: // (untested)
		sz := rv.Len()

		// encode data
		for i := 0; i < sz; i++ {
			elem := rv.Index(i).Addr()
			pkt.Encode(elem.Addr().Interface())
		}
	case reflect.Uint8:
		pkt.writeByte(byte(rv.Uint()))
	case reflect.Uint16:
		tmp := make([]byte, 2)
		pkt.ByteOrder.PutUint16(tmp, uint16(rv.Uint()))
		pkt.Write(tmp)
	case reflect.Uint32:
		tmp := make([]byte, 4)
		pkt.ByteOrder.PutUint32(tmp, uint32(rv.Uint()))
		pkt.Write(tmp)
	case reflect.Uint64:
		tmp := make([]byte, 8)
		pkt.ByteOrder.PutUint64(tmp, uint64(rv.Uint()))
		pkt.Write(tmp)
	// im using the PutUintX() api here because it's functionally the same and apparently PutIntX doesn't exist?
	case reflect.Int8:
		pkt.writeByte(byte(rv.Int()))
	case reflect.Int16:
		tmp := make([]byte, 2)
		pkt.ByteOrder.PutUint16(tmp, uint16(rv.Int()))
		pkt.Write(tmp)
	case reflect.Int32:
		tmp := make([]byte, 4)
		pkt.ByteOrder.PutUint32(tmp, uint32(rv.Int()))
		pkt.Write(tmp)
	case reflect.Int64:
		tmp := make([]byte, 8)
		pkt.ByteOrder.PutUint64(tmp, uint64(rv.Int()))
		pkt.Write(tmp)
	}
}

func (pkt *Packet) decodeStructField(field reflect.StructField, value reflect.Value) {
	log.Printf("Decoding '%s', current cursor: %d", field.Name, pkt.cursor)

	switch field.Type.Kind() {
	case reflect.String: // all strings in fusionfall packets are encoded as utf16, we'll need to decode it
		sz, err := strconv.Atoi(field.Tag.Get("size"))
		if err != nil {
			panic(fmt.Errorf("Failed to grab string 'size' tag!!"))
		}

		// sz * sizeof(char16_t)
		sz *= 2
		buf := make([]byte, 0, sz)
		buf = append(buf, pkt.Read(sz)...)

		buf16 := *(*[]uint16)(unsafe.Pointer(&buf))

		// find null terminator
		var realSize int
		for ; realSize < len(buf16); realSize++ {
			if buf16[realSize] == 0 {
				break
			}
		}

		str := string(utf16.Decode(buf16[:realSize]))
		log.Printf(" got: %#.v or '%s'", buf16, str)
		value.SetString(str)
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
	case reflect.Array: // (untested)
		sz := rv.Len()

		// decode data
		log.Print("reading array length ", sz, " cursor: ", pkt.cursor)
		for i := 0; i < sz; i++ {
			elem := rv.Index(i)
			pkt.Decode(elem.Addr().Interface())
		}
	case reflect.Uint8:
		rv.SetUint(uint64(pkt.readByte()))
	case reflect.Uint16:
		rv.SetUint(uint64(pkt.ByteOrder.Uint16(pkt.Read(2))))
	case reflect.Uint32:
		rv.SetUint(uint64(pkt.ByteOrder.Uint32(pkt.Read(4))))
	case reflect.Uint64:
		rv.SetUint(pkt.ByteOrder.Uint64(pkt.Read(8)))
	case reflect.Int8:
		rv.SetInt(int64(pkt.readByte()))
	case reflect.Int16:
		rv.SetInt(int64(pkt.ByteOrder.Uint16(pkt.Read(2))))
	case reflect.Int32:
		rv.SetInt(int64(pkt.ByteOrder.Uint32(pkt.Read(4))))
	case reflect.Int64:
		rv.SetInt(int64(pkt.ByteOrder.Uint64(pkt.Read(8))))
	}
}
