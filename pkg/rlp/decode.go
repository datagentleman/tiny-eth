package rlp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func Decode(encodings []byte, values ...interface{}) error {
	enc := bytes.NewBuffer(encodings)

	for _, v := range values {
		switch v := v.(type) {
		case *int8:
			encoding := nextEncoding(enc)
			*v = int8(encoding[0])

		case *int16:
			encoding := nextEncoding(enc)
			ensureLen(&encoding, 2)
			*v = (int16)(binary.BigEndian.Uint16(encoding))

		case *int32:
			encoding := nextEncoding(enc)
			ensureLen(&encoding, 4)
			*v = (int32)(binary.BigEndian.Uint32(encoding))

		case *int64:
			encoding := nextEncoding(enc)
			ensureLen(&encoding, 8)
			*v = (int64)(binary.BigEndian.Uint64(encoding))

		default:
			fmt.Println("Unknown type:")
			fmt.Println(reflect.TypeOf(v))
		}
	}

	return nil
}

func nextEncoding(encodings *bytes.Buffer) []byte {
	firstByte := encodings.Next(1)[0]

	if firstByte <= 0x7f {
		return []byte{firstByte}
	}

	return nil
}

func ensureLen(buf *[]byte, length int) {
	bufLen := len(*buf)

	if bufLen < length {
		l := length - bufLen

		b := make([]byte, length)
		*buf = append(b[:l], *buf...)
	}
}
