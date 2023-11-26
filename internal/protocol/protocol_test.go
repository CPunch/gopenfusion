package protocol_test

import (
	"bytes"
	"testing"

	"github.com/CPunch/gopenfusion/internal/protocol"
)

type TestPacketData struct {
	A        int32
	B        int32
	UTF16Str string `size:"32"`
	Pad      int16  `pad:"2"`
	C        int32
}

var (
	testStruct = TestPacketData{
		A:        1,
		B:        2,
		UTF16Str: "hello world",
		C:        3,
	}

	// this is the data we expect to get from encoding the above struct
	testData = [...]byte{
		0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
		0x68, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00,
		0x6f, 0x00, 0x20, 0x00, 0x77, 0x00, 0x6f, 0x00,
		0x72, 0x00, 0x6c, 0x00, 0x64, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
	}

	// this is the data we expect to get from EncryptData(testData, []byte(protocol.DEFAULT_KEY))
	encTestData = []byte{
		0x23, 0x40, 0x72, 0x51, 0x6c, 0x7e, 0x57, 0x6c,
		0x05, 0x3b, 0x17, 0x51, 0x02, 0x7e, 0x40, 0x23,
		0x02, 0x40, 0x7e, 0x51, 0x19, 0x52, 0x38, 0x23,
		0x1f, 0x40, 0x1e, 0x0a, 0x51, 0x7e, 0x57, 0x23,
		0x6d, 0x40, 0x72, 0x6e, 0x51, 0x7e, 0x57, 0x23,
		0x23, 0x40, 0x72, 0x51, 0x6e, 0x7e, 0x57, 0x6d,
		0x6d, 0x57, 0x72, 0x51, 0x6e, 0x7e, 0x40, 0x23,
		0x6d, 0x40, 0x7e, 0x51, 0x6e, 0x72, 0x57, 0x23,
		0x6d, 0x40, 0x72, 0x6e, 0x51, 0x7e, 0x57, 0x23,
		0x6d, 0x40, 0x72, 0x6d, 0x51, 0x7e, 0x57, 0x23,
	}
)

func TestPacketEncode(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	pkt := protocol.NewPacket(buf)

	if err := pkt.Encode(testStruct); err != nil {
		t.Error(err)
	}

	if !bytes.Equal(buf.Bytes(), testData[:]) {
		t.Error("packet data does not match!")
	}
}

func TestPacketDecode(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	pkt := protocol.NewPacket(buf)
	buf.Write(testData[:])

	var test TestPacketData
	if err := pkt.Decode(&test); err != nil {
		t.Error(err)
	}

	if test != testStruct {
		t.Error("packet data does not match!")
	}
}

func TestDataEncrypt(t *testing.T) {
	buf := make([]byte, len(testData))
	copy(buf, testData[:])

	protocol.EncryptData(buf, []byte(protocol.DEFAULT_KEY))

	if !bytes.Equal(buf, encTestData) {
		t.Error("encrypted data does not match!")
	}
}

func TestDataDecrypt(t *testing.T) {
	buf := make([]byte, len(encTestData))
	copy(buf, encTestData)

	protocol.DecryptData(buf, []byte(protocol.DEFAULT_KEY))

	if !bytes.Equal(buf, testData[:]) {
		t.Error("decrypted data does not match!")
	}
}

func TestCreateNewKey(t *testing.T) {
	key := protocol.CreateNewKey(123456789, 0x1234567890abcdef, 0x1234567890abcdef)

	if !bytes.Equal(key, []byte{0x0, 0x31, 0xb8, 0xcd, 0xd, 0xc3, 0xad, 0x67}) {
		t.Error("key does not match!")
	}
}
