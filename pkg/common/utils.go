package common

import "encoding/binary"

func NumberToBytes(number uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, number)

	return b
}
