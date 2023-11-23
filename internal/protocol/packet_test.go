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

func TestPacket(t *testing.T) {
	pkt := NewPacket(&bytes.Buffer{})
	if err := pkt.Encode(TestPacketData{
		A:        1,
		B:        2,
		UTF16Str: "hello world",
		C:        3,
	}); err != nil {
		t.Error(err)
	}

	var test TestPacketData
	if err := pkt.Decode(&test); err != nil {
		t.Error(err)
	}

	if test.A != 1 || test.B != 2 || test.C != 3 || test.UTF16Str != "hello world" {
		t.Error("packet data does not match!")
	}
}
