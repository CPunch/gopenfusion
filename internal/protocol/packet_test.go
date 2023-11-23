package protocol

import (
	"bytes"
	"testing"
)

type TestPacketData struct {
	A        int32
	B        int32
	UTF16Str string `size:"32"`
	Pad      int16  `pad:"2"`
	C        int32
}

// this is the data we expect to get from encoding the above struct with:
//
//	Encode(TestPacketData{
//		A:        1,
//		B:        2,
//		UTF16Str: "hello world",
//		C:        3,
//	})
var testData = [...]byte{0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x68, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x20, 0x00, 0x77, 0x00, 0x6f, 0x00, 0x72, 0x00, 0x6c, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}

func TestPacketEncodeDecode(t *testing.T) {
	pkt := NewPacket(&bytes.Buffer{})
	if err := pkt.Encode(TestPacketData{
		A:        1,
		B:        2,
		UTF16Str: "hello world",
		C:        3,
	}); err != nil {
		t.Error(err)
	}

	if !bytes.Equal(pkt.readWriter.(*bytes.Buffer).Bytes(), testData[:]) {
		t.Error("packet data does not match!")
	}

	var test TestPacketData
	if err := pkt.Decode(&test); err != nil {
		t.Error(err)
	}

	if test.A != 1 || test.B != 2 || test.C != 3 || test.UTF16Str != "hello world" {
		t.Error("decoded packet data does not match!")
	}
}
