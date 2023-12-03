package protocol

import (
	"crypto/rand"
	"encoding/binary"
)

const (
	DEFAULT_KEY = "m@rQn~W#" // if you change this, make sure to update the test data in protocol_test.go
	KEY_LENGTH  = 8
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

func EncryptData(buff, key []byte) {
	ERSize := len(buff)%(KEY_LENGTH/2+1)*2 + KEY_LENGTH
	xorData(buff, key, len(buff))
	encrypt_byte_change_A(ERSize, buff)
}

func DecryptData(buff, key []byte) {
	ERSize := len(buff)%(KEY_LENGTH/2+1)*2 + KEY_LENGTH
	size := encrypt_byte_change_A(ERSize, buff)
	xorData(buff, key, size)
}

func CreateNewKey(uTime, iv1, iv2 uint64) []byte {
	dEKey := binary.LittleEndian.Uint64([]byte(DEFAULT_KEY))

	num := iv1 + 1
	num2 := iv2 + 1
	key := dEKey * (uTime * num * num2)

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, key)
	return buf
}

func GenSerialKey() (int64, error) {
	tmp := [8]byte{}
	if _, err := rand.Read(tmp[:]); err != nil {
		return 0, err
	}

	// convert to uint64 && return
	return int64(binary.LittleEndian.Uint64(tmp[:])), nil
}
