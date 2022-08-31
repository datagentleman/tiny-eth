package rlp

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"reflect"
)

func Encode(values ...interface{}) []byte {
	encoding := []byte{}

	for _, v := range values {
		if isList(v) && !isByteArray(v) {
			encodedList := Encode(toList(v)...)

			firstByte := 0xc0 + len(bytes.TrimLeft(encodedList, "\x00"))

			encoding = append(encoding, byte(firstByte))
			encoding = append(encoding, bytes.TrimLeft(encodedList, "\x00")...)
		} else {
			bs, _ := toBytes(v)

			if len(bytes.TrimLeft(bs, "\x00")) == 1 && bytes.TrimLeft(bs, "\x00")[0] <= 0x7f {
				encoding = append(encoding, bytes.TrimLeft(bs, "\x00")[0])
			} else if len(bs) <= 55 {
				firstByte := 0x80 + len(bytes.TrimLeft(bs, "\x00"))

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, bytes.TrimLeft(bs, "\x00")...)
			} else {
				lenInBytes, _ := toBytes(len(bs))
				lenInBytes = bytes.TrimLeft(lenInBytes, "\x00")

				firstByte := 0xb7 + len(lenInBytes)

				encoding = append(encoding, byte(firstByte))
				encoding = append(encoding, lenInBytes...)
				encoding = append(encoding, bs...)
			}
		}
	}

	return encoding
}

func isList(v interface{}) bool {
	slice := reflect.TypeOf(v).Kind() == reflect.Slice
	array := reflect.TypeOf(v).Kind() == reflect.Array

	return (slice || array)
}

func isByteArray(v interface{}) bool {
	if isList(v) {
		switch v.(type) {
		case []uint8:
			return true
		default:
			return false
		}
	}

	return false
}

func toList(v interface{}) []interface{} {
	if !isList(v) {
		return nil
	}

	s := []interface{}{}
	l := reflect.ValueOf(v)

	for i := 0; i < l.Len(); i++ {
		s = append(s, l.Index(i).Interface())
	}

	return s
}

func toBytes(v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case *big.Int:
		return t.Bytes(), nil
	case int:
		return toBytes(int64(t))
	case string:
		return []byte(t), nil
	default:
		var buf bytes.Buffer
		err := binary.Write(&buf, binary.BigEndian, v)
		if err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}
}
